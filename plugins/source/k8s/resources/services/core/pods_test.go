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

func createPods(t *testing.T, ctrl *gomock.Controller) kubernetes.Interface {
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

	serviceClient.EXPECT().Pods("").Return(resourceClient)

	cl := mocks.NewMockInterface(ctrl)
	cl.EXPECT().CoreV1().Return(serviceClient)

	return cl
}

func TestPods(t *testing.T) {
	client.K8sMockTestHelper(t, Pods(), createPods)
}
