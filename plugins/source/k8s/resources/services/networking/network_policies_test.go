package networking

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/mocks"

	resourcemock "github.com/cloudquery/cloudquery/plugins/source/k8s/mocks/networking/v1"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	resource "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func createNetworkPolicies(t *testing.T, ctrl *gomock.Controller) kubernetes.Interface {
	r := resource.NetworkPolicy{}
	if err := faker.FakeObject(&r); err != nil {
		t.Fatal(err)
	}

	r.Spec.Ingress = []resource.NetworkPolicyIngressRule{}
	r.Spec.Egress = []resource.NetworkPolicyEgressRule{}

	resourceClient := resourcemock.NewMockNetworkPolicyInterface(ctrl)
	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&resource.NetworkPolicyList{Items: []resource.NetworkPolicy{r}}, nil,
	)

	serviceClient := resourcemock.NewMockNetworkingV1Interface(ctrl)

	serviceClient.EXPECT().NetworkPolicies("").Return(resourceClient)

	cl := mocks.NewMockInterface(ctrl)
	cl.EXPECT().NetworkingV1().Return(serviceClient)

	return cl
}

func TestNetworkPolicies(t *testing.T) {
	client.K8sMockTestHelper(t, NetworkPolicies(), createNetworkPolicies)
}
