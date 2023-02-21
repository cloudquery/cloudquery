package secretsmanager

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchSecretsmanagerSecretsVersions(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, res chan<- any) error {
	secret := resource.Item.(*secretsmanager.DescribeSecretOutput)
	c := meta.(*client.Client)
	svc := c.Services().Secretsmanager
	paginator := secretsmanager.NewListSecretVersionIdsPaginator(svc, &secretsmanager.ListSecretVersionIdsInput{
		SecretId:          secret.ARN,
		IncludeDeprecated: aws.Bool(true),
	})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.Versions
	}
	return nil
}
