package ssm

import (
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func InventorySchemas() *schema.Table {
	return &schema.Table{
		Name:        "aws_ssm_inventory_schemas",
		Description: `https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_InventoryItemSchema.html`,
		Resolver:    fetchSsmInventorySchemas,
		Transform:   transformers.TransformWithStruct(&types.InventoryItemSchema{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("ssm"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:     "type_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TypeName"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Version"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
