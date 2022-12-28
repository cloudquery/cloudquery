package csv

import (
	"encoding/csv"
	"io"

	"github.com/cloudquery/plugin-sdk/schema"
)


func WriteTableBatch(w io.Writer, table *schema.Table, resources [][]any) error {
	writer := csv.NewWriter(w)	
	for _, resource := range resources {
		record := make([]string, len(resource))
		for i, v := range resource {
			record[i] = v.(string)
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}
	writer.Flush()
	return nil
}


