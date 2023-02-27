package glacier

import (
	"github.com/aws/aws-sdk-go-v2/service/glacier"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func VaultLockPolicies() *schema.Table {
	return &schema.Table{
		Name:        "aws_glacier_vault_lock_policies",
		Description: `https://docs.aws.amazon.com/amazonglacier/latest/dev/api-GetVaultLock.html`,
		Resolver:    fetchGlacierVaultLockPolicies,
		Transform:   transformers.TransformWithStruct(&glacier.GetVaultLockOutput{}, transformers.WithSkipFields("ResultMetadata")),
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
