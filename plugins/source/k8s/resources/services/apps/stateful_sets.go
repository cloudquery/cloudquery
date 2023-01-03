// Code generated by codegen; DO NOT EDIT.

package apps

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/plugin-sdk/schema"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func StatefulSets() *schema.Table {
	return &schema.Table{
		Name:      "k8s_apps_stateful_sets",
		Resolver:  fetchStatefulSets,
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
				Name:     "spec_replicas",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Spec.Replicas"),
			},
			{
				Name:     "spec_selector",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Spec.Selector"),
			},
			{
				Name:     "spec_template",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Spec.Template"),
			},
			{
				Name:     "spec_volume_claim_templates",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Spec.VolumeClaimTemplates"),
			},
			{
				Name:     "spec_service_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.ServiceName"),
			},
			{
				Name:     "spec_pod_management_policy",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.PodManagementPolicy"),
			},
			{
				Name:     "spec_update_strategy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Spec.UpdateStrategy"),
			},
			{
				Name:     "spec_revision_history_limit",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Spec.RevisionHistoryLimit"),
			},
			{
				Name:     "spec_min_ready_seconds",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Spec.MinReadySeconds"),
			},
			{
				Name:     "spec_persistent_volume_claim_retention_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Spec.PersistentVolumeClaimRetentionPolicy"),
			},
			{
				Name:     "spec_ordinals",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Spec.Ordinals"),
			},
			{
				Name:     "status_observed_generation",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Status.ObservedGeneration"),
			},
			{
				Name:     "status_replicas",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Status.Replicas"),
			},
			{
				Name:     "status_ready_replicas",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Status.ReadyReplicas"),
			},
			{
				Name:     "status_current_replicas",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Status.CurrentReplicas"),
			},
			{
				Name:     "status_updated_replicas",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Status.UpdatedReplicas"),
			},
			{
				Name:     "status_current_revision",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status.CurrentRevision"),
			},
			{
				Name:     "status_update_revision",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status.UpdateRevision"),
			},
			{
				Name:     "status_collision_count",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Status.CollisionCount"),
			},
			{
				Name:     "status_conditions",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Status.Conditions"),
			},
			{
				Name:     "status_available_replicas",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Status.AvailableReplicas"),
			},
		},
	}
}

func fetchStatefulSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {

	cl := meta.(*client.Client).Client().AppsV1().StatefulSets("")

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
