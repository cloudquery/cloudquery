//go:build mock
// +build mock

package batch

import (
	"testing"

	"github.com/cloudquery/cq-provider-k8s/client"
	"github.com/cloudquery/cq-provider-k8s/client/mocks"
	"github.com/cloudquery/cq-provider-k8s/resources/services/testData"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

	j.Spec.Template = testData.FakePodTemplateSpec(t)
	j.ManagedFields = []metav1.ManagedFieldsEntry{testData.FakeManagedFields(t)}
	j.Spec.Template.ManagedFields = []metav1.ManagedFieldsEntry{testData.FakeManagedFields(t)}
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
