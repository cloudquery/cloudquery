package provider_test

import (
	"os"
	"testing"

	"github.com/cloudquery/cq-provider-azure/resources/provider"
	"github.com/cloudquery/cq-provider-sdk/migration"
)

func TestMigrationPostgreSQL(t *testing.T) {
	dsn := os.Getenv("CQ_MIGRATION_TEST_PG_DSN")
	if dsn == "" {
		dsn = "postgres://postgres:pass@localhost:5432/postgres?sslmode=disable"
	}
	migration.RunMigrationsTestWithNewDB(t, dsn, "testpgmigration", provider.Provider(), []string{"latest"})
}

func TestMigrationTimescaleDB(t *testing.T) {
	dsn := os.Getenv("CQ_MIGRATION_TEST_TSDB_DSN")
	if dsn == "" {
		dsn = "tsdb://postgres:pass@localhost:5432/postgres?sslmode=disable"
	}
	migration.RunMigrationsTestWithNewDB(t, dsn, "testtsdbmigration", provider.Provider(), []string{"latest"})
}
