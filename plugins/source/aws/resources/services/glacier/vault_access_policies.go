package glacier

import (
	"github.com/aws/aws-sdk-go-v2/service/glacier/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func VaultAccessPolicies() *schema.Table {
	return &schema.Table{
		Name:        "aws_glacier_vault_access_policies",
		Description: `https://docs.aws.amazon.com/amazonglacier/latest/dev/api-GetVaultAccessPolicy.html`,
		Resolver:    fetchGlacierVaultAccessPolicies,
		Transform:   transformers.TransformWithStruct(&types.VaultAccessPolicy{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("glacier"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "vault_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
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
