package database

import (
	"context"
	"fmt"

	"github.com/cloudquery/cloudquery/pkg/client/database/postgres"
	"github.com/cloudquery/cloudquery/pkg/client/database/timescale"
	"github.com/cloudquery/cloudquery/pkg/client/history"
	sdkdb "github.com/cloudquery/cq-provider-sdk/database"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/hashicorp/go-hclog"
)

type DialectExecutor interface {
	// Setup is called on the dialect on initialization, returns the DSN (modified if necessary) to use for migrations
	Setup(context.Context) (string, error)

	// Validate is called before startup to check that the dialect can execute properly. If returns true and error is set, the error is merely logged.
	Validate(context.Context) (bool, error)

	// Finalize is called after migrations and upgrades are run
	Finalize(context.Context) error
}

var (
	_ DialectExecutor = (*postgres.Executor)(nil)
	_ DialectExecutor = (*timescale.Executor)(nil)
)

func GetExecutor(logger hclog.Logger, dsn string, c *history.Config) (schema.DialectType, DialectExecutor, error) {
	if dsn == "" {
		return schema.Postgres, nil, fmt.Errorf("missing DSN")
	}

	dType, dsn, err := sdkdb.ParseDialectDSN(dsn)
	if err != nil {
		return dType, nil, err
	}

	switch dType {
	case schema.Postgres:
		return dType, postgres.New(logger, dsn), nil
	case schema.TSDB:
		ts, err := timescale.New(logger, dsn, c)
		if err != nil {
			return dType, nil, err
		}
		return dType, ts, nil
	default:
		return dType, nil, fmt.Errorf("unhandled dialect type")
	}
}
