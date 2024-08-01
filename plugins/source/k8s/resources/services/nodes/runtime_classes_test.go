package nodes

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/mocks"

	resource "k8s.io/api/node/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	resourcemock "github.com/cloudquery/cloudquery/plugins/source/k8s/mocks/node/v1"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
)

func createRuntimeClasses(t *testing.T, ctrl *gomock.Controller) client.Services {
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

	return client.Services{CoreAPI: cl}
}

func TestRuntimeClasses(t *testing.T) {
	client.MockTestHelper(t, RuntimeClasses(), createRuntimeClasses)
}
