package client

import (
	apiextensionsclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
)

type Services struct {
	CoreAPI          kubernetes.Interface
	APIExtensionsAPI apiextensionsclientset.Interface
	DynamicAPI       dynamic.Interface
}
