package core

import (
	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/mocks"
	resourcemock "github.com/cloudquery/cloudquery/plugins/source/k8s/mocks/core/v1"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	resource "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"testing"
)

func createPodTemplates(t *testing.T, ctrl *gomock.Controller) kubernetes.Interface {
	r := resource.PodTemplate{}
	if err := faker.FakeObject(&r); err != nil {
		t.Fatal(err)
	}

	r.Template = resource.PodTemplateSpec{}

	resourceClient := resourcemock.NewMockPodTemplateInterface(ctrl)
	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&resource.PodTemplateList{Items: []resource.PodTemplate{r}}, nil,
	)

	serviceClient := resourcemock.NewMockCoreV1Interface(ctrl)
	serviceClient.EXPECT().PodTemplates("").Return(resourceClient)

	cl := mocks.NewMockInterface(ctrl)
	cl.EXPECT().CoreV1().Return(serviceClient)

	return cl
}

func TestPodTemplates(t *testing.T) {
	client.K8sMockTestHelper(t, PodTemplates(), createPodTemplates)
}
