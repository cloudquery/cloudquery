package admissionregistration

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	v1 "k8s.io/api/admissionregistration/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func MutatingWebhookConfigurations() *schema.Table {
	return &schema.Table{
		Name:      "k8s_admissionregistration_mutating_webhook_configurations",
		Resolver:  fetchMutatingWebhookConfigurations,
		Multiplex: client.ContextMultiplex,
		Transform: client.TransformWithStruct(&v1.MutatingWebhookConfiguration{}, transformers.WithPrimaryKeys("UID")),
		Columns:   schema.ColumnList{client.ContextColumn},
	}
}

func fetchMutatingWebhookConfigurations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client).Client().AdmissionregistrationV1().MutatingWebhookConfigurations()

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
