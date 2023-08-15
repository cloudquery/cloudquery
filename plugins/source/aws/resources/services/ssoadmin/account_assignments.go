package ssoadmin

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func accountAssignments() *schema.Table {
	tableName := "aws_ssoadmin_permission_set_account_assignments"
	return &schema.Table{
		Name: tableName,
		Description: `https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_AccountAssignment.html
The 'request_account_id' and 'request_region' columns are added to show the account_id and region of where the request was made from.`,
		Resolver:  fetchSsoadminAccountAssignments,
		Transform: transformers.TransformWithStruct(&types.AccountAssignment{}, transformers.WithPrimaryKeys("PermissionSetArn", "PrincipalId", "PrincipalType", "AccountId")),
		Columns: schema.ColumnList{
			{
				Name:     "request_account_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "request_region",
				Type:     arrow.BinaryTypes.String,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:       "instance_arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("instance_arn"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchSsoadminAccountAssignments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Ssoadmin
	configListAccountForPPS := ssoadmin.ListAccountsForProvisionedPermissionSetInput{
		InstanceArn:      parent.Parent.Item.(types.InstanceMetadata).InstanceArn,
		PermissionSetArn: parent.Item.(*types.PermissionSet).PermissionSetArn,
	}

	paginatorListAccountForPPS := ssoadmin.NewListAccountsForProvisionedPermissionSetPaginator(svc, &configListAccountForPPS)
	for paginatorListAccountForPPS.HasMorePages() {
		accounts, err := paginatorListAccountForPPS.NextPage(ctx, func(o *ssoadmin.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		for _, account := range accounts.AccountIds {
			configLAA := ssoadmin.ListAccountAssignmentsInput{
				AccountId:        aws.String(account),
				InstanceArn:      parent.Parent.Item.(types.InstanceMetadata).InstanceArn,
				PermissionSetArn: parent.Item.(*types.PermissionSet).PermissionSetArn,
			}
			paginator := ssoadmin.NewListAccountAssignmentsPaginator(svc, &configLAA)
			for paginator.HasMorePages() {
				page, err := paginator.NextPage(ctx, func(o *ssoadmin.Options) {
					o.Region = cl.Region
				})
				if err != nil {
					return err
				}
				res <- page.AccountAssignments
			}
		}
	}
	return nil
}
