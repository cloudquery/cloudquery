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

func createBatchJobs(t *testing.T, ctrl *gomock.Controller) client.Services {
	jobs := mocks.NewMockJobsClient(ctrl)
	j := batchv1.Job{}
	if err := faker.FakeObject(&j); err != nil {
		t.Fatal(err)
	}
	j.Spec.Template = k8sTesting.FakePodTemplateSpec(t)
	j.ManagedFields = []metav1.ManagedFieldsEntry{k8sTesting.FakeManagedFields(t)}
	j.Spec.Template.ManagedFields = []metav1.ManagedFieldsEntry{k8sTesting.FakeManagedFields(t)}
	jobs.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&batchv1.JobList{Items: []batchv1.Job{j}}, nil,
	)
	return client.Services{
		Jobs: jobs,
	}
}

func TestBatchJobs(t *testing.T) {
	client.K8sMockTestHelper(t, Jobs(), createBatchJobs, client.TestOptions{})
}
