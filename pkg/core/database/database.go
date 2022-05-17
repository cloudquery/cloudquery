package database

import (
	"context"
	"fmt"

	"github.com/cloudquery/cloudquery/pkg/core/database/postgres"
	sdkdb "github.com/cloudquery/cq-provider-sdk/database"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

type DialectExecutor interface {
	// Identifier returns a unique identifier for the database if possible, or "", false
	Identifier(context.Context) (string, bool)

	// Validate is called before startup to check that the dialect can execute properly. If returns true and error is set, the error is merely logged.
	Validate(context.Context) (bool, error)
}

var (
	_ DialectExecutor = (*postgres.Executor)(nil)
)

func GetExecutor(dsn string) (schema.DialectType, DialectExecutor, error) {
	if dsn == "" {
		return schema.Postgres, nil, fmt.Errorf("missing DSN")
	}

	dType, dsn, err := sdkdb.ParseDialectDSN(dsn)
	if err != nil {
		return dType, nil, err
	}

	switch dType {
	case schema.Postgres:
		return dType, postgres.New(dsn), nil
	case schema.TSDB:
		return dType, nil, fmt.Errorf("history feature is removed. See more at https://TODO") // TODO
	default:
		return dType, nil, fmt.Errorf("unhandled dialect type")
	}
}
