package secretsmanager

import (
	"context"
	"encoding/json"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ARN"),
				PrimaryKey: true,
			},
			{
				Name:        "policy",
				Type:        sdkTypes.ExtensionTypes.JSON,
				Resolver:    fetchSecretsmanagerSecretPolicy,
				Description: `A JSON-formatted string that describes the permissions that are associated with the attached secret.`,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
		},
		Relations: []*schema.Table{
			secretVersions(),
		},
	}
}

func fetchSecretsmanagerSecrets(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Secretsmanager
	paginator := secretsmanager.NewListSecretsPaginator(svc, &secretsmanager.ListSecretsInput{})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(o *secretsmanager.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.SecretList
	}
	return nil
}

func getSecret(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Secretsmanager
	n := resource.Item.(types.SecretListEntry)

	// get more details about the secret
	resp, err := svc.DescribeSecret(ctx, &secretsmanager.DescribeSecretInput{
		SecretId: n.ARN,
	}, func(o *secretsmanager.Options) {
		o.Region = cl.Region
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
	response, err := svc.GetResourcePolicy(ctx, &cfg, func(o *secretsmanager.Options) {
		o.Region = cl.Region
	})
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
