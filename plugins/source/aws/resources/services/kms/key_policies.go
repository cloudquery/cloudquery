package kms

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

type KeyPolicy struct {
	Name   string
	Policy *string
}

func keyPolicies() *schema.Table {
	tableName := "aws_kms_key_policies"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/kms/latest/APIReference/API_GetKeyPolicy.html`,
		Resolver:    fetchKeyPolicies,
		Transform:   transformers.TransformWithStruct(&KeyPolicy{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "kms"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "key_arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("arn"),
				PrimaryKey: true,
			},
			{
				Name:       "name",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("Name"),
				PrimaryKey: true,
			},
			{
				Name:     "policy",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("Policy"),
			},
		},
	}
}

func fetchKeyPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Kms

	const policyName = "default"

	k := parent.Item.(*types.KeyMetadata)
	d, err := svc.GetKeyPolicy(ctx, &kms.GetKeyPolicyInput{
		KeyId:      k.Arn,
		PolicyName: aws.String(policyName),
	}, func(o *kms.Options) {
		o.Region = cl.Region
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
