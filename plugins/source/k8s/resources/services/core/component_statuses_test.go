package core

import (
	"testing"

	client "github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	mocks "github.com/cloudquery/cloudquery/plugins/source/k8s/mocks"

	resourcemock "github.com/cloudquery/cloudquery/plugins/source/k8s/mocks/core/v1"
	// k8sTesting "github.com/cloudquery/cloudquery/plugins/source/k8s/resources/services/testing"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	resource "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func createComponentStatuses(t *testing.T, ctrl *gomock.Controller) kubernetes.Interface {
	r := resource.ComponentStatus{}
	if err := faker.FakeObject(&r); err != nil {
		t.Fatal(err)
	}

	resourceClient := resourcemock.NewMockComponentStatusInterface(ctrl)
	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&resource.ComponentStatusList{Items: []resource.ComponentStatus{r}}, nil,
	)

	serviceClient := resourcemock.NewMockCoreV1Interface(ctrl)

	serviceClient.EXPECT().ComponentStatuses().Return(resourceClient)

	client := mocks.NewMockInterface(ctrl)
	client.EXPECT().CoreV1().Return(serviceClient)

	return client
}

func TestComponentStatuses(t *testing.T) {
	client.K8sMockTestHelper(t, ComponentStatuses(), createComponentStatuses)
}
