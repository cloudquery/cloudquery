package crd

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/mocks/apiextensions"
	mocks "github.com/cloudquery/cloudquery/plugins/source/k8s/mocks/apiextensions/v1"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	resource "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	apiextensionsclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func createCRDs(t *testing.T, ctrl *gomock.Controller) apiextensionsclientset.Interface {
	var crd resource.CustomResourceDefinition
	if err := faker.FakeObject(&crd); err != nil {
		t.Fatal(err)
	}

	resourceClient := mocks.NewMockCustomResourceDefinitionInterface(ctrl)
	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&resource.CustomResourceDefinitionList{Items: []resource.CustomResourceDefinition{crd}}, nil,
	)

	serviceClient := mocks.NewMockApiextensionsV1Interface(ctrl)
	serviceClient.EXPECT().CustomResourceDefinitions().Return(resourceClient)

	cl := apiextensions.NewMockInterface(ctrl)
	cl.EXPECT().ApiextensionsV1().Return(serviceClient)
	return cl
}

func TestCRDs(t *testing.T) {
	client.APIExtensionsMockTestHelper(t, CRDs(), createCRDs)
}
