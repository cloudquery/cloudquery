package secretsmanager

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchSecretsmanagerSecrets(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().SecretsManager
	cfg := secretsmanager.ListSecretsInput{}
	for {
		response, err := svc.ListSecrets(ctx, &cfg)
		if err != nil {
			return err
		}

		var secrets []*secretsmanager.DescribeSecretOutput

		// get more details about the secret
		for _, n := range response.SecretList {
			cfg := secretsmanager.DescribeSecretInput{
				SecretId: n.ARN,
			}
			resp, err := svc.DescribeSecret(ctx, &cfg, func(options *secretsmanager.Options) {
				options.Region = c.Region
			})
			if err != nil {
				return err
			}

			secrets = append(secrets, resp)
		}

		res <- secrets

		if aws.ToString(response.NextToken) == "" {
			break
		}
		cfg.NextToken = response.NextToken
	}
	return nil
}

func fetchSecretsmanagerSecretPolicy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*secretsmanager.DescribeSecretOutput)
	cl := meta.(*client.Client)
	svc := cl.Services().SecretsManager
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

	v := map[string]interface{}{}
	err = json.Unmarshal([]byte(*response.ResourcePolicy), &v)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, v)
}
