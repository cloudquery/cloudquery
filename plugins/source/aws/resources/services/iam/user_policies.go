package iam

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/aws/aws-sdk-go-v2/aws"
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
		Transform:           transformers.TransformWithStruct(&iam.GetUserPolicyOutput{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "iam"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:     "user_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "user_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
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
	for {
		output, err := svc.ListUserPolicies(ctx, &config)
		if err != nil {
			if c.IsNotFoundError(err) {
				return nil
			}
			return err
		}
		res <- output.PolicyNames

		if aws.ToString(output.Marker) == "" {
			break
		}
		config.Marker = output.Marker
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
