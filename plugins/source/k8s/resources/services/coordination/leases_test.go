package coordination

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/mocks"

	resourcemock "github.com/cloudquery/cloudquery/plugins/source/k8s/mocks/coordination/v1"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	resource "k8s.io/api/coordination/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func createLeases(t *testing.T, ctrl *gomock.Controller) kubernetes.Interface {
	r := resource.Lease{}
	if err := faker.FakeObject(&r); err != nil {
		t.Fatal(err)
	}

	resourceClient := resourcemock.NewMockLeaseInterface(ctrl)
	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&resource.LeaseList{Items: []resource.Lease{r}}, nil,
	)

	serviceClient := resourcemock.NewMockCoordinationV1Interface(ctrl)

	serviceClient.EXPECT().Leases("").Return(resourceClient)

	cl := mocks.NewMockInterface(ctrl)
	cl.EXPECT().CoordinationV1().Return(serviceClient)

	return cl
}

func TestLeases(t *testing.T) {
	client.K8sMockTestHelper(t, Leases(), createLeases)
}
