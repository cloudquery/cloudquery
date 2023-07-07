package ssoadmin

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func customerManagedPolicies() *schema.Table {
	tableName := "aws_ssoadmin_permission_set_customer_managed_policies"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_ListManagedPoliciesInPermissionSet.html`,
		Resolver:    fetchCustomerManagedPolicies,
		Transform:   transformers.TransformWithStruct(&types.CustomerManagedPolicyReference{}, transformers.WithPrimaryKeys("Name", "Path")),
		Columns: []schema.Column{
			{
				Name:       "permission_set_arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("permission_set_arn"),
				PrimaryKey: true,
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

func fetchCustomerManagedPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Ssoadmin

	permissionSetARN := parent.Item.(*types.PermissionSet).PermissionSetArn
	instanceARN := parent.Parent.Item.(types.InstanceMetadata).InstanceArn
	config := ssoadmin.ListCustomerManagedPolicyReferencesInPermissionSetInput{
		InstanceArn:      instanceARN,
		PermissionSetArn: permissionSetARN,
	}
	paginator := ssoadmin.NewListCustomerManagedPolicyReferencesInPermissionSetPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(o *ssoadmin.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.CustomerManagedPolicyReferences
	}
	return nil
}
