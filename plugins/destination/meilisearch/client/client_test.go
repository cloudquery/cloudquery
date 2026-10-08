package client

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	internalPlugin "github.com/cloudquery/cloudquery/plugins/destination/meilisearch/v2/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func getTestSpec() *Spec {
	apiKey := os.Getenv("CQ_DEST_MEILI_TEST_API_KEY")
	if len(apiKey) == 0 {
		apiKey = "test"
	}
	host := os.Getenv("CQ_DEST_MEILI_TEST_HOST")
	if len(host) == 0 {
		host = "http://localhost:7700"
	}

	return &Spec{Host: host, APIKey: apiKey}
}

func TestSafeMigrateMatchesAssessment(t *testing.T) {
	ctx := context.Background()
	specBytes, err := json.Marshal(getTestSpec())
	require.NoError(t, err)
	pluginClient, err := New(ctx, zerolog.Nop(), specBytes, plugin.NewClientOptions{})
	require.NoError(t, err)
	c := pluginClient.(*Client)
	t.Cleanup(func() { require.NoError(t, c.Close(ctx)) })

	datadogOld := &schema.Table{Name: "datadog_monitors_" + uuid.NewString()[:8], Columns: schema.ColumnList{
		{Name: "id", Type: arrow.PrimitiveTypes.Int64, PrimaryKey: true},
		{Name: "tags", Type: arrow.ListOf(arrow.BinaryTypes.String)},
	}}
	datadogNew := datadogOld.Copy(nil)
	datadogNew.Columns[1].Type = types.ExtensionTypes.JSON

	githubOld := &schema.Table{Name: "github_repositories_" + uuid.NewString()[:8], Columns: schema.ColumnList{
		{Name: "id", Type: arrow.PrimitiveTypes.Int64, PrimaryKey: true},
	}}
	githubNew := githubOld.Copy(nil)
	githubNew.Columns = append(githubNew.Columns, schema.Column{Name: "topics", Type: types.ExtensionTypes.JSON})

	oktaOld := &schema.Table{Name: "okta_policy_rules_" + uuid.NewString()[:8], Columns: schema.ColumnList{
		{Name: "id", Type: arrow.BinaryTypes.String, PrimaryKey: true},
	}}
	oktaNew := oktaOld.Copy(nil)
	oktaNew.Columns = append(oktaNew.Columns, schema.Column{Name: "policy_id", Type: arrow.BinaryTypes.String, PrimaryKey: true})

	tests := []struct {
		pair               plugin.TablePair
		category           plugin.AssessCategory
		row                map[string]any
		documentIDsChanged bool
	}{
		{
			pair:     plugin.TablePair{Old: datadogOld, New: datadogNew},
			category: plugin.AssessCategoryNoChange,
			row:      map[string]any{"id": 1},
		},
		{
			pair:     plugin.TablePair{Old: githubOld, New: githubNew},
			category: plugin.AssessCategoryAutomaticallyMigratable,
			row:      map[string]any{"id": 1},
		},
		{
			pair:               plugin.TablePair{Old: oktaOld, New: oktaNew},
			category:           plugin.AssessCategoryManualMigrationRequired,
			row:                map[string]any{"id": "rule", "policy_id": "policy"},
			documentIDsChanged: true,
		},
	}
	for _, tc := range tests {
		name := tc.pair.TableName()
		t.Run(name, func(t *testing.T) {
			findings, err := c.AssessTables(ctx, []plugin.TablePair{tc.pair}, plugin.AssessOptions{})
			require.NoError(t, err)
			require.Equal(t, tc.category, findings[0].Category)

			t.Cleanup(func() { require.NoError(t, c.deleteIndex(ctx, &indexSchema{UID: name})) })
			require.NoError(t, c.MigrateTables(ctx, message.WriteMigrateTables{{Table: tc.pair.Old}}))
			before, err := c.indexes()
			require.NoError(t, err)

			require.NoError(t, c.MigrateTables(ctx, message.WriteMigrateTables{{Table: tc.pair.New}}))
			after, err := c.indexes()
			require.NoError(t, err)
			require.Equal(t, before[name], after[name])

			filterable, err := c.Meilisearch.Index(name).GetFilterableAttributes()
			require.NoError(t, err)
			wantFilterable := make([]any, 0, len(tc.pair.New.Columns))
			for _, column := range tc.pair.New.Columns.Names() {
				wantFilterable = append(wantFilterable, column)
			}
			require.ElementsMatch(t, wantFilterable, *filterable)

			oldID, newID := hashUUID(tc.pair.Old)(tc.row), hashUUID(tc.pair.New)(tc.row)
			require.Equal(t, tc.documentIDsChanged, oldID != newID)
		})
	}
}

func TestPlugin(t *testing.T) {
	ctx := context.Background()
	p := plugin.NewPlugin("meilisearch", internalPlugin.Version, New)
	spec := getTestSpec()
	specBytes, err := json.Marshal(spec)
	if err != nil {
		t.Fatal(err)
	}
	if err := p.Init(ctx, specBytes, plugin.NewClientOptions{}); err != nil {
		t.Fatal(err)
	}

	plugin.TestWriterSuiteRunner(t,
		p,
		plugin.WriterTestSuiteTests{
			SkipDeleteStale:  true,
			SkipDeleteRecord: true,
			SkipMigrate:      true,
			SafeMigrations: plugin.SafeMigrations{
				AddColumn:           true,
				AddColumnNotNull:    true,
				RemoveColumn:        true,
				RemoveColumnNotNull: true,
			},
		},
	)
}
