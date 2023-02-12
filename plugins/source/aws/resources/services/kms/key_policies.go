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
		Name:        "aws_kms_key_policies",
		Description: `https://docs.aws.amazon.com/kms/latest/APIReference/API_GetKeyPolicy.html`,
		Resolver:    fetchKeyPolicies,
		Transform:   transformers.TransformWithStruct(&KeyPolicy{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("kms"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
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

func fetchKeyPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Kms

	const policyName = "default"

	k := parent.Item.(*types.KeyMetadata)
	d, err := svc.GetKeyPolicy(ctx, &kms.GetKeyPolicyInput{
		KeyId:      k.Arn,
		PolicyName: aws.String(policyName),
	})
	if err != nil {
		return err
	}

	res <- KeyPolicy{
		Name:   policyName,
		Policy: d.Policy,
	}
	return nil
}
