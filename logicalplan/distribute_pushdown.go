package logicalplan

import (
	"github.com/thanos-community/promql-engine/api"
	"github.com/thanos-community/promql-engine/internal/prometheus/parser"
)

type DistributedPushDownOptimizer struct {
	Endpoints api.RemoteEndpoints
}

func (m DistributedPushDownOptimizer) Optimize(plan parser.Expr, opts *Opts) parser.Expr {
	engines := m.Endpoints.Engines(opts.UserID, opts.Query, opts.Start, opts.End, opts.LookbackDelta)
	if len(engines) == 1 {
		return RemoteExecution{
			Engine:          engines[0],
			Query:           plan.String(),
			QueryRangeStart: opts.Start,
		}
	} else if len(engines) == 0 {
		return plan
	}

	traverseBottomUp(nil, &plan, func(parent, current *parser.Expr) (stop bool) {
		// If the current operation is not distributive, stop the traversal.
		if !isDistributive(current) {
			return true
		}

		// If the current node is an aggregation, distribute the operation and
		// stop the traversal.
		if aggr, ok := (*current).(*parser.AggregateExpr); ok {
			localAggregation := aggr.Op
			if aggr.Op == parser.COUNT {
				localAggregation = parser.SUM
			}

			remoteAggregation := newRemoteAggregation(aggr, engines)
			subQueries := m.distributeQuery(&remoteAggregation, engines, opts)
			*current = &parser.AggregateExpr{
				Op:       localAggregation,
				Expr:     subQueries,
				Param:    aggr.Param,
				Grouping: aggr.Grouping,
				Without:  aggr.Without,
				PosRange: aggr.PosRange,
			}
			return true
		}

		// If the parent operation is distributive, continue the traversal.
		if isDistributive(parent) {
			return false
		}

		*current = m.distributeQuery(current, engines, opts)
		return true
	})
	return plan
}

func (m DistributedPushDownOptimizer) distributeQuery(expr *parser.Expr, engines []api.RemoteEngine, opts *Opts) parser.Expr {
	remoteQueries := make(RemoteExecutions, 0, len(engines))
	for _, e := range engines {
		remoteQueries = append(remoteQueries, RemoteExecution{
			Engine:          e,
			Query:           (*expr).String(),
			QueryRangeStart: opts.Start,
		})
	}

	return Coalesce{
		Expressions: remoteQueries,
	}
}
