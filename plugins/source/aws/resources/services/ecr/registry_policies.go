package ecr

import (
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func RegistryPolicies() *schema.Table {
	return &schema.Table{
		Name:      "aws_ecr_registry_policies",
		Resolver:  fetchEcrRegistryPolicies,
		Multiplex: client.ServiceAccountRegionMultiplexer("api.ecr"),
		Transform: transformers.TransformWithStruct(&ecr.GetRegistryPolicyOutput{}),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
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
