package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/iam/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/mitchellh/mapstructure"
)

func fetchIamAccounts(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client).Services().Iam

	summary, err := svc.GetAccountSummary(ctx, &iam.GetAccountSummaryInput{})
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
	config := iam.ListAccountAliasesInput{}
	for {
		response, err := svc.ListAccountAliases(ctx, &config)
		if err != nil {
			return err
		}

		accSummary.Aliases = append(accSummary.Aliases, response.AccountAliases...)

		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}
	res <- accSummary
	return nil
}
