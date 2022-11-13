package batch

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/plugin-sdk/schema"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func fetchBatchCronJobs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	jobs := meta.(*client.Client).Services().Client.BatchV1().Jobs("")
	opts := metav1.ListOptions{}
	for {
		result, err := jobs.List(ctx, opts)
		if err != nil {
			return err
		}
		res <- result.Items
		next := result.GetContinue()
		if next == "" {
			return nil
		}
		opts.Continue = next
	}
}
