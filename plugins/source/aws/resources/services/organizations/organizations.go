package organizations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Organizations() *schema.Table {
	tableName := "aws_organizations"
	return &schema.Table{
		Name: tableName,
		Description: `https://docs.aws.amazon.com/organizations/latest/APIReference/API_Organization.html
The 'request_account_id' column is added to show from where the request was made.`,
		Resolver: fetchOrganizationsOrganizations,
		Transform: transformers.TransformWithStruct(
			&types.Organization{},
			transformers.WithSkipFields(
				"AvailablePolicyTypes", // deprecated and misleading field according to docs
			),
			transformers.WithPrimaryKeyComponents("Arn"),
		),
		Multiplex: client.ServiceAccountRegionMultiplexer(tableName, "organizations"),
		Columns: []schema.Column{
			client.RequestAccountIDColumn(true),
		},
	}
}
func fetchOrganizationsOrganizations(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceOrganizations).Organizations
	o, err := svc.DescribeOrganization(ctx, &organizations.DescribeOrganizationInput{}, func(options *organizations.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}

	res <- o.Organization
	return nil
}
