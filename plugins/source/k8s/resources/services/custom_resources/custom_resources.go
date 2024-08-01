package custom_resources

import (
	"context"

	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	kubernetesschema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func resources() *schema.Table {
	return &schema.Table{
		Name:      "k8s_custom_resources",
		Resolver:  fetchCustomResources,
		Multiplex: client.ContextMultiplex,
		Transform: client.TransformWithStruct(
			&customResource{},
			transformers.WithPrimaryKeys("UID"),
			client.WithMoreSkipFields("Generation", "ResourceVersion", "DeletionGracePeriodSeconds"),
		),
		Columns: schema.ColumnList{client.ContextColumn},
	}
}

type customResource struct {
	metav1.TypeMeta   `json:",inline"` //nolint:all
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              map[string]any `json:"spec,omitempty"`
	Status            map[string]any `json:"status,omitempty"`
}

func fetchCustomResources(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	p := parent.Item.(v1.CustomResourceDefinition)

	for _, version := range p.Spec.Versions {
		c.Logger().Debug().Msgf("Fetching resources for CRD %s/%s", p.Name, version.Name)

		gvk := kubernetesschema.GroupVersionResource{
			Group:    p.Spec.Group,
			Version:  version.Name,
			Resource: p.Spec.Names.Plural,
		}

		var resource dynamic.ResourceInterface
		if p.Spec.Scope == v1.NamespaceScoped {
			resource = c.DynamicAPI().Resource(gvk).Namespace(metav1.NamespaceAll)
		} else {
			resource = c.DynamicAPI().Resource(gvk)
		}

		opts := metav1.ListOptions{}
		for {
			rs, err := resource.List(ctx, opts)
			if err != nil {
				c.Logger().Warn().Msgf("Error listing resources for %s: %s", gvk.String(), err.Error())
				break
			}

			for _, item := range rs.Items {
				data, err := runtime.DefaultUnstructuredConverter.ToUnstructured(&item)
				if err != nil {
					return err
				}

				var cr customResource
				if err := runtime.DefaultUnstructuredConverter.FromUnstructured(data, &cr); err != nil {
					return err
				}

				res <- cr
			}

			if rs.GetContinue() == "" {
				break
			}
			opts.Continue = rs.GetContinue()
		}
	}

	return nil
}
