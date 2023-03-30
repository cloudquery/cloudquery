package secretsmanager

import (
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Secrets() *schema.Table {
	tableName := "aws_secretsmanager_secrets"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_ListSecrets.html`,
		Resolver:            fetchSecretsmanagerSecrets,
		Transform:           client.TransformWithStruct(&secretsmanager.DescribeSecretOutput{}, transformers.WithSkipFields("ResultMetadata")),
		PreResourceResolver: getSecret,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "secretsmanager"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:        "policy",
				Type:        schema.TypeJSON,
				Resolver:    fetchSecretsmanagerSecretPolicy,
				Description: `A JSON-formatted string that describes the permissions that are associated with the attached secret.`,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},
		Relations: []*schema.Table{
			SecretVersions(),
		},
	}
}
