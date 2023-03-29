package crd

import (
	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
)

func CRDs() *schema.Table {
	return &schema.Table{
		Name:      "k8s_crds",
		Resolver:  fetchCRDs,
		Multiplex: client.ContextMultiplex,
		Transform: client.TransformWithStruct(new(v1beta1.CustomResourceDefinition), transformers.WithPrimaryKeys("UID")),
		Columns:   schema.ColumnList{client.ContextColumn},
	}

	return nil
}
