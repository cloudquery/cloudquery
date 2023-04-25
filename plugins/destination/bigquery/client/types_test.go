package client

import (
	"testing"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/plugin-sdk/v2/testdata"
)

func TestGetValueForBigQuery(t *testing.T) {
	for arrowType := range TypeMapping {
		f := arrow.Field{
			Name:     arrowType.Name(),
			Type:     arrowType,
			Nullable: true,
		}
		sc := arrow.NewSchema([]arrow.Field{f}, nil)
		opts := testdata.GenTestDataOptions{
			MaxRows: 1,
		}
		records := testdata.GenTestData(sc, opts)
		for _, r := range records {
			for i := 0; i < int(r.NumCols()); i++ {
				v := GetValueForBigQuery(r.Column(i), 0)
				if v == nil {
					t.Errorf("GetValueForBigQuery(%s) = nil", arrowType.Name())
				}
			}
		}
	}
}
