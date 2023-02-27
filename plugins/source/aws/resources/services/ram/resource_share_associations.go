package ram

import (
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ResourceShareAssociations() *schema.Table {
	return &schema.Table{
		Name:        "aws_ram_resource_share_associations",
		Description: `https://docs.aws.amazon.com/ram/latest/APIReference/API_ResourceShareAssociation.html`,
		Resolver:    fetchRamResourceShareAssociations,
		Transform:   transformers.TransformWithStruct(&types.ResourceShareAssociation{}, transformers.WithPrimaryKeys("AssociatedEntity", "ResourceShareArn")),
		Multiplex:   client.ServiceAccountRegionMultiplexer("ram"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
		},
	}
}
