package resourcegroups

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/resourcegroups/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ResourceGroups() *schema.Table {
	return &schema.Table{
		Name:                "aws_resourcegroups_resource_groups",
		Description:         `https://docs.aws.amazon.com/ARG/latest/APIReference/API_GetGroupQuery.html`,
		Resolver:            fetchResourcegroupsResourceGroups,
		PreResourceResolver: getResourceGroup,
		Transform:           transformers.TransformWithStruct(&models.ResourceGroupWrapper{}, transformers.WithUnwrapAllEmbeddedStructs()),
		Multiplex:           client.ServiceAccountRegionMultiplexer("resource-groups"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GroupArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveResourcegroupsResourceGroupTags,
			},
		},
	}
}
