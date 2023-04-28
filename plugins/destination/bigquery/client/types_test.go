package client

import (
	"testing"

	"cloud.google.com/go/bigquery"
	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/plugin-sdk/v2/types"
)

func TestSchemaTypeToBigQueryType(t *testing.T) {
	cases := []struct {
		dataType     arrow.DataType
		want         bigquery.FieldType
		wantRepeated bool
	}{
		{dataType: arrow.FixedWidthTypes.Boolean, want: bigquery.BooleanFieldType, wantRepeated: false},
		{dataType: types.ExtensionTypes.JSON, want: bigquery.JSONFieldType, wantRepeated: false},
		{dataType: arrow.ListOf(arrow.BinaryTypes.String), want: bigquery.StringFieldType, wantRepeated: true},
		{dataType: arrow.ListOf(types.ExtensionTypes.JSON), want: bigquery.JSONFieldType, wantRepeated: true},
	}
	cl := &Client{}
	for _, c := range cases {
		got, gotRepeated := cl.SchemaTypeToBigQueryType(c.dataType)
		if got != c.want {
			t.Errorf("SchemaTypeToBigQueryType(%v) got %v, want %v", c.dataType, got, c.want)
		}
		if gotRepeated != c.wantRepeated {
			t.Errorf("SchemaTypeToBigQueryType(%v) got repeated=%v, want %v", c.dataType, gotRepeated, c.wantRepeated)
		}
	}
}
