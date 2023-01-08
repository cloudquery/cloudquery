package storage

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	v1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CsiStorageCapacities() *schema.Table {
	return &schema.Table{
		Name:      "k8s_storage_csi_storage_capacities",
		Resolver:  fetchCsiStorageCapacities,
		Multiplex: client.ContextMultiplex,
		Transform: transformers.TransformWithStruct(&v1.CSIStorageCapacity{}, client.SharedTransformers()...),
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
		},
	}
}

func fetchCsiStorageCapacities(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {

	cl := meta.(*client.Client).Client().StorageV1().CSIStorageCapacities("")

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
