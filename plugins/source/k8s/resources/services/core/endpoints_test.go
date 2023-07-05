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

func createEndpoints(t *testing.T, ctrl *gomock.Controller) kubernetes.Interface {
	r := resource.Endpoints{}
	if err := faker.FakeObject(&r); err != nil {
		t.Fatal(err)
	}

	resourceClient := resourcemock.NewMockEndpointsInterface(ctrl)
	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&resource.EndpointsList{Items: []resource.Endpoints{r}}, nil,
	)

	serviceClient := resourcemock.NewMockCoreV1Interface(ctrl)

	serviceClient.EXPECT().Endpoints("").Return(resourceClient)

	cl := mocks.NewMockInterface(ctrl)
	cl.EXPECT().CoreV1().Return(serviceClient)

	return cl
}

func TestEndpoints(t *testing.T) {
	client.K8sMockTestHelper(t, Endpoints(), createEndpoints)
}
