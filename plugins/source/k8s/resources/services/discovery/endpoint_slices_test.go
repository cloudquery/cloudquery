package discovery

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/mocks"

	resource "k8s.io/api/discovery/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	resourcemock "github.com/cloudquery/cloudquery/plugins/source/k8s/mocks/discovery/v1"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
)

func createEndpointSlices(t *testing.T, ctrl *gomock.Controller) client.Services {
	r := resource.EndpointSlice{}
	if err := faker.FakeObject(&r); err != nil {
		t.Fatal(err)
	}

	resourceClient := resourcemock.NewMockEndpointSliceInterface(ctrl)
	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&resource.EndpointSliceList{Items: []resource.EndpointSlice{r}}, nil,
	)

	serviceClient := resourcemock.NewMockDiscoveryV1Interface(ctrl)

	serviceClient.EXPECT().EndpointSlices(metav1.NamespaceAll).Return(resourceClient)

	cl := mocks.NewMockInterface(ctrl)
	cl.EXPECT().DiscoveryV1().Return(serviceClient)

	return client.Services{CoreAPI: cl}
}

func TestEndpointSlices(t *testing.T) {
	client.MockTestHelper(t, EndpointSlices(), createEndpointSlices)
}
