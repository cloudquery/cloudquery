package account

import (
	"github.com/aws/aws-sdk-go-v2/service/account/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Contacts() *schema.Table {
	return &schema.Table{
		Name:        "aws_account_contacts",
		Description: `https://docs.aws.amazon.com/accounts/latest/reference/API_ContactInformation.html`,
		Resolver:    fetchAccountContacts,
		Multiplex:   client.ServiceAccountRegionMultiplexer("account"),
		Transform:   transformers.TransformWithStruct(&types.ContactInformation{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
		},
	}
}
