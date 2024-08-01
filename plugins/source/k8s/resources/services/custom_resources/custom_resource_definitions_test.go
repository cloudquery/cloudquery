package custom_resources

import (
	"testing"

	crd "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/mocks/apiextensions"
	apiextensionmocks "github.com/cloudquery/cloudquery/plugins/source/k8s/mocks/apiextensions/v1"
	dynamicmocks "github.com/cloudquery/cloudquery/plugins/source/k8s/mocks/dynamic"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
)

func createCRDs(t *testing.T, ctrl *gomock.Controller) client.Services {
	apiExtensionsCl := apiextensions.NewMockInterface(ctrl)
	dynamicCl := dynamicmocks.NewMockInterface(ctrl)

	var crd2 crd.CustomResourceDefinition
	if err := faker.FakeObject(&crd2); err != nil {
		t.Fatal(err)
	}

	customResourceDefinitionClient := apiextensionmocks.NewMockCustomResourceDefinitionInterface(ctrl)
	customResourceDefinitionClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&crd.CustomResourceDefinitionList{Items: []crd.CustomResourceDefinition{crd2}}, nil,
	)

	apiextensionsV1Client := apiextensionmocks.NewMockApiextensionsV1Interface(ctrl)
	apiextensionsV1Client.EXPECT().CustomResourceDefinitions().Return(customResourceDefinitionClient)
	apiExtensionsCl.EXPECT().ApiextensionsV1().Return(apiextensionsV1Client)

	var rs dynamicmocks.MockNamespaceableResourceInterface
	if err := faker.FakeObject(&rs); err != nil {
		t.Fatal(err)
	}

	var crl unstructured.UnstructuredList
	if err := faker.FakeObject(&crl); err != nil {
		t.Fatal(err)
	}

	crl.Items[0].SetUnstructuredContent(map[string]any{
		"spec": map[string]any{
			"test": "test",
		},
		"status": map[string]any{
			"test": "test",
		},
	})
	crl.Items[0].SetLabels(map[string]string{"test": "test"})
	crl.Items[0].SetAnnotations(map[string]string{"test": "test"})
	crl.Items[0].SetOwnerReferences([]metav1.OwnerReference{{Name: "test"}})
	crl.Items[0].SetFinalizers([]string{"test"})

	cr := dynamicmocks.NewMockNamespaceableResourceInterface(ctrl)
	cr.EXPECT().List(gomock.Any(), gomock.Any()).Return(&crl, nil)

	dynamicCl.EXPECT().Resource(gomock.Any()).Return(cr)

	return client.Services{APIExtensionsAPI: apiExtensionsCl, DynamicAPI: dynamicCl}
}

func TestCRDs(t *testing.T) {
	client.MockTestHelper(t, CRDs(), createCRDs)
}
