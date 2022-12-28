package json

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"

	"github.com/cloudquery/plugin-sdk/schema"
)

const maxJsonSize = 1024 * 1024 * 20

func Read(f io.Reader, table *schema.Table, sourceName string, res chan<- []any) error {
	sourceNameIndex := table.Columns.Index(schema.CqSourceNameColumn.Name)
	if sourceNameIndex == -1 {
		return fmt.Errorf("could not find column %s in table %s", schema.CqSourceNameColumn.Name, table.Name)
	}
	scanner := bufio.NewScanner(f)
	scanner.Buffer(make([]byte, maxJsonSize), maxJsonSize)
	for scanner.Scan() {
		jsonObj := make(map[string]interface{}, len(table.Columns))
		row := scanner.Bytes()
		if err := json.Unmarshal(row, &jsonObj); err != nil {
			return err
		}
		if jsonObj[schema.CqSourceNameColumn.Name] != sourceName {
			continue
		}
		jsonArr := make([]interface{}, len(table.Columns))
		for i, col := range table.Columns {
			jsonArr[i] = jsonObj[col.Name]
		}
		res <- jsonArr
	}

	return scanner.Err()
}