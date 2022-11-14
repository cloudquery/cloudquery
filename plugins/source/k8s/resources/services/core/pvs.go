// Code generated by codegen; DO NOT EDIT.

package core

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/plugin-sdk/schema"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Pvs() *schema.Table {
	return &schema.Table{
		Name:      "k8s_core_pvs",
		Resolver:  fetchPvs,
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
				Name:     "spec_capacity",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Spec.Capacity"),
			},
			{
				Name:     "spec_persistent_volume_source",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Spec.PersistentVolumeSource"),
			},
			{
				Name:     "spec_access_modes",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Spec.AccessModes"),
			},
			{
				Name:     "spec_claim_ref",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Spec.ClaimRef"),
			},
			{
				Name:     "spec_persistent_volume_reclaim_policy",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.PersistentVolumeReclaimPolicy"),
			},
			{
				Name:     "spec_storage_class_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.StorageClassName"),
			},
			{
				Name:     "spec_mount_options",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Spec.MountOptions"),
			},
			{
				Name:     "spec_volume_mode",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.VolumeMode"),
			},
			{
				Name:     "spec_node_affinity",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Spec.NodeAffinity"),
			},
			{
				Name:     "status_phase",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status.Phase"),
			},
			{
				Name:     "status_message",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status.Message"),
			},
			{
				Name:     "status_reason",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status.Reason"),
			},
		},
	}
}

func fetchPvs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {

	cl := meta.(*client.Client).Client().CoreV1().PersistentVolumes()

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
