package iam

import (
	"context"
	"encoding/json"
	"net/url"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func groupPolicies() *schema.Table {
	tableName := "aws_iam_group_policies"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetGroupPolicy.html`,
		Resolver:            fetchIamGroupPolicies,
		PreResourceResolver: getGroupPolicy,
		Transform:           transformers.TransformWithStruct(&iam.GetGroupPolicyOutput{}, transformers.WithPrimaryKeys("PolicyName")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:       "group_arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("arn"),
				PrimaryKey: true,
			},
			{
				Name:     "policy_document",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveIamGroupPolicyPolicyDocument,
			},
		},
	}
}

func fetchIamGroupPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Iam
	group := parent.Item.(types.Group)
	config := iam.ListGroupPoliciesInput{
		GroupName: group.GroupName,
	}
	paginator := iam.NewListGroupPoliciesPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *iam.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			if cl.IsNotFoundError(err) {
				return nil
			}
			return err
		}

		res <- page.PolicyNames
	}
	return nil
}

func getGroupPolicy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Iam
	p := resource.Item.(string)
	group := resource.Parent.Item.(types.Group)

	policyResult, err := svc.GetGroupPolicy(ctx, &iam.GetGroupPolicyInput{PolicyName: &p, GroupName: group.GroupName}, func(options *iam.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	resource.Item = policyResult
	return nil
}

func resolveIamGroupPolicyPolicyDocument(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*iam.GetGroupPolicyOutput)

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
