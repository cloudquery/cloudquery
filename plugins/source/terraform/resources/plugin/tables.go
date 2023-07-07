package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/terraform/resources/services"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func getTables() schema.Tables {
	tables := schema.Tables{services.TFData()}
	if err := transformers.TransformTables(tables); err != nil {
		panic(err)
	}
	for _, t := range tables {
		schema.AddCqIDs(t)
	}
	return tables
}
