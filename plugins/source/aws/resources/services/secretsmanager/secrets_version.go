package secretsmanager

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
				Name:       "secret_arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("arn"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchSecretsmanagerSecretsVersions(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, res chan<- any) error {
	secret := resource.Item.(*secretsmanager.DescribeSecretOutput)
	cl := meta.(*client.Client)
	svc := cl.Services().Secretsmanager
	paginator := secretsmanager.NewListSecretVersionIdsPaginator(svc, &secretsmanager.ListSecretVersionIdsInput{
		SecretId:          secret.ARN,
		IncludeDeprecated: aws.Bool(true),
	})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(o *secretsmanager.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Versions
	}
	return nil
}
