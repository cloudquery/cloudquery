package recipes

import (
	"reflect"
	"strings"

	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/dms/models"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func DMSResources() []*Resource {
	resources := []*Resource{
		{
			SubService:  "replication_instances",
			Struct:      &models.ReplicationInstanceWrapper{},
			Description: "https://docs.aws.amazon.com/dms/latest/APIReference/API_ReplicationInstance.html",
			SkipFields:  []string{"ReplicationInstanceArn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ReplicationInstanceArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "dms"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("dms")`
		structName := reflect.ValueOf(r.Struct).Elem().Type().Name()
		if strings.Contains(structName, "Wrapper") {
			r.UnwrapEmbeddedStructs = true
		}
	}
	return resources
}
