package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func AccountAuthorizationDetails() *schema.Table {
	tableName := "aws_iam_account_authorization_details"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetAccountAuthorizationDetails.html`,
		Resolver:    fetchIamAccountAuthorizationDetails,
		Transform:   transformers.TransformWithStruct(&iam.GetAccountAuthorizationDetailsOutput{}, transformers.WithSkipFields("Marker", "IsTruncated", "ResultMetadata")),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "iam"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
		},
	}
}

func fetchIamAccountAuthorizationDetails(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	config := iam.GetAccountAuthorizationDetailsInput{
		Filter: types.EntityType("").Values(),
	}
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceIam).Iam
	paginator := iam.NewGetAccountAuthorizationDetailsPaginator(svc, &config)
	aggregatedOutput := &iam.GetAccountAuthorizationDetailsOutput{}
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *iam.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}

		aggregatedOutput.GroupDetailList = append(aggregatedOutput.GroupDetailList, page.GroupDetailList...)
		aggregatedOutput.RoleDetailList = append(aggregatedOutput.RoleDetailList, page.RoleDetailList...)
		aggregatedOutput.UserDetailList = append(aggregatedOutput.UserDetailList, page.UserDetailList...)
		aggregatedOutput.Policies = append(aggregatedOutput.Policies, page.Policies...)
	}
	res <- aggregatedOutput
	return nil
}
