package core

import (
	"testing"

	resource "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/mocks"
	resourcemock "github.com/cloudquery/cloudquery/plugins/source/k8s/mocks/core/v1"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
)

func createPodTemplates(t *testing.T, ctrl *gomock.Controller) client.Services {
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
	serviceClient.EXPECT().PodTemplates(metav1.NamespaceAll).Return(resourceClient)

	cl := mocks.NewMockInterface(ctrl)
	cl.EXPECT().CoreV1().Return(serviceClient)

	return client.Services{CoreAPI: cl}
}

func TestPodTemplates(t *testing.T) {
	client.MockTestHelper(t, PodTemplates(), createPodTemplates)
}
