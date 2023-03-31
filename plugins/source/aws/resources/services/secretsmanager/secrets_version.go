package secretsmanager

import (
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func secretVersions() *schema.Table {
	tableName := "aws_secretsmanager_secret_versions"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_ListSecretVersionIds.html`,
		Resolver:    fetchSecretsmanagerSecretsVersions,
		Transform:   transformers.TransformWithStruct(&types.SecretVersionsListEntry{}, transformers.WithPrimaryKeys("VersionId")),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "secretsmanager"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "secret_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
