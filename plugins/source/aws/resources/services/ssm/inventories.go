package ssm

import (
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Inventories() *schema.Table {
	return &schema.Table{
		Name:        "aws_ssm_inventories",
		Description: `https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_InventoryResultEntity.html`,
		Resolver:    fetchSsmInventories,
		Transform:   transformers.TransformWithStruct(&types.InventoryResultEntity{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("ssm"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
