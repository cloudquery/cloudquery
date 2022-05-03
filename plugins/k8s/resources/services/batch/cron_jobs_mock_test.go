//go:build mock
// +build mock

package batch

import (
	"testing"

	"github.com/cloudquery/faker/v3"

	"github.com/cloudquery/cq-provider-k8s/client"
	"github.com/cloudquery/cq-provider-k8s/client/mocks"
	"github.com/golang/mock/gomock"
	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func createBatchCronJobs(t *testing.T, ctrl *gomock.Controller) client.Services {
	cronJobs := mocks.NewMockCronJobsClient(ctrl)

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
	if err := faker.FakeDataSkipFields(&job, []string{
		"Spec"}); err != nil {
		t.Fatal(err)
	}
	if err := faker.FakeDataSkipFields(&job.Spec, []string{
		"JobTemplate", "ConcurrencyPolicy"}); err != nil {
		t.Fatal(err)
	}
	if err := faker.FakeData(&job.Spec.JobTemplate.ObjectMeta); err != nil {
		t.Fatal(err)
	}
	job.ManagedFields = []metav1.ManagedFieldsEntry{testing.FakeManagedFields(t)}
	job.Spec.JobTemplate.ManagedFields = []metav1.ManagedFieldsEntry{testing.FakeManagedFields(t)}
	job.Spec.JobTemplate.Spec.Template = testing.FakePodTemplateSpec(t)
	return job
}

func TestBatchCronJobs(t *testing.T) {
	client.K8sMockTestHelper(t, CronJobs(), createBatchCronJobs, client.TestOptions{})

}
