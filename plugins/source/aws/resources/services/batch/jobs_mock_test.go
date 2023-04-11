package batch

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/batch"
	"github.com/aws/aws-sdk-go-v2/service/batch/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v2/faker"
	"github.com/golang/mock/gomock"
)

func buildBatchJobsMock(t *testing.T, m *mocks.MockBatchClient) client.Services {
	services := client.Services{
		Batch: m,
	}
	a := types.JobSummary{}
	err := faker.FakeObject(&a)
	if err != nil {
		t.Fatal(err)
	}

	d := types.JobDetail{}
	err = faker.FakeObject(&d)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListJobs(gomock.Any(), gomock.Any()).Return(
		&batch.ListJobsOutput{
			JobSummaryList: []types.JobSummary{a},
		}, nil).Times(len(allJobStatuses))

	m.EXPECT().DescribeJobs(gomock.Any(), &batch.DescribeJobsInput{
		Jobs: []string{*a.JobId},
	}).Return(
		&batch.DescribeJobsOutput{
			Jobs: []types.JobDetail{d},
		}, nil).Times(len(allJobStatuses))

	tagResponse := batch.ListTagsForResourceOutput{}
	err = faker.FakeObject(&tagResponse)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tagResponse, nil).Times(len(allJobStatuses))

	return services
}
