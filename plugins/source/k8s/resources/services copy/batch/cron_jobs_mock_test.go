package batch

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/client/mocks"
	k8sTesting "github.com/cloudquery/cloudquery/plugins/source/k8s/resources/services/testing"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func createBatchCronJobs(t *testing.T, ctrl *gomock.Controller) client.Services {
	
	// cronJobs := mocks.NewMockCronJobsClient(ctrl)
	batchv1 := mocks.NewMockBatchV1Interface(ctrl)
	batchv1.EXPECT().CronJobs("").Return(mocks.NewMockCronJobsClient(ctrl))

	cronJobs.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&batchv1.CronJobList{Items: []batchv1.CronJob{fakeCronJob(t)}},
		nil,
	)
	return client.Services{
		CronJobs: cronJobs,
	}
}

func fakeCronJob(t *testing.T) batchv1.CronJob {
	var job batchv1.CronJob
	if err := faker.FakeObject(&job); err != nil {
		t.Fatal(err)
	}
	job.ManagedFields = []metav1.ManagedFieldsEntry{k8sTesting.FakeManagedFields(t)}
	job.Spec.JobTemplate.ManagedFields = []metav1.ManagedFieldsEntry{k8sTesting.FakeManagedFields(t)}
	job.Spec.JobTemplate.Spec.Template = k8sTesting.FakePodTemplateSpec(t)
	return job
}

func TestBatchCronJobs(t *testing.T) {
	client.K8sMockTestHelper(t, CronJobs(), createBatchCronJobs, client.TestOptions{})
}
