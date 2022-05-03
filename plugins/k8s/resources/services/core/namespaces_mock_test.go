//go:build mock
// +build mock

package core

import (
	"testing"

	"github.com/cloudquery/cq-provider-k8s/client"
	"github.com/cloudquery/cq-provider-k8s/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func createCoreNamespace(t *testing.T, ctrl *gomock.Controller) client.Services {
	s := mocks.NewMockNamespacesClient(ctrl)
	var namespace corev1.Namespace
	if err := faker.FakeData(&namespace); err != nil {
		t.Fatal(err)
	}
	namespace.ManagedFields = []metav1.ManagedFieldsEntry{testing.FakeManagedFields(t)}
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
