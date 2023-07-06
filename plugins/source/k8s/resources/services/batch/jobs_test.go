package batch

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/mocks"

	resourcemock "github.com/cloudquery/cloudquery/plugins/source/k8s/mocks/batch/v1"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	resource "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	corev1 "k8s.io/api/core/v1"
)

func createJobs(t *testing.T, ctrl *gomock.Controller) kubernetes.Interface {
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

	serviceClient.EXPECT().Jobs("").Return(resourceClient)

	cl := mocks.NewMockInterface(ctrl)
	cl.EXPECT().BatchV1().Return(serviceClient)

	return cl
}

func TestJobs(t *testing.T) {
	client.K8sMockTestHelper(t, Jobs(), createJobs)
}
