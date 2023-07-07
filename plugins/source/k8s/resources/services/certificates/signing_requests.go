package certificates

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	v1 "k8s.io/api/certificates/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func SigningRequests() *schema.Table {
	return &schema.Table{
		Name:      "k8s_certificates_signing_requests",
		Resolver:  fetchSigningRequests,
		Multiplex: client.ContextMultiplex,
		Transform: client.TransformWithStruct(&v1.CertificateSigningRequest{}, transformers.WithPrimaryKeys("UID")),
		Columns:   schema.ColumnList{client.ContextColumn},
	}
}

func fetchSigningRequests(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client).Client().CertificatesV1().CertificateSigningRequests()

	opts := metav1.ListOptions{}
	for {
		result, err := cl.List(ctx, opts)
		if err != nil {
			return err
		}
		res <- result.Items
		if result.GetContinue() == "" {
			return nil
		}
		opts.Continue = result.GetContinue()
	}
}
