package certificates

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/gandi/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchCertificates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	list, err := cl.Services.CertificateClient.ListCertificates()
	if err != nil {
		return err
	}
	res <- list
	return nil
}
