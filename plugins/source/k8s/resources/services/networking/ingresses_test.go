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

func createIngresses(t *testing.T, ctrl *gomock.Controller) kubernetes.Interface {
	r := resource.Ingress{}
	if err := faker.FakeObject(&r); err != nil {
		t.Fatal(err)
	}

	resourceClient := resourcemock.NewMockIngressInterface(ctrl)
	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&resource.IngressList{Items: []resource.Ingress{r}}, nil,
	)

	serviceClient := resourcemock.NewMockNetworkingV1Interface(ctrl)

	serviceClient.EXPECT().Ingresses("").Return(resourceClient)

	cl := mocks.NewMockInterface(ctrl)
	cl.EXPECT().NetworkingV1().Return(serviceClient)

	return cl
}

func TestIngresses(t *testing.T) {
	client.K8sMockTestHelper(t, Ingresses(), createIngresses)
}
