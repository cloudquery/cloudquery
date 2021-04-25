package resources_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/georgysavva/scany/pgxscan"

	"github.com/cloudquery/cq-provider-sdk/logging"
	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/cloudquery/faker/v3"
	"github.com/hashicorp/go-hclog"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

type ProviderTestData struct {
	Provider  func() *provider.Provider
	Resources []ResourceTestData
}

type ResourceTestData struct {
	Table     *schema.Table
	Config    interface{}
	Configure func(logger hclog.Logger, data interface{}) (schema.ClientMeta, error)
}

func testResource(t *testing.T, provider func() *provider.Provider, resource ResourceTestData) {
	testData := ProviderTestData{
		Provider:  provider,
		Resources: []ResourceTestData{resource},
	}
	if err := faker.SetRandomMapAndSliceMinSize(1); err != nil {
		t.Fatal(err)
	}
	if err := faker.SetRandomMapAndSliceMaxSize(1); err != nil {
		t.Fatal(err)
	}
	conn, err := setupDatabase()
	if err != nil {
		t.Fatal(err)
	}
	testProvider := testData.Provider()
	testProvider.Logger = logging.New(hclog.DefaultOptions)
	testProvider.Configure = resource.Configure
	err = testProvider.Init("", "host=localhost user=postgres password=pass DB.name=postgres port=5432", false)
	assert.Nil(t, err)
	data, err := yaml.Marshal(resource.Config)
	assert.Nil(t, err)
	err = testProvider.Fetch(data)
	assert.Nil(t, err)
	verifyNoEmptyColumns(t, resource, conn)
}

func setupDatabase() (*pgx.Conn, error) {
	dbCfg, err := pgx.ParseConfig(getEnv("DATABASE_URL",
		"host=localhost user=postgres password=pass DB.name=postgres port=5432"))
	if err != nil {
		return nil, fmt.Errorf("failed to parse config. %w", err)
	}
	ctx := context.Background()
	conn, err := pgx.ConnectConfig(ctx, dbCfg)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database. %w", err)
	}
	return conn, nil

}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func verifyNoEmptyColumns(t *testing.T, tc ResourceTestData, conn pgxscan.Querier) {
	// Test that we don't have missing columns and have exactly one entry for each table
	for _, table := range getTablesFromMainTable(tc.Table) {
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
