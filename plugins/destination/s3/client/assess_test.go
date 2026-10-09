package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/s3/v7/client/spec"
	"github.com/cloudquery/filetypes/v4"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func newNoConnectionClient(t *testing.T, s []byte) *Client {
	t.Helper()
	c, err := New(context.Background(), zerolog.Nop(), s, plugin.NewClientOptions{NoConnection: true})
	require.NoError(t, err)
	return c.(*Client)
}

func noConnectionSpec(t *testing.T, format filetypes.FormatType, endpoint string) []byte {
	t.Helper()
	b, err := json.Marshal(spec.Spec{
		Bucket:   bucket,
		Region:   region,
		Path:     "{{TABLE}}/{{SYNC_ID}}/{{UUID}}.{{FORMAT}}",
		Endpoint: endpoint,
		FileSpec: filetypes.FileSpec{Format: format},
	})
	require.NoError(t, err)
	return b
}

func assessTable(t *testing.T, format filetypes.FormatType, pair plugin.TablePair) plugin.TableFinding {
	t.Helper()
	c := newNoConnectionClient(t, noConnectionSpec(t, format, ""))
	findings, err := c.AssessTables(context.Background(), []plugin.TablePair{pair}, plugin.AssessOptions{})
	require.NoError(t, err)
	require.Len(t, findings, 1)
	return findings[0]
}

func cloudflareCertificatePacks() plugin.TablePair {
	oldTable := &schema.Table{Name: "cloudflare_certificate_packs", Columns: schema.ColumnList{
		schema.CqIDColumn,
		schema.CqParentIDColumn,
		{Name: "account_id", Type: arrow.BinaryTypes.String},
		{Name: "zone_id", Type: arrow.BinaryTypes.String},
		{Name: "id", Type: arrow.BinaryTypes.String, PrimaryKey: true},
		{Name: "primary_certificate", Type: arrow.BinaryTypes.String},
		{Name: "certificate_authority", Type: arrow.BinaryTypes.String},
		{Name: "type", Type: arrow.BinaryTypes.String},
		{Name: "hosts", Type: arrow.ListOf(arrow.BinaryTypes.String)},
		{Name: "status", Type: arrow.BinaryTypes.String},
		{Name: "certificates", Type: types.ExtensionTypes.JSON},
		{Name: "created_on", Type: arrow.BinaryTypes.String},
		{Name: "validity_days", Type: arrow.PrimitiveTypes.Float64},
		{Name: "validation_method", Type: arrow.BinaryTypes.String},
	}}
	newTable := &schema.Table{Name: "cloudflare_certificate_packs", Columns: schema.ColumnList{
		schema.CqIDColumn,
		schema.CqParentIDColumn,
		{Name: "account_id", Type: arrow.BinaryTypes.String},
		{Name: "zone_id", Type: arrow.BinaryTypes.String},
		{Name: "id", Type: arrow.BinaryTypes.String, PrimaryKey: true},
		{Name: "validity_days", Type: arrow.PrimitiveTypes.Float64},
		{Name: "certificates", Type: types.ExtensionTypes.JSON},
		{Name: "hosts", Type: arrow.ListOf(arrow.BinaryTypes.String)},
		{Name: "status", Type: arrow.BinaryTypes.String},
		{Name: "type", Type: arrow.BinaryTypes.String},
		{Name: "certificate_authority", Type: arrow.BinaryTypes.String},
		{Name: "cloudflare_branding", Type: arrow.FixedWidthTypes.Boolean},
		{Name: "primary_certificate", Type: arrow.BinaryTypes.String},
		{Name: "validation_errors", Type: types.ExtensionTypes.JSON},
		{Name: "validation_method", Type: arrow.BinaryTypes.String},
		{Name: "validation_records", Type: types.ExtensionTypes.JSON},
	}}
	return plugin.TablePair{Old: oldTable, New: newTable}
}

func datadogMonitorTags() plugin.TablePair {
	oldTable := &schema.Table{Name: "datadog_monitors", Columns: schema.ColumnList{
		{Name: "id", Type: arrow.PrimitiveTypes.Int64, PrimaryKey: true},
		{Name: "tags", Type: arrow.ListOf(arrow.BinaryTypes.String)},
	}}
	newTable := oldTable.Copy(nil)
	newTable.Columns[1].Type = types.ExtensionTypes.JSON
	return plugin.TablePair{Old: oldTable, New: newTable}
}

func TestNewWithNoConnectionMakesNoS3Calls(t *testing.T) {
	var requests atomic.Int64
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		requests.Add(1)
		w.WriteHeader(http.StatusForbidden)
	}))
	defer server.Close()

	c := newNoConnectionClient(t, noConnectionSpec(t, filetypes.FormatTypeParquet, server.URL))

	require.NotNil(t, c.Client)
	require.Nil(t, c.s3Client)
	require.Nil(t, c.transferManager)
	require.NoError(t, c.Close(context.Background()))
	require.Zero(t, requests.Load())
}

func TestAssessTablesWithoutSpec(t *testing.T) {
	for _, s := range []string{"", "null"} {
		c := newNoConnectionClient(t, []byte(s))
		findings, err := c.AssessTables(context.Background(), []plugin.TablePair{cloudflareCertificatePacks()}, plugin.AssessOptions{})
		require.NoError(t, err)
		require.Equal(t, []plugin.TableFinding{{
			TableName:                "cloudflare_certificate_packs",
			Category:                 plugin.AssessCategoryUnknown,
			IncompleteCoverageReason: noSpecReason,
		}}, findings)
		require.NoError(t, c.Close(context.Background()))
	}
}

func TestAssessTablesParquetSchemaChange(t *testing.T) {
	require.Equal(t, plugin.TableFinding{
		TableName: "cloudflare_certificate_packs",
		Category:  plugin.AssessCategoryFileSchemaChanged,
		Columns: []plugin.ColumnFinding{
			{ColumnName: "created_on", Category: plugin.AssessCategoryFileSchemaChanged, OldType: "optional byte_array (String)"},
			{ColumnName: "cloudflare_branding", Category: plugin.AssessCategoryAutomaticallyMigratable, NewType: "optional boolean"},
			{ColumnName: "validation_errors", Category: plugin.AssessCategoryAutomaticallyMigratable, NewType: "optional byte_array (String)"},
			{ColumnName: "validation_records", Category: plugin.AssessCategoryAutomaticallyMigratable, NewType: "optional byte_array (String)"},
		},
	}, assessTable(t, filetypes.FormatTypeParquet, cloudflareCertificatePacks()))
}

func TestAssessTablesJSONSameOutputKeepsTypeChangeEvidence(t *testing.T) {
	finding := assessTable(t, filetypes.FormatTypeJSON, datadogMonitorTags())

	require.Equal(t, plugin.AssessCategoryNoChange, finding.Category)
	require.Len(t, finding.Columns, 1)
	column := finding.Columns[0]
	require.Equal(t, "tags", column.ColumnName)
	require.Equal(t, plugin.AssessCategoryNoChange, column.Category)
	require.Equal(t, "list<item: utf8, nullable>", column.OldType)
	require.Equal(t, "json", column.NewType)
	require.NotEmpty(t, column.Evidence)
	require.Equal(t, plugin.Evidence{
		SyntheticValue: `["env:prod"]`,
		Before:         `{"tags":["env:prod"]}`,
		After:          `{"tags":["env:prod"]}`,
	}, column.Evidence[0])
	for _, evidence := range column.Evidence {
		require.Equal(t, evidence.Before, evidence.After)
	}
}

func TestAssessTablesCSVHeaderChange(t *testing.T) {
	finding := assessTable(t, filetypes.FormatTypeCSV, cloudflareCertificatePacks())

	require.Equal(t, plugin.AssessCategoryFileSchemaChanged, finding.Category)
	require.Len(t, finding.Evidence, 1)
	require.Contains(t, finding.Evidence[0].Before, ",created_on,")
	require.NotContains(t, finding.Evidence[0].After, ",created_on,")
}
