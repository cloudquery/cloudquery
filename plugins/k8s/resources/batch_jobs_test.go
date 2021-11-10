package resources

import (
	"github.com/cloudquery/cq-provider-k8s/client"
	"github.com/cloudquery/cq-provider-k8s/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func createBatchJobs(t *testing.T, ctrl *gomock.Controller) client.Services {
	jobs := mocks.NewMockJobsClient(ctrl)
	j := batchv1.Job{}
	if err := faker.FakeDataSkipFields(&j, []string{
		"Spec"}); err != nil {
		t.Fatal(err)
	}
	if err := faker.FakeDataSkipFields(&j.Spec, []string{
		"Template"}); err != nil {
		t.Fatal(err)
	}

	j.Spec.Template = fakePodTemplateSpec(t)
	j.ManagedFields = []metav1.ManagedFieldsEntry{fakeManagedFields(t)}
	j.Spec.Template.ManagedFields = []metav1.ManagedFieldsEntry{fakeManagedFields(t)}
	jobs.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&batchv1.JobList{Items: []batchv1.Job{j}}, nil,
	)
	return client.Services{
		Jobs: jobs,
	}
}

func TestBatchJobs(t *testing.T) {
	k8sTestHelper(t, BatchJobs(), createBatchJobs)
}
