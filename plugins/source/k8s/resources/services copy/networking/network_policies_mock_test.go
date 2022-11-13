package networking

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/client/mocks"
	k8sTesting "github.com/cloudquery/cloudquery/plugins/source/k8s/resources/services/testing"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func createNetworkingNetworkPolicies(t *testing.T, ctrl *gomock.Controller) client.Services {
	s := mocks.NewMockNetworkPoliciesClient(ctrl)
	var networkPolicy networkingv1.NetworkPolicy
	if err := faker.FakeObject(&networkPolicy); err != nil {
		t.Fatal(err)
	}
	networkPolicy.ManagedFields = []metav1.ManagedFieldsEntry{k8sTesting.FakeManagedFields(t)}
	intStr1 := intstr.FromInt(100)
	intStr2 := intstr.FromInt(200)
	networkPolicy.Spec.Ingress[0].Ports[0].Port = &intStr1
	networkPolicy.Spec.Egress[0].Ports[0].Port = &intStr2

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
