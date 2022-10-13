package core

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/client/mocks"
	k8sTesting "github.com/cloudquery/cloudquery/plugins/source/k8s/resources/services/testing"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func createCoreNamespace(t *testing.T, ctrl *gomock.Controller) client.Services {
	s := mocks.NewMockNamespacesClient(ctrl)
	var namespace corev1.Namespace
	if err := faker.FakeObject(&namespace); err != nil {
		t.Fatal(err)
	}
	namespace.ManagedFields = []metav1.ManagedFieldsEntry{k8sTesting.FakeManagedFields(t)}
	s.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&corev1.NamespaceList{Items: []corev1.Namespace{namespace}}, nil,
	)
	return client.Services{
		Namespaces: s,
	}
}

func TestCoreNamespaces(t *testing.T) {
	client.K8sMockTestHelper(t, Namespaces(), createCoreNamespace, client.TestOptions{})
}
