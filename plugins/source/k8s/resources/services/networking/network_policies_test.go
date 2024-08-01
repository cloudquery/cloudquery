package networking

import (
	"testing"

	"k8s.io/apimachinery/pkg/util/intstr"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/mocks"

	resource "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	resourcemock "github.com/cloudquery/cloudquery/plugins/source/k8s/mocks/networking/v1"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
)

func createNetworkPolicies(t *testing.T, ctrl *gomock.Controller) client.Services {
	r := resource.NetworkPolicy{}
	if err := faker.FakeObject(&r); err != nil {
		t.Fatal(err)
	}

	p := intstr.FromInt(80)
	r.Spec.Ingress[0].Ports[0].Port = &p
	r.Spec.Egress[0].Ports[0].Port = &p

	resourceClient := resourcemock.NewMockNetworkPolicyInterface(ctrl)
	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&resource.NetworkPolicyList{Items: []resource.NetworkPolicy{r}}, nil,
	)

	serviceClient := resourcemock.NewMockNetworkingV1Interface(ctrl)

	serviceClient.EXPECT().NetworkPolicies(metav1.NamespaceAll).Return(resourceClient)

	cl := mocks.NewMockInterface(ctrl)
	cl.EXPECT().NetworkingV1().Return(serviceClient)

	return client.Services{CoreAPI: cl}
}

func TestNetworkPolicies(t *testing.T) {
	client.MockTestHelper(t, NetworkPolicies(), createNetworkPolicies)
}
