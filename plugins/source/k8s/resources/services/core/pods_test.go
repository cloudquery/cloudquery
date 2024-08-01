package core

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/mocks"

	resource "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	resourcemock "github.com/cloudquery/cloudquery/plugins/source/k8s/mocks/core/v1"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
)

func createPods(t *testing.T, ctrl *gomock.Controller) client.Services {
	r := resource.Pod{}
	if err := faker.FakeObject(&r); err != nil {
		t.Fatal(err)
	}

	r.Status.HostIP = "8.8.8.8"
	r.Status.PodIP = "1.1.1.1"
	r.Status.PodIPs = []resource.PodIP{{IP: "1.1.1.1"}}
	r.Spec.Containers = []resource.Container{{Name: "test"}}
	r.Spec.InitContainers = []resource.Container{{Name: "test"}}

	resourceClient := resourcemock.NewMockPodInterface(ctrl)
	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&resource.PodList{Items: []resource.Pod{r}}, nil,
	)

	serviceClient := resourcemock.NewMockCoreV1Interface(ctrl)

	serviceClient.EXPECT().Pods(metav1.NamespaceAll).Return(resourceClient)

	cl := mocks.NewMockInterface(ctrl)
	cl.EXPECT().CoreV1().Return(serviceClient)

	return client.Services{CoreAPI: cl}
}

func TestPods(t *testing.T) {
	client.MockTestHelper(t, Pods(), createPods)
}
