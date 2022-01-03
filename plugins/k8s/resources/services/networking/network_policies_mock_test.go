//go:build mock
// +build mock

package networking

import (
	"testing"

	"github.com/cloudquery/cq-provider-k8s/client"
	"github.com/cloudquery/cq-provider-k8s/client/mocks"
	"github.com/cloudquery/cq-provider-k8s/resources/services/testData"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func createNetworkingNetworkPolicies(t *testing.T, ctrl *gomock.Controller) client.Services {
	s := mocks.NewMockNetworkPoliciesClient(ctrl)
	var networkPolicy networkingv1.NetworkPolicy
	if err := faker.FakeData(&networkPolicy); err != nil {
		t.Fatal(err)
	}
	networkPolicy.ManagedFields = []metav1.ManagedFieldsEntry{testData.FakeManagedFields(t)}

	s.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&networkingv1.NetworkPolicyList{Items: []networkingv1.NetworkPolicy{networkPolicy}}, nil,
	)
	return client.Services{
		NetworkPolicies: s,
	}
}

func TestNetworkingNetworkPolicies(t *testing.T) {
	client.K8sMockTestHelper(t, NetworkPolicies(), createNetworkingNetworkPolicies, client.TestOptions{})
}
