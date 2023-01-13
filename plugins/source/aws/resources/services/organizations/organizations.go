package organizations

import (
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Organizations() *schema.Table {
	return &schema.Table{
		Name:        "aws_organizations",
		Description: `https://docs.aws.amazon.com/organizations/latest/APIReference/API_Organization.html`,
		Resolver:    fetchOrganizationsOrganizations,
		Transform: transformers.TransformWithStruct(
			&types.Organization{},
			transformers.WithSkipFields(
				"AvailablePolicyTypes", // deprecated and misleading field according to docs
			),
		),
		Multiplex: client.AccountMultiplex,
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
