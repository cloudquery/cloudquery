package account

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/account"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchAccountContacts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Account
	var input account.GetContactInformationInput
	output, err := svc.GetContactInformation(ctx, &input)
	if err != nil {
		return err
	}
	res <- output.ContactInformation
	return nil
}
