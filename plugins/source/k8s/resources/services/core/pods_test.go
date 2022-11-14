// Code generated by codegen; DO NOT EDIT.

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

func createPods(t *testing.T, ctrl *gomock.Controller) kubernetes.Interface {
	r := resource.Pod{}
	if err := faker.FakeObject(&r); err != nil {
		t.Fatal(err)
	}

	r.Status.HostIP = "8.8.8.8"
	r.Status.PodIP = "1.1.1.1"
	r.Status.PodIPs = []resource.PodIP{resource.PodIP{IP: "1.1.1.1"}}
	r.Spec.Containers = []resource.Container{resource.Container{Name: "test"}}
	r.Spec.InitContainers = []resource.Container{resource.Container{Name: "test"}}

	resourceClient := resourcemock.NewMockPodInterface(ctrl)
	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&resource.PodList{Items: []resource.Pod{r}}, nil,
	)

	serviceClient := resourcemock.NewMockCoreV1Interface(ctrl)

	serviceClient.EXPECT().Pods("").Return(resourceClient)

	client := mocks.NewMockInterface(ctrl)
	client.EXPECT().CoreV1().Return(serviceClient)

	return client
}

func TestPods(t *testing.T) {
	client.K8sMockTestHelper(t, Pods(), createPods)
}
