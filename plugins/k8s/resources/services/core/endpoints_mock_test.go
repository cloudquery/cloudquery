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

func createCoreEndpoints(t *testing.T, ctrl *gomock.Controller) client.Services {
	endpoints := mocks.NewMockEndpointsClient(ctrl)
	e := corev1.Endpoints{}
	if err := faker.FakeDataSkipFields(&e, []string{"Subsets"}); err != nil {
		t.Fatal(err)
	}
	e.ManagedFields = []metav1.ManagedFieldsEntry{testData.FakeManagedFields(t)}
	subset := corev1.EndpointSubset{}
	if err := faker.FakeDataSkipFields(&subset, []string{"Addresses", "NotReadyAddresses"}); err != nil {
		t.Fatal(err)
	}
	address := corev1.EndpointAddress{}
	if err := faker.FakeDataSkipFields(&address, []string{"IP"}); err != nil {
		t.Fatal(err)
	}
	address.IP = "127.0.0.1"
	subset.Addresses = []corev1.EndpointAddress{address}
	subset.NotReadyAddresses = []corev1.EndpointAddress{address}
	e.Subsets = []corev1.EndpointSubset{subset}
	endpoints.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&corev1.EndpointsList{Items: []corev1.Endpoints{e}}, nil,
	)
	return client.Services{
		Endpoints: endpoints,
	}
}

func TestCoreEndpoints(t *testing.T) {
	client.K8sMockTestHelper(t, Endpoints(), createCoreEndpoints, client.TestOptions{})
}
