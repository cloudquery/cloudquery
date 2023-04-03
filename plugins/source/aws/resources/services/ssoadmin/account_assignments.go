package ssoadmin

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func accountAssignments() *schema.Table {
	tableName := "aws_ssoadmin_account_assignments"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_AccountAssignment.html`,
		Resolver:    fetchSsoadminAccountAssignments,
		Transform:   transformers.TransformWithStruct(&types.AccountAssignment{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "identitystore"),
	}
}

func fetchSsoadminAccountAssignments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Ssoadmin
	permission_set_arn := parent.Item.(*types.PermissionSet).PermissionSetArn
	instance_arn := parent.Parent.Item.(types.InstanceMetadata).InstanceArn
	config := ssoadmin.ListAccountAssignmentsInput{
		AccountId:        &cl.AccountID,
		InstanceArn:      instance_arn,
		PermissionSetArn: permission_set_arn,
	}
	// TODO: replace with paginator
	for {
		response, err := svc.ListAccountAssignments(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.AccountAssignments
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}

	return nil
}
