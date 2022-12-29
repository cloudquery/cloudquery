package csv

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"

	"github.com/cloudquery/plugin-sdk/schema"
)


func Read(f io.Reader, table *schema.Table, sourceName string, res chan<- []any) error {
	reader := csv.NewReader(f)
	sourceNameIndex := table.Columns.Index(schema.CqSourceNameColumn.Name)
	if sourceNameIndex == -1 {
		return fmt.Errorf("could not find column %s in table %s", schema.CqSourceNameColumn.Name, table.Name)
	}
	for {
		record, err := reader.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}
		if record[sourceNameIndex] != sourceName {
			continue
		}
		values := make([]any, len(record))
		for i, v := range record {
			values[i] = v
		}
		res <- values
	}
	return nil
}