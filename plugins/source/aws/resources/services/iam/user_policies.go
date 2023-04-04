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

func userPolicies() *schema.Table {
	tableName := "aws_iam_user_policies"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetUserPolicy.html`,
		Resolver:            fetchIamUserPolicies,
		PreResourceResolver: getUserPolicy,
		Transform:           transformers.TransformWithStruct(&iam.GetUserPolicyOutput{}, transformers.WithPrimaryKeys("PolicyName")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:     "user_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "user_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("user_id"),
			},
			{
				Name:     "policy_document",
				Type:     schema.TypeJSON,
				Resolver: resolveIamUserPolicyPolicyDocument,
			},
		},
	}
}

func fetchIamUserPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Iam
	user := parent.Item.(*types.User)
	config := iam.ListUserPoliciesInput{UserName: user.UserName}
	paginator := iam.NewListUserPoliciesPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			if c.IsNotFoundError(err) {
				return nil
			}
			return err
		}
		res <- page.PolicyNames
	}
	return nil
}

func getUserPolicy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	svc := meta.(*client.Client).Services().Iam
	p := resource.Item.(string)
	user := resource.Parent.Item.(*types.User)

	policyResult, err := svc.GetUserPolicy(ctx, &iam.GetUserPolicyInput{PolicyName: &p, UserName: user.UserName})
	if err != nil {
		return err
	}

	resource.Item = policyResult
	return nil
}

func resolveIamUserPolicyPolicyDocument(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*iam.GetUserPolicyOutput)

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
