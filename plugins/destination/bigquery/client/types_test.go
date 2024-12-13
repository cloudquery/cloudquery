package client

import (
	"testing"

	"cloud.google.com/go/bigquery"
	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/google/go-cmp/cmp"
)

func TestClient_ColumnToBigQuerySchema(t *testing.T) {
	cases := []struct {
		col  schema.Column
		want *bigquery.FieldSchema
	}{
		{col: schema.Column{Name: "int32", Description: "bar", NotNull: true, Type: arrow.PrimitiveTypes.Int32}, want: &bigquery.FieldSchema{Name: "int32", Description: "bar", Required: false, Type: bigquery.IntegerFieldType}},
		{col: schema.Column{Name: "bool", Description: "bar", NotNull: true, Type: arrow.FixedWidthTypes.Boolean}, want: &bigquery.FieldSchema{Name: "bool", Description: "bar", Required: false, Type: bigquery.BooleanFieldType}},
		{col: schema.Column{Name: "json", Description: "bar", NotNull: true, Type: types.ExtensionTypes.JSON}, want: &bigquery.FieldSchema{Name: "json", Description: "bar", Required: false, Type: bigquery.JSONFieldType}},
		{col: schema.Column{Name: "string_list", Description: "bar", NotNull: true, Type: arrow.ListOf(arrow.BinaryTypes.String)}, want: &bigquery.FieldSchema{Name: "string_list", Description: "bar", Required: false, Type: bigquery.StringFieldType, Repeated: true}},
		{col: schema.Column{Name: "json_list", Description: "bar", NotNull: true, Type: arrow.ListOf(types.ExtensionTypes.JSON)}, want: &bigquery.FieldSchema{Name: "json_list", Description: "bar", Required: false, Type: bigquery.JSONFieldType, Repeated: true}},
		{col: schema.Column{Name: "struct", Description: "bar", NotNull: true, Type: arrow.StructOf([]arrow.Field{{Name: "foo", Type: arrow.BinaryTypes.String}}...)}, want: &bigquery.FieldSchema{Name: "struct", Description: "bar", Required: false, Type: bigquery.RecordFieldType, Schema: bigquery.Schema{{Name: "foo", Type: bigquery.StringFieldType}}}},
	}
	cl := &Client{}
	for _, c := range cases {
		got := cl.ColumnToBigQuerySchema(c.col)
		if diff := cmp.Diff(got, c.want); diff != "" {
			t.Errorf("ColumnToBigQuerySchema(%v) mismatch (-got +want):\n%s", c.col, diff)
		}
	}
}

func TestClient_DataTypeToBigQueryType(t *testing.T) {
	cases := []struct {
		dataType     arrow.DataType
		want         bigquery.FieldType
		wantRepeated bool
	}{
		{dataType: arrow.FixedWidthTypes.Boolean, want: bigquery.BooleanFieldType},
		{dataType: types.ExtensionTypes.JSON, want: bigquery.JSONFieldType},
		{dataType: arrow.ListOf(arrow.BinaryTypes.String), want: bigquery.StringFieldType},
		{dataType: arrow.ListOf(types.ExtensionTypes.JSON), want: bigquery.JSONFieldType},
		{dataType: arrow.StructOf([]arrow.Field{{Name: "foo", Type: arrow.BinaryTypes.String}}...), want: bigquery.RecordFieldType},
	}
	cl := &Client{}
	for _, c := range cases {
		got := cl.DataTypeToBigQueryType(c.dataType)
		if got != c.want {
			t.Errorf("ColumnToBigQuerySchema(%v) type = %v, want %v", c.dataType, got, c.want)
		}
	}
}
