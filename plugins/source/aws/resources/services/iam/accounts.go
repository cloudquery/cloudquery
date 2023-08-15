package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/iam/models"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/mitchellh/mapstructure"
)

func Accounts() *schema.Table {
	tableName := "aws_iam_accounts"
	return &schema.Table{
		Name:        tableName,
		Description: "https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetAccountSummary.html",
		Resolver:    fetchIamAccounts,
		Transform:   transformers.TransformWithStruct(&models.Account{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "iam"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
		},
	}
}

func fetchIamAccounts(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Iam

	summary, err := svc.GetAccountSummary(ctx, &iam.GetAccountSummaryInput{}, func(options *iam.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	var accSummary models.Account
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{WeaklyTypedInput: true, Result: &accSummary})
	if err != nil {
		return err
	}
	if err := decoder.Decode(summary.SummaryMap); err != nil {
		return err
	}
	paginator := iam.NewListAccountAliasesPaginator(svc, &iam.ListAccountAliasesInput{})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *iam.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		accSummary.Aliases = append(accSummary.Aliases, page.AccountAliases...)
	}
	res <- accSummary
	return nil
}
