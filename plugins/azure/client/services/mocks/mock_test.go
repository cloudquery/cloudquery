package mocks_test

import (
	"context"
	"fmt"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/resources"
	"os"
	"testing"

	"github.com/cloudquery/cq-provider-sdk/logging"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/hashicorp/go-hclog"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

type TestResource struct {
	resource    string
	mockBuilder func(*testing.T, *gomock.Controller) services.Services
	mainTable   *schema.Table
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func TestResources(t *testing.T) {
	dbCfg, err := pgx.ParseConfig(getEnv("DATABASE_URL",
		"host=localhost user=postgres password=pass DB.name=postgres port=5432"))
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	conn, err := pgx.ConnectConfig(ctx, dbCfg)
	if err != nil {
		t.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(ctx)
	_ = faker.SetRandomMapAndSliceMinSize(1)
	err = faker.SetRandomMapAndSliceMaxSize(1)
	if err != nil {
		t.Fatal(err)
	}
	ctrl := gomock.NewController(t)
	testResourcesTable := []TestResource{
		{
			resource:    "compute.disks",
			mockBuilder: buildComputeDiskMock,
			mainTable:   resources.ComputeDisks(),
		},
		{
			resource:    "resources.groups",
			mockBuilder: buildResourceGroupMock,
			mainTable:   resources.ResourceGroups(),
		},
		{
			resource:    "keyvault.vaults",
			mockBuilder: buildKeyVaultMock,
			mainTable:   resources.KeyVaultVaults(),
		},
		{
			resource:    "storage.accounts",
			mockBuilder: buildStorageMock,
			mainTable:   resources.StorageAccounts(),
		},
		{
			resource:    "mysql.servers",
			mockBuilder: buildMySQLServerMock,
			mainTable:   resources.MySQLServers(),
		},
		{
			resource:    "postgresql.servers",
			mockBuilder: buildPostgresServerMock,
			mainTable:   resources.PostgresqlServers(),
		},
		{
			resource:    "sql.servers",
			mockBuilder: buildSQLServerMock,
			mainTable:   resources.SQLServers(),
		},
		// Faker is recursive, cause overflow
		//{
		//	resource: "network.virtual_networks",
		//	mockBuilder: buildVirtualNetworkMock,
		//	mainTable: resources.NetworkVirtualNetworks(),
		//},
	}
	for _, tc := range testResourcesTable {
		t.Run(tc.resource, func(t *testing.T) {
			cfg := client.Config{
				Subscriptions: []string{"test_sub"},
				Resources: []client.Resource{{
					Name: tc.resource,
				},
				},
			}
			testProvider := resources.Provider()
			testProvider.Logger = logging.New(hclog.DefaultOptions)
			testProvider.Configure = func(logger hclog.Logger, i interface{}) (schema.ClientMeta, error) {
				c := client.NewAzureClient(logging.New(&hclog.LoggerOptions{
					Level: hclog.Warn,
				}), []string{"test_sub"})
				c.SetSubscriptionServices("test_sub", tc.mockBuilder(t, ctrl))
				return c, nil
			}
			err := testProvider.Init("", "host=localhost user=postgres password=pass DB.name=postgres port=5432", false)
			assert.Nil(t, err)
			data, err := yaml.Marshal(cfg)
			assert.Nil(t, err)
			err = testProvider.Fetch(data)
			assert.Nil(t, err)
			verifyNoEmptyColumns(t, tc, conn)
		})
	}
}

func verifyNoEmptyColumns(t *testing.T, tc TestResource, conn *pgx.Conn) {
	// Test that we don't have missing columns and have exactly one entry for each table
	for _, table := range getTablesFromMainTable(tc.mainTable) {

		query := fmt.Sprintf("select * FROM %s ", table)
		rows, err := conn.Query(context.Background(), query)
		if err != nil {
			t.Fatal(err)
		}
		count := 0
		for rows.Next() {
			count += 1
		}
		if count < 1 {
			t.Fatalf("expected to have at least 1 entry at table %s got %d", table, count)
		}

		query = fmt.Sprintf("select t.* FROM %s as t WHERE to_jsonb(t) = jsonb_strip_nulls(to_jsonb(t))", table)
		rows, err = conn.Query(context.Background(), query)
		if err != nil {
			t.Fatal(err)
		}
		count = 0
		for rows.Next() {
			count += 1
		}
		if count < 1 {
			t.Fatalf("row at table %s has an empty column", table)
		}
	}
}

func getTablesFromMainTable(table *schema.Table) []string {
	var res []string
	res = append(res, table.Name)
	for _, t := range table.Relations {
		res = append(res, getTablesFromMainTable(t)...)
	}
	return res
}
