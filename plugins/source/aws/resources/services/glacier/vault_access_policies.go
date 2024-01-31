package glacier

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/service/glacier"
	"github.com/aws/aws-sdk-go-v2/service/glacier/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"
)

func vaultAccessPolicies() *schema.Table {
	tableName := "aws_glacier_vault_access_policies"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/amazonglacier/latest/dev/api-GetVaultAccessPolicy.html`,
		Resolver:    fetchGlacierVaultAccessPolicies,
		Transform:   transformers.TransformWithStruct(&types.VaultAccessPolicy{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "vault_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("arn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "policy",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("Policy"),
			},
		},
	}
}

func fetchGlacierVaultAccessPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceGlacier).Glacier
	p := parent.Item.(types.DescribeVaultOutput)

	response, err := svc.GetVaultAccessPolicy(ctx, &glacier.GetVaultAccessPolicyInput{
		VaultName: p.VaultName,
	}, func(options *glacier.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	res <- response.Policy
	return nil
}
