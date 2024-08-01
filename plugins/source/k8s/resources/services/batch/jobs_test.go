package batch

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/mocks"

	resource "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	resourcemock "github.com/cloudquery/cloudquery/plugins/source/k8s/mocks/batch/v1"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
)

func createJobs(t *testing.T, ctrl *gomock.Controller) client.Services {
	r := resource.Job{}
	if err := faker.FakeObject(&r); err != nil {
		t.Fatal(err)
	}

	r.Spec.Template = corev1.PodTemplateSpec{}

	resourceClient := resourcemock.NewMockJobInterface(ctrl)
	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&resource.JobList{Items: []resource.Job{r}}, nil,
	)

	serviceClient := resourcemock.NewMockBatchV1Interface(ctrl)

	serviceClient.EXPECT().Jobs(metav1.NamespaceAll).Return(resourceClient)

	cl := mocks.NewMockInterface(ctrl)
	cl.EXPECT().BatchV1().Return(serviceClient)

	return client.Services{CoreAPI: cl}
}

func TestJobs(t *testing.T) {
	client.MockTestHelper(t, Jobs(), createJobs)
}
