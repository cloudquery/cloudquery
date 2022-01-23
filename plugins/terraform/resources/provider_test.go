package resources_test

import (
	"testing"

	"github.com/cloudquery/cq-provider-sdk/migration"
	"github.com/cloudquery/cq-provider-terraform/resources"
)

func TestMigrations(t *testing.T) {
	migration.RunMigrationsTest(t, resources.Provider(), nil)
}
