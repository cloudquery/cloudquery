package client

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/destination/mongodb/v2/client/spec"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func newMigrateClient(t *testing.T) *Client {
	t.Helper()
	ctx := context.Background()
	s := &spec.Spec{ConnectionString: getTestConnection(), Database: "destination_mongodb_assess_test"}
	specBytes, err := json.Marshal(s)
	require.NoError(t, err)
	pc, err := New(ctx, zerolog.Nop(), specBytes, plugin.NewClientOptions{})
	require.NoError(t, err)
	c := pc.(*Client)
	require.NoError(t, c.client.Database(s.Database).Drop(ctx))
	t.Cleanup(func() {
		_ = c.client.Database(s.Database).Drop(ctx)
		_ = c.Close(ctx)
	})
	return c
}

func TestAssessTablesMatchesMigrate(t *testing.T) {
	primaryKeyAdded := withoutPrimaryKeysPair(oktaPolicyIDPair())
	primaryKeyAdded.New = oktaPolicyIDPair().New

	tests := []struct {
		name string
		pair plugin.TablePair
	}{
		{name: "datadog tags list to json", pair: datadogTagsPair()},
		{name: "okta policy_id primary key", pair: oktaPolicyIDPair()},
		{name: "okta append write mode", pair: withoutPrimaryKeysPair(oktaPolicyIDPair())},
		{name: "primary key added", pair: primaryKeyAdded},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			c := newMigrateClient(t)
			require.NoError(t, c.MigrateTables(ctx, message.WriteMigrateTables{{Table: tc.pair.Old}}))
			collection := c.client.Database(c.spec.Database).Collection(tc.pair.Old.Name)
			seedDocuments(t, collection, tc.pair.Old)

			finding := assessOneTable(t, tc.pair)
			safeErr := c.MigrateTables(ctx, message.WriteMigrateTables{{Table: tc.pair.New}})
			if finding.Category == plugin.AssessCategoryManualMigrationRequired {
				require.Error(t, safeErr)
				require.NoError(t, c.MigrateTables(ctx, message.WriteMigrateTables{{Table: tc.pair.New, MigrateForce: true}}))
			} else {
				require.NoError(t, safeErr)
			}

			count, err := collection.CountDocuments(ctx, bson.D{})
			require.NoError(t, err)
			require.EqualValues(t, seededDocuments, count)
			if newKeys := tc.pair.New.PrimaryKeys(); len(newKeys) > 0 {
				requireUniquePrimaryKeyIndex(t, collection, newKeys)
			}
		})
	}
}

const seededDocuments = 2

func seedDocuments(t *testing.T, collection *mongo.Collection, table *schema.Table) {
	t.Helper()
	documents := make([]any, seededDocuments)
	for i := range documents {
		document := make(bson.D, 0, len(table.Columns))
		for _, column := range table.Columns {
			document = append(document, bson.E{Key: column.Name, Value: fmt.Sprintf("%s-%d", column.Name, i)})
		}
		documents[i] = document
	}
	_, err := collection.InsertMany(context.Background(), documents)
	require.NoError(t, err)
}

func requireUniquePrimaryKeyIndex(t *testing.T, collection *mongo.Collection, keys []string) {
	t.Helper()
	specs, err := collection.Indexes().ListSpecifications(context.Background())
	require.NoError(t, err)
	for _, index := range specs {
		if index.Name != "cq_pk" {
			continue
		}
		elements, err := index.KeysDocument.Elements()
		require.NoError(t, err)
		indexKeys := make([]string, len(elements))
		for i, element := range elements {
			indexKeys[i] = element.Key()
		}
		require.Equal(t, keys, indexKeys)
		require.NotNil(t, index.Unique)
		require.True(t, *index.Unique)
		return
	}
	require.Fail(t, "cq_pk index not found")
}
