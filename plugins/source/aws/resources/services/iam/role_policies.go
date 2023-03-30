package iam

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func rolePolicies() *schema.Table {
	tableName := "aws_iam_role_policies"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetRolePolicy.html`,
		Resolver:            fetchIamRolePolicies,
		PreResourceResolver: getRolePolicy,
		Transform:           transformers.TransformWithStruct(&iam.GetRolePolicyOutput{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "iam"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:     "role_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "policy_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PolicyName"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "policy_document",
				Type:     schema.TypeJSON,
				Resolver: resolveRolePoliciesPolicyDocument,
			},
		},
	}
}

func fetchIamRolePolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Iam
	role := parent.Item.(*types.Role)
	paginator := iam.NewListRolePoliciesPaginator(svc, &iam.ListRolePoliciesInput{
		RoleName: role.RoleName,
	})
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			if c.IsNotFoundError(err) {
				return nil
			}
			return err
		}
		res <- output.PolicyNames
	}
	return nil
}

func getRolePolicy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Iam
	p := resource.Item.(string)
	role := resource.Parent.Item.(*types.Role)

	policyResult, err := svc.GetRolePolicy(ctx, &iam.GetRolePolicyInput{PolicyName: &p, RoleName: role.RoleName})
	if err != nil {
		return err
	}
	resource.Item = policyResult
	return nil
}

func resolveRolePoliciesPolicyDocument(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*iam.GetRolePolicyOutput)

	decodedDocument, err := url.QueryUnescape(*r.PolicyDocument)
	if err != nil {
		return err
	}

	var document map[string]any
	err = json.Unmarshal([]byte(decodedDocument), &document)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, document)
}
