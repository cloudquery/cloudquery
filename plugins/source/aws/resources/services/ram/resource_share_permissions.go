package ram

import (
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ResourceSharePermissions() *schema.Table {
	return &schema.Table{
		Name:        "aws_ram_resource_share_permissions",
		Description: `https://docs.aws.amazon.com/ram/latest/APIReference/API_ResourceSharePermissionSummary.html`,
		Transform:   transformers.TransformWithStruct(&types.ResourceSharePermissionSummary{}, transformers.WithPrimaryKeys("Arn", "Version")),
		Resolver:    fetchRamResourceSharePermissions,
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:     "permission",
				Type:     schema.TypeJSON,
				Resolver: resolveResourceSharePermissionDetailPermission,
			},
		},
	}
}
