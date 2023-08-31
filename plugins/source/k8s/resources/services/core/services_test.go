package core

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/mocks"

	resourcemock "github.com/cloudquery/cloudquery/plugins/source/k8s/mocks/core/v1"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	resource "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func createServices(t *testing.T, ctrl *gomock.Controller) kubernetes.Interface {
	r := resource.Service{}
	if err := faker.FakeObject(&r); err != nil {
		t.Fatal(err)
	}

	r.Spec.ClusterIP = "8.8.8.8"
	r.Spec.ClusterIPs = []string{"1.1.1.1"}
	r.Spec.ExternalIPs = []string{"1.1.1.1"}
	r.Spec.LoadBalancerIP = "1.1.1.1"
	r.Spec.Ports = []resource.ServicePort{}

	resourceClient := resourcemock.NewMockServiceInterface(ctrl)
	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&resource.ServiceList{Items: []resource.Service{r}}, nil,
	)

	serviceClient := resourcemock.NewMockCoreV1Interface(ctrl)

	serviceClient.EXPECT().Services("").Return(resourceClient)

	cl := mocks.NewMockInterface(ctrl)
	cl.EXPECT().CoreV1().Return(serviceClient)

	return cl
}

func TestServices(t *testing.T) {
	client.K8sMockTestHelper(t, Services(), createServices)
}
