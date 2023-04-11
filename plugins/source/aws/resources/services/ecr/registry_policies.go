package ecr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func RegistryPolicies() *schema.Table {
	tableName := "aws_ecr_registry_policies"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonECR/latest/APIReference/API_GetRegistryPolicy.html`,
		Resolver:    fetchEcrRegistryPolicies,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "api.ecr"),
		Transform:   transformers.TransformWithStruct(&ecr.GetRegistryPolicyOutput{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:     "registry_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RegistryId"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "policy_text",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PolicyText"),
			},
		},
	}
}
func fetchEcrRegistryPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Ecr
	output, err := svc.GetRegistryPolicy(ctx, &ecr.GetRegistryPolicyInput{})
	if err != nil {
		if client.IsAWSError(err, "RegistryPolicyNotFoundException") {
			return nil
		}
		return err
	}
	res <- output
	return nil
}
