package ram

import (
	"context"
	"strconv"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ram"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
				Name:       "resource_share_arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("arn"),
				PrimaryKey: true,
			},
			{
				Name:     "permission",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveResourceSharePermissionDetailPermission,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}

func fetchRamResourceSharePermissions(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	input := &ram.ListResourceSharePermissionsInput{
		MaxResults:       aws.Int32(500),
		ResourceShareArn: resource.Item.(types.ResourceShare).ResourceShareArn,
	}
	paginator := ram.NewListResourceSharePermissionsPaginator(meta.(*client.Client).Services().Ram, input)
	for paginator.HasMorePages() {
		response, err := paginator.NextPage(ctx, func(options *ram.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.Permissions
	}
	return nil
}

func resolveResourceSharePermissionDetailPermission(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Ram
	permission := resource.Item.(types.ResourceSharePermissionSummary)
	version, err := strconv.ParseInt(aws.ToString(permission.Version), 10, 32)
	if err != nil {
		return err
	}
	input := &ram.GetPermissionInput{
		PermissionArn:     permission.Arn,
		PermissionVersion: aws.Int32(int32(version)),
	}
	response, err := svc.GetPermission(ctx, input, func(options *ram.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, response.Permission.Permission)
}
