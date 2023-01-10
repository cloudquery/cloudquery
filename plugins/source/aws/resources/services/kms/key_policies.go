package kms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

type KeyPolicy struct {
	Name   string
	Policy *string
}

func KeyPolicies() *schema.Table {
	return &schema.Table{
		Name:                "aws_kms_key_policies",
		Description:         `https://docs.aws.amazon.com/kms/latest/APIReference/API_GetKeyPolicy.html`,
		Resolver:            fetchKeyPolicies,
		PreResourceResolver: getKeyPolicy,
		Transform:           transformers.TransformWithStruct(&KeyPolicy{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer("kms"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "key_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Policy"),
			},
		},
	}
}

type keyAndPolicyName struct {
	KeyId      string
	PolicyName string
}

func fetchKeyPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	k := parent.Item.(*types.KeyMetadata)
	config := kms.ListKeyPoliciesInput{
		KeyId: k.Arn,
		Limit: aws.Int32(1000),
	}

	c := meta.(*client.Client)
	svc := c.Services().Kms
	p := kms.NewListKeyPoliciesPaginator(svc, &config)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx)
		if err != nil {
			return err
		}
		for _, n := range response.PolicyNames {
			res <- keyAndPolicyName{
				KeyId:      aws.ToString(k.Arn),
				PolicyName: n,
			}
		}
	}
	return nil
}

func getKeyPolicy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Kms
	item := resource.Item.(keyAndPolicyName)

	d, err := svc.GetKeyPolicy(ctx, &kms.GetKeyPolicyInput{
		KeyId:      aws.String(item.KeyId),
		PolicyName: aws.String(item.PolicyName),
	})
	if err != nil {
		return err
	}
	resource.Item = KeyPolicy{
		Name:   item.PolicyName,
		Policy: d.Policy,
	}
	return nil
}
