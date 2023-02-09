package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/plugin"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
)

func main() {
	args := os.Args[1:]
	specReader, err := specs.NewSpecReader(args)
	if err != nil {
		log.Fatalf("failed to load spec(s) from %s. Error: %v", strings.Join(args, ", "), err)
	}
	aws := plugin.AWS()

	tables := make(schema.Tables, 0)
	ctx := context.Background()
	for _, s := range specReader.Sources {
		err := aws.Init(ctx, *s)
		if err != nil {
			log.Fatalf("failed to init aws: %v", err)
		}
		tables = append(tables, aws.GetDynamicTables()...)
	}
	fmt.Println(generateViewQuery(tables))
}

func generateViewQuery(tables schema.Tables) string {
	var sb strings.Builder
	sb.WriteString(`CREATE VIEW aws_resources AS (`)
	sb.WriteString("\n")
	include := func(t *schema.Table) bool {
		return t.Columns.Get("account_id") != nil && t.Columns.Get("arn") != nil
	}
	exclude := func(t *schema.Table) bool {
		return false
	}
	filteredTables := tables.FilterDfsFunc(include, exclude, true)
	for i, t := range filteredTables {
		if i != 0 {
			sb.WriteString("  UNION ALL\n")
		}
		region := "region"
		if t.Columns.Get("region") == nil {
			region = "''"
		}
		tags := "tags"
		if t.Columns.Get("tags") == nil {
			tags = "'{}'"
		}
		q := `    SELECT _cq_id, _cq_source_name, _cq_sync_time, '%s' as _cq_table, account_id, %s as region, arn, %s as tags FROM %s`
		sb.WriteString(fmt.Sprintf(q, t.Name, region, tags, t.Name))
		sb.WriteString("\n")
	}
	sb.WriteString(`)`)
	return sb.String()
}
