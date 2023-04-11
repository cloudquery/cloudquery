package ram

import (
	"context"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ram"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func resourceSharePermissions() *schema.Table {
	return &schema.Table{
		Name:        "aws_ram_resource_share_permissions",
		Description: `https://docs.aws.amazon.com/ram/latest/APIReference/API_ResourceSharePermissionSummary.html`,
		Transform:   transformers.TransformWithStruct(&types.ResourceSharePermissionSummary{}, transformers.WithPrimaryKeys("Arn", "Version")),
		Resolver:    fetchRamResourceSharePermissions,
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:     "resource_share_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "permission",
				Type:     schema.TypeJSON,
				Resolver: resolveResourceSharePermissionDetailPermission,
			},
		},
	}
}

func fetchRamResourceSharePermissions(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, res chan<- any) error {
	input := &ram.ListResourceSharePermissionsInput{
		MaxResults:       aws.Int32(500),
		ResourceShareArn: resource.Item.(types.ResourceShare).ResourceShareArn,
	}
	paginator := ram.NewListResourceSharePermissionsPaginator(meta.(*client.Client).Services().Ram, input)
	for paginator.HasMorePages() {
		response, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- response.Permissions
	}
	return nil
}

func resolveResourceSharePermissionDetailPermission(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	permission := resource.Item.(types.ResourceSharePermissionSummary)
	version, err := strconv.ParseInt(aws.ToString(permission.Version), 10, 32)
	if err != nil {
		return err
	}
	input := &ram.GetPermissionInput{
		PermissionArn:     permission.Arn,
		PermissionVersion: aws.Int32(int32(version)),
	}
	response, err := meta.(*client.Client).Services().Ram.GetPermission(ctx, input)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, response.Permission.Permission)
}
