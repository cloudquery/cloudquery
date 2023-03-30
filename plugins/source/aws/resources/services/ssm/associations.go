package ssm

import (
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Associations() *schema.Table {
	tableName := "aws_ssm_associations"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_Association.html`,
		Resolver:    fetchSsmAssociations,
		Transform:   transformers.TransformWithStruct(&types.Association{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ssm"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:     "association_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AssociationId"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
