package ssoadmin

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func getSsoadminPermissionSetInlinePolicy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	svc := meta.(*client.Client).Services().Ssoadmin
	permissionSetARN := resource.Item.(*types.PermissionSet).PermissionSetArn
	instanceARN := resource.Parent.Item.(types.InstanceMetadata).InstanceArn
	config := ssoadmin.GetInlinePolicyForPermissionSetInput{
		InstanceArn:      instanceARN,
		PermissionSetArn: permissionSetARN,
	}

	response, err := svc.GetInlinePolicyForPermissionSet(ctx, &config)
	if err != nil {
		return err
	}

	return resource.Set(c.Name, response.InlinePolicy)
}

func getSsoadminPermissionSet(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	svc := meta.(*client.Client).Services().Ssoadmin
	permission_set_arn := resource.Item.(string)
	instance_arn := resource.Parent.Item.(types.InstanceMetadata).InstanceArn
	config := ssoadmin.DescribePermissionSetInput{
		InstanceArn:      instance_arn,
		PermissionSetArn: &permission_set_arn,
	}

	response, err := svc.DescribePermissionSet(ctx, &config)
	if err != nil {
		return err
	}
	resource.Item = response.PermissionSet
	return nil
}

func fetchSsoadminPermissionSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Ssoadmin
	instance_arn := parent.Item.(types.InstanceMetadata).InstanceArn
	config := ssoadmin.ListPermissionSetsInput{
		InstanceArn: instance_arn,
	}

	for {
		response, err := svc.ListPermissionSets(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.PermissionSets
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}

	return nil
}
