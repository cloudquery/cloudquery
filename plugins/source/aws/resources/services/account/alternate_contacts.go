package account

import (
	"github.com/aws/aws-sdk-go-v2/service/account/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func AlternateContacts() *schema.Table {
	return &schema.Table{
		Name:        "aws_account_alternate_contacts",
		Description: `https://docs.aws.amazon.com/accounts/latest/reference/API_AlternateContact.html`,
		Resolver:    fetchAccountAlternateContacts,
		Multiplex:   client.AccountMultiplex,
		Transform:   transformers.TransformWithStruct(&types.AlternateContact{}),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
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
