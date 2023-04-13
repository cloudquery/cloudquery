package secretsmanager

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func Secrets() *schema.Table {
	tableName := "aws_secretsmanager_secrets"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_ListSecrets.html`,
		Resolver:            fetchSecretsmanagerSecrets,
		Transform:           transformers.TransformWithStruct(&secretsmanager.DescribeSecretOutput{}, transformers.WithSkipFields("ResultMetadata")),
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
			secretVersions(),
		},
	}
}

func fetchSecretsmanagerSecrets(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client).Services().Secretsmanager
	paginator := secretsmanager.NewListSecretsPaginator(svc, &secretsmanager.ListSecretsInput{})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.SecretList
	}
	return nil
}

func getSecret(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Secretsmanager
	n := resource.Item.(types.SecretListEntry)

	// get more details about the secret
	resp, err := svc.DescribeSecret(ctx, &secretsmanager.DescribeSecretInput{
		SecretId: n.ARN,
	})
	if err != nil {
		return err
	}

	resource.Item = resp
	return nil
}

func fetchSecretsmanagerSecretPolicy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*secretsmanager.DescribeSecretOutput)
	cl := meta.(*client.Client)
	svc := cl.Services().Secretsmanager
	cfg := secretsmanager.GetResourcePolicyInput{
		SecretId: r.ARN,
	}
	response, err := svc.GetResourcePolicy(ctx, &cfg)
	if err != nil {
		return err
	}
	// guard against nil-pointer dereference
	if response.ResourcePolicy == nil {
		return nil
	}

	v := map[string]any{}
	err = json.Unmarshal([]byte(*response.ResourcePolicy), &v)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, v)
}
