package account

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/account"
	"github.com/aws/aws-sdk-go-v2/service/account/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Contacts() *schema.Table {
	tableName := "aws_account_contacts"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/accounts/latest/reference/API_ContactInformation.html`,
		Resolver:    fetchAccountContacts,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "account"),
		Transform:   transformers.TransformWithStruct(&types.ContactInformation{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
		},
	}
}

func fetchAccountContacts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Account
	var input account.GetContactInformationInput
	output, err := svc.GetContactInformation(ctx, &input, func(options *account.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	res <- output.ContactInformation
	return nil
}
