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
)

func createCronJobs(t *testing.T, ctrl *gomock.Controller) kubernetes.Interface {
	r := resource.CronJob{}
	if err := faker.FakeObject(&r); err != nil {
		t.Fatal(err)
	}

	r.Spec.JobTemplate = resource.JobTemplateSpec{}

	resourceClient := resourcemock.NewMockCronJobInterface(ctrl)
	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&resource.CronJobList{Items: []resource.CronJob{r}}, nil,
	)

	serviceClient := resourcemock.NewMockBatchV1Interface(ctrl)

	serviceClient.EXPECT().CronJobs("").Return(resourceClient)

	cl := mocks.NewMockInterface(ctrl)
	cl.EXPECT().BatchV1().Return(serviceClient)

	return cl
}

func TestCronJobs(t *testing.T) {
	client.K8sMockTestHelper(t, CronJobs(), createCronJobs)
}
