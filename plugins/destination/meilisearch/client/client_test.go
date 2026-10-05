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

	oktaOld := &schema.Table{Name: "okta_policy_rules_" + uuid.NewString()[:8], Columns: schema.ColumnList{
		{Name: "id", Type: arrow.BinaryTypes.String, PrimaryKey: true},
	}}
	oktaNew := oktaOld.Copy(nil)
	oktaNew.Columns = append(oktaNew.Columns, schema.Column{Name: "policy_id", Type: arrow.BinaryTypes.String, PrimaryKey: true})

	for _, pair := range []plugin.TablePair{{Old: datadogOld, New: datadogNew}, {Old: oktaOld, New: oktaNew}} {
		t.Run(pair.TableName(), func(t *testing.T) {
			require.NoError(t, c.MigrateTables(ctx, message.WriteMigrateTables{{Table: pair.Old}}))
			before, err := c.indexes()
			require.NoError(t, err)

			require.NoError(t, c.MigrateTables(ctx, message.WriteMigrateTables{{Table: pair.New}}))
			after, err := c.indexes()
			require.NoError(t, err)
			require.Equal(t, before[pair.TableName()], after[pair.TableName()])
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
