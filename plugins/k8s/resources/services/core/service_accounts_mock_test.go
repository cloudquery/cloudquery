//go:build mock
// +build mock

package core

import (
	"testing"

	"github.com/cloudquery/cq-provider-k8s/client"
	"github.com/cloudquery/cq-provider-k8s/client/mocks"
	"github.com/cloudquery/cq-provider-k8s/resources/services/testData"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func createCoreServiceAccounts(t *testing.T, ctrl *gomock.Controller) client.Services {
	serviceAccounts := mocks.NewMockServiceAccountsClient(ctrl)
	e := corev1.ServiceAccount{}
	if err := faker.FakeData(&e); err != nil {
		t.Fatal(err)
	}
	e.ManagedFields = []metav1.ManagedFieldsEntry{testData.FakeManagedFields(t)}
	serviceAccounts.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&corev1.ServiceAccountList{Items: []corev1.ServiceAccount{e}}, nil,
	)
	return client.Services{
		ServiceAccounts: serviceAccounts,
	}
}

func TestCoreServiceAccounts(t *testing.T) {
	client.K8sMockTestHelper(t, ServiceAccounts(), createCoreServiceAccounts, client.TestOptions{})

}
