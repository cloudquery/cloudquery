package ecr

import (
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func RegistryPolicies() *schema.Table {
	return &schema.Table{
		Name:        "aws_ecr_registry_policies",
		Description: `https://docs.aws.amazon.com/AmazonECR/latest/APIReference/API_GetRegistryPolicy.html`,
		Resolver:    fetchEcrRegistryPolicies,
		Multiplex:   client.ServiceAccountRegionMultiplexer("api.ecr"),
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
