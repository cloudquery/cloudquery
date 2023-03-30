package account

import (
	"github.com/aws/aws-sdk-go-v2/service/account/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func AlternateContacts() *schema.Table {
	tableName := "aws_account_alternate_contacts"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/accounts/latest/reference/API_AlternateContact.html`,
		Resolver:    fetchAccountAlternateContacts,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "account"),
		Transform:   client.TransformWithStruct(&types.AlternateContact{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:     "alternate_contact_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AlternateContactType"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
