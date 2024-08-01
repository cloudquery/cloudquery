package core

import (
	"testing"

	"k8s.io/apimachinery/pkg/util/intstr"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/mocks"

	resource "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	resourcemock "github.com/cloudquery/cloudquery/plugins/source/k8s/mocks/core/v1"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
)

func createServices(t *testing.T, ctrl *gomock.Controller) client.Services {
	r := resource.Service{}
	if err := faker.FakeObject(&r); err != nil {
		t.Fatal(err)
	}

	r.Spec.ClusterIP = "8.8.8.8"
	r.Spec.ClusterIPs = []string{"1.1.1.1"}
	r.Spec.ExternalIPs = []string{"1.1.1.1"}
	r.Spec.LoadBalancerIP = "1.1.1.1"
	r.Spec.Ports[0].TargetPort = intstr.FromInt(80)

	resourceClient := resourcemock.NewMockServiceInterface(ctrl)
	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&resource.ServiceList{Items: []resource.Service{r}}, nil,
	)

	serviceClient := resourcemock.NewMockCoreV1Interface(ctrl)

	serviceClient.EXPECT().Services(metav1.NamespaceAll).Return(resourceClient)

	cl := mocks.NewMockInterface(ctrl)
	cl.EXPECT().CoreV1().Return(serviceClient)

	return client.Services{CoreAPI: cl}
}

func TestServices(t *testing.T) {
	client.MockTestHelper(t, Services(), createServices)
}
