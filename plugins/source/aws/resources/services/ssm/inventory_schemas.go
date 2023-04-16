package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func InventorySchemas() *schema.Table {
	tableName := "aws_ssm_inventory_schemas"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_InventoryItemSchema.html`,
		Resolver:    fetchSsmInventorySchemas,
		Transform:   transformers.TransformWithStruct(&types.InventoryItemSchema{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ssm"),
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

func fetchSsmInventorySchemas(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Ssm

	paginator := ssm.NewGetInventorySchemaPaginator(svc, nil)
	for paginator.HasMorePages() {
		v, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- v.Schemas
	}
	return nil
}
