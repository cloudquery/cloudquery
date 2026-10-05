package client

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/destination/mongodb/v2/client/spec"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
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

			finding := assessOneTable(t, tc.pair)
			safeErr := c.MigrateTables(ctx, message.WriteMigrateTables{{Table: tc.pair.New}})
			if finding.Category == plugin.AssessCategoryManualMigrationRequired {
				require.Error(t, safeErr)
				require.NoError(t, c.MigrateTables(ctx, message.WriteMigrateTables{{Table: tc.pair.New, MigrateForce: true}}))
			} else {
				require.NoError(t, safeErr)
			}
		})
	}
}
