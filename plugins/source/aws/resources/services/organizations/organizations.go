package organizations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Organizations() *schema.Table {
	tableName := "aws_organizations"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/organizations/latest/APIReference/API_Organization.html`,
		Resolver:    fetchOrganizationsOrganizations,
		Transform: client.TransformWithStruct(&types.Organization{},
			transformers.WithSkipFields(
				"AvailablePolicyTypes", // deprecated and misleading field according to docs
			),
			transformers.WithPrimaryKeys("Arn"),
		),
		Multiplex: client.ServiceAccountRegionMultiplexer(tableName, "organizations"),
		Columns:   []schema.Column{client.DefaultAccountIDColumn(true)},
	}
}

func fetchOrganizationsOrganizations(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	o, err := c.Services().Organizations.DescribeOrganization(ctx, &organizations.DescribeOrganizationInput{})
	if err != nil {
		return err
	}

	res <- o.Organization
	return nil
}
