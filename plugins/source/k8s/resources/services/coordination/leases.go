// Code generated by codegen; DO NOT EDIT.

package coordination

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/plugin-sdk/schema"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Leases() *schema.Table {
	return &schema.Table{
		Name:      "k8s_coordination_leases",
		Resolver:  fetchLeases,
		Multiplex: client.ContextMultiplex,
		Columns: []schema.Column{
			{
				Name:     "context",
				Type:     schema.TypeString,
				Resolver: client.ResolveContext,
			},
			{
				Name:     "uid",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Kind"),
			},
			{
				Name:     "api_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("APIVersion"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "namespace",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Namespace"),
			},
			{
				Name:     "resource_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceVersion"),
			},
			{
				Name:     "generation",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Generation"),
			},
			{
				Name:     "deletion_grace_period_seconds",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("DeletionGracePeriodSeconds"),
			},
			{
				Name:     "labels",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Labels"),
			},
			{
				Name:     "annotations",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Annotations"),
			},
			{
				Name:     "owner_references",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("OwnerReferences"),
			},
			{
				Name:     "finalizers",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Finalizers"),
			},
			{
				Name:     "spec_holder_identity",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.HolderIdentity"),
			},
			{
				Name:     "spec_lease_duration_seconds",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Spec.LeaseDurationSeconds"),
			},
			{
				Name:     "spec_acquire_time",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Spec.AcquireTime"),
			},
			{
				Name:     "spec_renew_time",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Spec.RenewTime"),
			},
			{
				Name:     "spec_lease_transitions",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Spec.LeaseTransitions"),
			},
		},
	}
}

func fetchLeases(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {

	cl := meta.(*client.Client).Client().CoordinationV1().Leases("")

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
