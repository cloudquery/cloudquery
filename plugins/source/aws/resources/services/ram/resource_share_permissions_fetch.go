package ram

import (
	"context"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ram"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchRamResourceSharePermissions(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	input := &ram.ListResourceSharePermissionsInput{MaxResults: aws.Int32(500)}
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

func resolveResourceSharePermissionDetailPermission(
	ctx context.Context,
	meta schema.ClientMeta,
	resource *schema.Resource,
	c schema.Column,
) error {
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
