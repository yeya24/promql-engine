// Copyright (c) The Thanos Community Authors.
// Licensed under the Apache License 2.0.

package logicalplan

import (
	"sort"
	"testing"

	"github.com/efficientgo/core/testutil"
	"github.com/prometheus/prometheus/model/labels"
	"github.com/prometheus/prometheus/promql/parser"
	"github.com/thanos-io/promql-engine/query"
)

func TestMergeSelects(t *testing.T) {
	cases := []struct {
		expr     string
		expected string
	}{
		{
			expr:     `X{a="b"}/X`,
			expected: `filter([a="b"], X) / X`,
		},
		{
			expr:     `floor(X{a="b"})/X`,
			expected: `floor(filter([a="b"], X)) / X`,
		},
		{
			expr:     `X/floor(X{a="b"})`,
			expected: `X / floor(filter([a="b"], X))`,
		},
		{
			expr:     `X{a="b"}/floor(X)`,
			expected: `filter([a="b"], X) / floor(X)`,
		},
		{
			expr:     `quantile by (pod) (scalar(min(http_requests_total)), http_requests_total)`,
			expected: `quantile by (pod) (scalar(min(http_requests_total)), http_requests_total)`,
		},
	}
	optimizers := []Optimizer{MergeSelectsOptimizer{}}
	for _, tcase := range cases {
		t.Run(tcase.expr, func(t *testing.T) {
			expr, err := parser.ParseExpr(tcase.expr)
			testutil.Ok(t, err)

			plan := NewFromAST(expr, &query.Options{}, PlanOptions{})
			optimizedPlan, _ := plan.Optimize(optimizers)
			testutil.Equals(t, tcase.expected, renderExprTree(optimizedPlan.Root()))
		})
	}
}

func TestMergeProjections(t *testing.T) {
	cases := []struct {
		name        string
		projection1 *Projection
		projection2 *Projection
		expected    *Projection
	}{
		{
			name:        "nil projection1",
			projection1: nil,
			projection2: &Projection{Include: true, Labels: []string{"a", "b"}},
			expected:    &Projection{Include: true, Labels: []string{"a", "b"}},
		},
		{
			name:        "nil projection2",
			projection1: &Projection{Include: true, Labels: []string{"a", "b"}},
			projection2: nil,
			expected:    &Projection{Include: true, Labels: []string{"a", "b"}},
		},
		{
			name:        "both include projections",
			projection1: &Projection{Include: true, Labels: []string{"a", "b"}},
			projection2: &Projection{Include: true, Labels: []string{"b", "c"}},
			expected:    &Projection{Include: true, Labels: []string{"a", "b", "c"}},
		},
		{
			name:        "both exclude projections",
			projection1: &Projection{Include: false, Labels: []string{"a", "b", "c"}},
			projection2: &Projection{Include: false, Labels: []string{"b", "c", "d"}},
			expected:    &Projection{Include: false, Labels: []string{"b", "c"}},
		},
		{
			name:        "mixed include/exclude projections",
			projection1: &Projection{Include: true, Labels: []string{"a", "b", "c"}},
			projection2: &Projection{Include: false, Labels: []string{"b", "c", "d"}},
			expected:    &Projection{Include: false, Labels: []string{"d"}},
		},
		{
			name:        "mixed exclude/include projections",
			projection1: &Projection{Include: false, Labels: []string{"a", "b", "c"}},
			projection2: &Projection{Include: true, Labels: []string{"b", "c", "d"}},
			expected:    &Projection{Include: false, Labels: []string{"a"}},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := mergeProjections(tc.projection1, tc.projection2)
			testutil.Equals(t, tc.expected.Include, result.Include)
			sort.Strings(tc.expected.Labels)
			sort.Strings(result.Labels)
			testutil.Equals(t, tc.expected.Labels, result.Labels)
		})
	}
}

func TestReplaceProjections(t *testing.T) {
	cases := []struct {
		name          string
		selectors     matcherHeap
		projectionMap map[string]*Projection
		node          Node
		expectedNode  Node
		shouldReplace bool
	}{
		{
			name:          "empty projection map",
			selectors:     matcherHeap{},
			projectionMap: map[string]*Projection{},
			node: &VectorSelector{
				VectorSelector: &parser.VectorSelector{
					LabelMatchers: []*labels.Matcher{
						{Type: labels.MatchEqual, Name: labels.MetricName, Value: "metric"},
					},
				},
				Projection: &Projection{Include: true, Labels: []string{"a"}},
			},
			expectedNode: &VectorSelector{
				VectorSelector: &parser.VectorSelector{
					LabelMatchers: []*labels.Matcher{
						{Type: labels.MatchEqual, Name: labels.MetricName, Value: "metric"},
					},
				},
				Projection: &Projection{Include: true, Labels: []string{"a"}},
			},
			shouldReplace: false,
		},
		{
			name: "vector selector with no projection",
			selectors: matcherHeap{
				"metric": []*labels.Matcher{{Type: labels.MatchEqual, Name: labels.MetricName, Value: "metric"}},
			},
			projectionMap: map[string]*Projection{
				"metric": {Include: true, Labels: []string{"a", "b"}},
			},
			node: &VectorSelector{
				VectorSelector: &parser.VectorSelector{
					LabelMatchers: []*labels.Matcher{
						{Type: labels.MatchEqual, Name: labels.MetricName, Value: "metric"},
					},
				},
			},
			expectedNode: &VectorSelector{
				VectorSelector: &parser.VectorSelector{
					LabelMatchers: []*labels.Matcher{
						{Type: labels.MatchEqual, Name: labels.MetricName, Value: "metric"},
					},
				},
			},
			shouldReplace: false,
		},
		{
			name: "matching projection with different labels",
			selectors: matcherHeap{
				"metric": []*labels.Matcher{{Type: labels.MatchEqual, Name: labels.MetricName, Value: "metric"}},
			},
			projectionMap: map[string]*Projection{
				"metric": {Include: true, Labels: []string{"a", "b"}},
			},
			node: &VectorSelector{
				VectorSelector: &parser.VectorSelector{
					LabelMatchers: []*labels.Matcher{
						{Type: labels.MatchEqual, Name: labels.MetricName, Value: "metric"},
					},
				},
				Projection: &Projection{Include: true, Labels: []string{"c"}},
			},
			expectedNode: &VectorSelector{
				VectorSelector: &parser.VectorSelector{
					LabelMatchers: []*labels.Matcher{
						{Type: labels.MatchEqual, Name: labels.MetricName, Value: "metric"},
					},
				},
				Projection:       &Projection{Include: true, Labels: []string{"a", "b"}},
				ProjectionFilter: &Projection{Include: true, Labels: []string{"c", "__series_hash__"}},
			},
			shouldReplace: true,
		},
		{
			name: "empty final projection",
			selectors: matcherHeap{
				"metric": []*labels.Matcher{{Type: labels.MatchEqual, Name: labels.MetricName, Value: "metric"}},
			},
			projectionMap: map[string]*Projection{
				"metric": {Include: false, Labels: []string{}},
			},
			node: &VectorSelector{
				VectorSelector: &parser.VectorSelector{
					LabelMatchers: []*labels.Matcher{
						{Type: labels.MatchEqual, Name: labels.MetricName, Value: "metric"},
					},
				},
				Projection: &Projection{Include: true, Labels: []string{"a"}},
			},
			expectedNode: &VectorSelector{
				VectorSelector: &parser.VectorSelector{
					LabelMatchers: []*labels.Matcher{
						{Type: labels.MatchEqual, Name: labels.MetricName, Value: "metric"},
					},
				},
				Projection: &Projection{},
			},
			shouldReplace: true,
		},
		{
			name: "equal projections",
			selectors: matcherHeap{
				"metric": []*labels.Matcher{{Type: labels.MatchEqual, Name: labels.MetricName, Value: "metric"}},
			},
			projectionMap: map[string]*Projection{
				"metric": {Include: true, Labels: []string{"a", "b"}},
			},
			node: &VectorSelector{
				VectorSelector: &parser.VectorSelector{
					LabelMatchers: []*labels.Matcher{
						{Type: labels.MatchEqual, Name: labels.MetricName, Value: "metric"},
					},
				},
				Projection: &Projection{Include: true, Labels: []string{"a", "b"}},
			},
			expectedNode: &VectorSelector{
				VectorSelector: &parser.VectorSelector{
					LabelMatchers: []*labels.Matcher{
						{Type: labels.MatchEqual, Name: labels.MetricName, Value: "metric"},
					},
				},
				Projection: &Projection{Include: true, Labels: []string{"a", "b"}},
			},
			shouldReplace: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			node := tc.node
			replaceProjections(tc.selectors, tc.projectionMap, &node)

			vs1, ok1 := node.(*VectorSelector)
			vs2, ok2 := tc.expectedNode.(*VectorSelector)
			testutil.Equals(t, ok1, ok2)
			if !ok1 {
				return
			}

			testutil.Equals(t, vs2.LabelMatchers, vs1.LabelMatchers)

			if vs1.Projection != nil && vs2.Projection != nil {
				testutil.Equals(t, vs2.Projection.Include, vs1.Projection.Include)
				sort.Strings(vs1.Projection.Labels)
				sort.Strings(vs2.Projection.Labels)
				testutil.Equals(t, vs2.Projection.Labels, vs1.Projection.Labels)
			} else {
				testutil.Equals(t, vs2.Projection, vs1.Projection)
			}

			if vs1.ProjectionFilter != nil && vs2.ProjectionFilter != nil {
				testutil.Equals(t, vs2.ProjectionFilter.Include, vs1.ProjectionFilter.Include)
				sort.Strings(vs1.ProjectionFilter.Labels)
				sort.Strings(vs2.ProjectionFilter.Labels)
				testutil.Equals(t, vs2.ProjectionFilter.Labels, vs1.ProjectionFilter.Labels)
			} else {
				testutil.Equals(t, vs2.ProjectionFilter, vs1.ProjectionFilter)
			}
		})
	}
}

func TestRemoveFilterLabelsFromProjections(t *testing.T) {
	cases := []struct {
		name            string
		projectionMap   map[string]*Projection
		filterLabelsMap map[string]map[string]struct{}
		expected        map[string]*Projection
	}{
		{
			name:          "empty projection map",
			projectionMap: map[string]*Projection{},
			filterLabelsMap: map[string]map[string]struct{}{
				"metric": {"label1": struct{}{}, "label2": struct{}{}},
			},
			expected: map[string]*Projection{},
		},
		{
			name: "no matching filter labels",
			projectionMap: map[string]*Projection{
				"metric1": {Include: true, Labels: []string{"a", "b"}},
			},
			filterLabelsMap: map[string]map[string]struct{}{
				"metric2": {"label1": struct{}{}, "label2": struct{}{}},
			},
			expected: map[string]*Projection{
				"metric1": {Include: true, Labels: []string{"a", "b"}},
			},
		},
		{
			name: "matching filter labels with include projection",
			projectionMap: map[string]*Projection{
				"metric": {Include: true, Labels: []string{"a", "b"}},
			},
			filterLabelsMap: map[string]map[string]struct{}{
				"metric": {"c": struct{}{}, "d": struct{}{}},
			},
			expected: map[string]*Projection{
				"metric": {Include: true, Labels: []string{"a", "b", "c", "d"}},
			},
		},
		{
			name: "matching filter labels with exclude projection",
			projectionMap: map[string]*Projection{
				"metric": {Include: false, Labels: []string{"a", "b", "c"}},
			},
			filterLabelsMap: map[string]map[string]struct{}{
				"metric": {"b": struct{}{}, "d": struct{}{}},
			},
			expected: map[string]*Projection{
				"metric": {Include: false, Labels: []string{"a", "c"}},
			},
		},
		{
			name: "multiple metrics with mixed projections",
			projectionMap: map[string]*Projection{
				"metric1": {Include: true, Labels: []string{"a", "b"}},
				"metric2": {Include: false, Labels: []string{"c", "d"}},
				"metric3": {Include: true, Labels: []string{"e", "f"}},
			},
			filterLabelsMap: map[string]map[string]struct{}{
				"metric1": {"x": struct{}{}, "y": struct{}{}},
				"metric2": {"d": struct{}{}, "z": struct{}{}},
			},
			expected: map[string]*Projection{
				"metric1": {Include: true, Labels: []string{"a", "b", "x", "y"}},
				"metric2": {Include: false, Labels: []string{"c"}},
				"metric3": {Include: true, Labels: []string{"e", "f"}},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a copy of the projection map to avoid modifying the test case
			projMap := make(map[string]*Projection)
			for k, v := range tc.projectionMap {
				projMap[k] = &Projection{
					Include: v.Include,
					Labels:  append([]string{}, v.Labels...),
				}
			}

			removeFilterLabelsFromProjections(projMap, tc.filterLabelsMap)

			testutil.Equals(t, len(tc.expected), len(projMap))
			for metric, expectedProj := range tc.expected {
				actualProj, ok := projMap[metric]
				testutil.Assert(t, ok, "missing projection for metric %s", metric)
				testutil.Equals(t, expectedProj.Include, actualProj.Include)

				// Sort labels for stable comparison
				sort.Strings(expectedProj.Labels)
				sort.Strings(actualProj.Labels)
				testutil.Equals(t, expectedProj.Labels, actualProj.Labels)
			}
		})
	}
}
