package policy

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	policy "k8s.io/api/policy/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func PodDisruptionBudgets() *schema.Table {
	return &schema.Table{
		Name:      "k8s_policy_pod_disruption_budgets",
		Resolver:  fetchPodDisruptionBudgets,
		Multiplex: client.ContextNamespaceMultiplex,
		Transform: client.TransformWithStruct(&policy.PodDisruptionBudget{}, transformers.WithPrimaryKeys("UID")),
		Columns:   schema.ColumnList{client.ContextColumn},
	}
}

func fetchPodDisruptionBudgets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	cl := c.Client().PolicyV1().PodDisruptionBudgets(c.Namespace)

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
