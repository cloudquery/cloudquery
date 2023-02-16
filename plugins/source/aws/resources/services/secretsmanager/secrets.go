package secretsmanager

import (
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Secrets() *schema.Table {
	return &schema.Table{
		Name:                "aws_secretsmanager_secrets",
		Description:         `https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_ListSecrets.html`,
		Resolver:            fetchSecretsmanagerSecrets,
		Transform:           transformers.TransformWithStruct(&secretsmanager.DescribeSecretOutput{}, transformers.WithSkipFields("ResultMetadata")),
		PreResourceResolver: getSecret,
		Multiplex:           client.ServiceAccountRegionMultiplexer("secretsmanager"),
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
