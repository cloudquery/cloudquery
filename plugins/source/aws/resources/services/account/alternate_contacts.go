package account

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/account"
	"github.com/aws/aws-sdk-go-v2/service/account/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func AlternateContacts() *schema.Table {
	tableName := "aws_account_alternate_contacts"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/accounts/latest/reference/API_AlternateContact.html`,
		Resolver:    fetchAccountAlternateContacts,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "account"),
		Transform:   transformers.TransformWithStruct(&types.AlternateContact{}),
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

func fetchAccountAlternateContacts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Account
	var contactTypes types.AlternateContactType
	for _, acType := range contactTypes.Values() {
		var input account.GetAlternateContactInput
		input.AlternateContactType = acType
		output, err := svc.GetAlternateContact(ctx, &input)
		if err != nil {
			if client.IsAWSError(err, "ResourceNotFoundException") {
				continue
			}
			return err
		}
		res <- output.AlternateContact
	}
	return nil
}
