package ssm

import (
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Inventories() *schema.Table {
	tableName := "aws_ssm_inventories"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_InventoryResultEntity.html`,
		Resolver:    fetchSsmInventories,
		Transform:   client.TransformWithStruct(&types.InventoryResultEntity{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ssm"),
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
