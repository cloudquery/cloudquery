package nodes

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/mocks"

	resourcemock "github.com/cloudquery/cloudquery/plugins/source/k8s/mocks/node/v1"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	resource "k8s.io/api/node/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func createRuntimeClasses(t *testing.T, ctrl *gomock.Controller) kubernetes.Interface {
	r := resource.RuntimeClass{}
	if err := faker.FakeObject(&r); err != nil {
		t.Fatal(err)
	}

	resourceClient := resourcemock.NewMockRuntimeClassInterface(ctrl)
	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&resource.RuntimeClassList{Items: []resource.RuntimeClass{r}}, nil,
	)

	serviceClient := resourcemock.NewMockNodeV1Interface(ctrl)

	serviceClient.EXPECT().RuntimeClasses().Return(resourceClient)

	cl := mocks.NewMockInterface(ctrl)
	cl.EXPECT().NodeV1().Return(serviceClient)

	return cl
}

func TestRuntimeClasses(t *testing.T) {
	client.K8sMockTestHelper(t, RuntimeClasses(), createRuntimeClasses)
}
