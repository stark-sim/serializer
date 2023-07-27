package db_tools

import (
	"context"
	"github.com/stark-sim/serializer/ent"
)

func P(next ent.Querier) ent.Querier {
	return ent.QuerierFunc(func(ctx context.Context, query ent.Query) (ent.Value, error) {

		return next.Query(ctx, query)
	})
}
