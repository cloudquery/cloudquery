package provider_test

import (
	"testing"

	"github.com/cloudquery/cq-provider-azure/resources/provider"
	"github.com/cloudquery/cq-provider-sdk/migration"
)

func TestMigrations(t *testing.T) {
	migration.RunMigrationsTest(t, provider.Provider(), nil)
}
