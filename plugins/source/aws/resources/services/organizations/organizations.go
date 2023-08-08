package organizations

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Organizations() *schema.Table {
	tableName := "aws_organizations"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/organizations/latest/APIReference/API_Organization.html`,
		Resolver:    fetchOrganizationsOrganizations,
		Transform: transformers.TransformWithStruct(
			&types.Organization{},
			transformers.WithSkipFields(
				"AvailablePolicyTypes", // deprecated and misleading field according to docs
			),
			transformers.WithPrimaryKeys("Arn"),
		),
		Multiplex: client.ServiceAccountRegionMultiplexer(tableName, "organizations"),
		Columns: []schema.Column{
			{
				Name:       "request_account_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveAWSAccount,
				PrimaryKey: true,
			},
		},
	}
}

func fetchOrganizationsOrganizations(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Organizations
	o, err := svc.DescribeOrganization(ctx, &organizations.DescribeOrganizationInput{}, func(options *organizations.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}

	res <- o.Organization
	return nil
}
