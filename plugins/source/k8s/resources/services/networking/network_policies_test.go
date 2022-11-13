// Code generated by codegen; DO NOT EDIT.

package networking

import (
	"testing"

	client "github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	mocks "github.com/cloudquery/cloudquery/plugins/source/k8s/mocks"

	resourcemock "github.com/cloudquery/cloudquery/plugins/source/k8s/mocks/networking/v1"
	// k8sTesting "github.com/cloudquery/cloudquery/plugins/source/k8s/resources/services/testing"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	resource "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func createNetworkPolicies(t *testing.T, ctrl *gomock.Controller) kubernetes.Interface {
	r := resource.NetworkPolicy{}
	if err := faker.FakeObject(&r,
		faker.WithSkipFields("Port"),
		faker.WithSkipTypeFields(),
		faker.WithFieldsValue(
			map[string]interface{}{},
		)); err != nil {
		t.Fatal(err)
	}

	resourceClient := resourcemock.NewMockNetworkPolicyInterface(ctrl)
	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&resource.NetworkPolicyList{Items: []resource.NetworkPolicy{r}}, nil,
	)

	serviceClient := resourcemock.NewMockNetworkingV1Interface(ctrl)

	serviceClient.EXPECT().NetworkPolicies("").Return(resourceClient)

	client := mocks.NewMockInterface(ctrl)
	client.EXPECT().NetworkingV1().Return(serviceClient)

	return client
}

func TestNetworkPolicies(t *testing.T) {
	client.K8sMockTestHelper(t, NetworkPolicies(), createNetworkPolicies)
}
