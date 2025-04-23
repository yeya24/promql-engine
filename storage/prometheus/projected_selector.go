package prometheus

import (
	"context"
	"sync"

	"github.com/prometheus/prometheus/model/labels"
	"github.com/prometheus/prometheus/storage"
	"github.com/thanos-io/promql-engine/logicalplan"
)

type projectedSelector struct {
	selector   SeriesSelector
	projection *logicalplan.Projection

	once   sync.Once
	series []SignedSeries
}

func NewProjectedSelector(selector SeriesSelector, projection *logicalplan.Projection) SeriesSelector {
	return &projectedSelector{
		selector:   selector,
		projection: projection,
	}
}

func (f *projectedSelector) Matchers() []*labels.Matcher {
	return f.selector.Matchers()
}

func (f *projectedSelector) GetSeries(ctx context.Context, shard, numShards int) ([]SignedSeries, error) {
	var err error
	f.once.Do(func() { err = f.loadSeries(ctx) })
	if err != nil {
		return nil, err
	}

	return seriesShard(f.series, shard, numShards), nil
}

func (f *projectedSelector) loadSeries(ctx context.Context) error {
	series, err := f.selector.GetSeries(ctx, 0, 1)
	if err != nil {
		return err
	}
	if f.projection == nil || (!f.projection.Include && len(f.projection.Labels) == 0) {
		f.series = series
		return nil
	}

	var i uint64
	f.series = make([]SignedSeries, 0, len(series))
	b := labels.NewBuilder(labels.EmptyLabels())
	for _, s := range series {
		b.Reset(s.Labels())
		if f.projection.Include {
			b.Keep(f.projection.Labels...)
		} else {
			b.Del(f.projection.Labels...)
		}

		f.series = append(f.series, SignedSeries{
			Series:    &projectedSeries{Series: s, lset: b.Labels()},
			Signature: i,
		})
		i++
	}

	return nil
}

// projectedSeries wraps a storage.Series but returns projected labels
type projectedSeries struct {
	storage.Series
	lset labels.Labels
}

func (s *projectedSeries) Labels() labels.Labels {
	return s.lset
}
