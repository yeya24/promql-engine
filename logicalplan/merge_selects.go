// Copyright (c) The Thanos Community Authors.
// Licensed under the Apache License 2.0.

package logicalplan

import (
	"github.com/thanos-io/promql-engine/query"

	"github.com/prometheus/prometheus/model/labels"
	"github.com/prometheus/prometheus/util/annotations"
)

// MergeSelectsOptimizer optimizes a binary expression where
// one select is a superset of the other select.
// For example, the expression:
//
//	metric{a="b", c="d"} / scalar(metric{a="b"}) becomes:
//	Filter(c="d", metric{a="b"}) / scalar(metric{a="b"}).
//
// The engine can then cache the result of `metric{a="b"}`
// and apply an additional filter for {c="d"}.
type MergeSelectsOptimizer struct{}

func (m MergeSelectsOptimizer) Optimize(plan Node, _ *query.Options) (Node, annotations.Annotations) {
	heap := make(matcherHeap)
	projectionMap := make(map[string]*Projection)
	filterLabelsMap := make(map[string]map[string]struct{})
	extractSelectors(heap, projectionMap, plan)
	replaceMatchers(heap, filterLabelsMap, &plan)
	removeFilterLabelsFromProjections(projectionMap, filterLabelsMap)
	replaceProjections(heap, projectionMap, &plan)

	return plan, nil
}

func extractSelectors(selectors matcherHeap, projectionMap map[string]*Projection, expr Node) {
	Traverse(&expr, func(node *Node) {
		e, ok := (*node).(*VectorSelector)
		if !ok {
			return
		}
		for _, l := range e.LabelMatchers {
			if l.Name == labels.MetricName {
				selectors.add(l.Value, e.LabelMatchers)
				if e.Projection != nil {
					projectionMap[l.Value] = mergeProjections(projectionMap[l.Value], e.Projection)
				}
			}
		}
	})
}

func mergeProjections(projection1, projection2 *Projection) *Projection {
	if projection1 == nil {
		return projection2
	}
	if projection2 == nil {
		return projection1
	}

	if projection1.Include && projection2.Include {
		return &Projection{
			Include: true,
			Labels:  union(projection1.Labels, projection2.Labels),
		}
	}
	if !projection1.Include && !projection2.Include {
		return &Projection{
			Include: false,
			Labels:  intersect(projection1.Labels, projection2.Labels),
		}
	}

	includeSide := projection1
	excludeSide := projection2
	if !projection1.Include {
		includeSide = projection2
		excludeSide = projection1
	}

	return &Projection{
		Include: false,
		Labels:  subtract(excludeSide.Labels, includeSide.Labels),
	}
}

func replaceMatchers(selectors matcherHeap, filterLabelsMap map[string]map[string]struct{}, expr *Node) {
	Traverse(expr, func(node *Node) {
		var matchers []*labels.Matcher
		switch e := (*node).(type) {
		case *VectorSelector:
			matchers = e.LabelMatchers
		default:
			return
		}

		for _, l := range matchers {
			if l.Name != labels.MetricName {
				continue
			}
			replacement, found, _ := selectors.findReplacement(l.Value, matchers)
			if !found {
				continue
			}

			// Make a copy of the original selectors to avoid modifying them while
			// trimming filters.
			filters := make([]*labels.Matcher, len(matchers))
			copy(filters, matchers)

			// All replacements are done on metrics name only,
			// so we can drop the explicit metric name selector.
			filters = dropMatcher(labels.MetricName, filters)

			// Drop filters which are already present as matchers in the replacement selector.
			for _, s := range replacement {
				for _, f := range filters {
					if s.Name == f.Name && s.Value == f.Value && s.Type == f.Type {
						filters = dropMatcher(f.Name, filters)
					}
				}
			}

			for _, f := range filters {
				if _, ok := filterLabelsMap[l.Value]; !ok {
					filterLabelsMap[l.Value] = make(map[string]struct{})
				}
				filterLabelsMap[l.Value][f.Name] = struct{}{}
			}

			switch e := (*node).(type) {
			case *VectorSelector:
				e.LabelMatchers = replacement
				e.Filters = filters
			}

			return
		}
	})
}

func replaceProjections(selectors matcherHeap, projectionMap map[string]*Projection, expr *Node) {
	// Return early if there are no projections to merge.
	if len(projectionMap) == 0 {
		return
	}

	Traverse(expr, func(node *Node) {
		var matchers []*labels.Matcher
		switch e := (*node).(type) {
		case *VectorSelector:
			if e.Projection == nil {
				return
			}
			matchers = e.LabelMatchers
		default:
			return
		}

		for _, l := range matchers {
			if l.Name != labels.MetricName {
				continue
			}
			projection, foundProjection := projectionMap[l.Value]
			if !foundProjection {
				return
			}

			_, found, isTop := selectors.findReplacement(l.Value, matchers)
			if !found && !isTop {
				continue
			}

			switch e := (*node).(type) {
			case *VectorSelector:
				// No need to update projection filter if it is the same as the final projection.
				if e.Projection.Equals(*projection) {
					return
				}
				// If the final projection is empty projection, we just disable merge projection
				// as empty projection won't fetch series hash so it might cause correctness issues.
				if !projection.Include && len(projection.Labels) == 0 {
					e.Projection = &Projection{}
					return
				}

				newProjection := extendProjection(*e.Projection, []string{"__series_hash__"})
				e.ProjectionFilter = &newProjection
				e.Projection = projection
			}

			return
		}
	})
}

// filter labels must be included in the projection in order to filter out the series.
func removeFilterLabelsFromProjections(projectionMap map[string]*Projection, filterLabelsMap map[string]map[string]struct{}) {
	// Return early if there are no projections to merge.
	if len(projectionMap) == 0 {
		return
	}

	for name, filterLabels := range filterLabelsMap {
		if _, ok := projectionMap[name]; !ok {
			continue
		}
		m := make([]string, 0, len(filterLabels))
		for label := range filterLabels {
			m = append(m, label)
		}

		projectionMap[name] = mergeProjections(projectionMap[name], &Projection{
			Include: true,
			Labels:  m,
		})
	}
}

func dropMatcher(matcherName string, originalMatchers []*labels.Matcher) []*labels.Matcher {
	i := 0
	for i < len(originalMatchers) {
		l := originalMatchers[i]
		if l.Name == matcherName {
			originalMatchers = append(originalMatchers[:i], originalMatchers[i+1:]...)
		} else {
			i++
		}
	}
	return originalMatchers
}

func matcherToMap(matchers []*labels.Matcher) map[string]*labels.Matcher {
	r := make(map[string]*labels.Matcher, len(matchers))
	for i := 0; i < len(matchers); i++ {
		r[matchers[i].Name] = matchers[i]
	}
	return r
}

// matcherHeap is a set of the most selective label matchers
// for each metrics discovered in a PromQL expression.
// The selectivity of a matcher is defined by how many series are
// matched by it. Since matchers in PromQL are open, selectors
// with the least amount of matchers are typically the most selective ones.
type matcherHeap map[string][]*labels.Matcher

func (m matcherHeap) add(metricName string, lessSelective []*labels.Matcher) {
	moreSelective, ok := m[metricName]
	if !ok {
		m[metricName] = lessSelective
		return
	}

	if len(lessSelective) < len(moreSelective) {
		moreSelective = lessSelective
	}

	m[metricName] = moreSelective
}

func (m matcherHeap) findReplacement(metricName string, matcher []*labels.Matcher) (matchers []*labels.Matcher, found bool, isTop bool) {
	top, ok := m[metricName]
	if !ok {
		return nil, false, false
	}

	matcherSet := matcherToMap(matcher)
	topSet := matcherToMap(top)
	for k, v := range topSet {
		m, ok := matcherSet[k]
		if !ok {
			return nil, false, false
		}

		equals := v.Name == m.Name && v.Type == m.Type && v.Value == m.Value
		if !equals {
			return nil, false, false
		}
	}

	// The top matcher and input matcher are equal. No replacement needed.
	if len(topSet) == len(matcherSet) {
		return nil, false, true
	}

	return top, true, false
}
