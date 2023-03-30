package account

import (
	"github.com/aws/aws-sdk-go-v2/service/account/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Contacts() *schema.Table {
	tableName := "aws_account_contacts"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/accounts/latest/reference/API_ContactInformation.html`,
		Resolver:    fetchAccountContacts,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "account"),
		Transform:   client.TransformWithStruct(&types.ContactInformation{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
		},
	}
}
