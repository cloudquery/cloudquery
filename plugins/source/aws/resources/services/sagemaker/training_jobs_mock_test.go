package sagemaker

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	types "github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildSageMakerTrainingJobs(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockSagemakerClient(ctrl)

	summ := types.TrainingJobSummary{}
	if err := faker.FakeObject(&summ); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListTrainingJobs(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&sagemaker.ListTrainingJobsOutput{TrainingJobSummaries: []types.TrainingJobSummary{summ}},
		nil,
	)

	note := sagemaker.DescribeTrainingJobOutput{}
	if err := faker.FakeObject(&note); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeTrainingJob(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&note,
		nil,
	)

	var tagsOut sagemaker.ListTagsOutput
	if err := faker.FakeObject(&tagsOut); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTags(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&tagsOut, nil,
	)

	return client.Services{
		Sagemaker: m,
	}
}

func TestSageMakerTrainingJobs(t *testing.T) {
	client.AwsMockTestHelper(t, TrainingJobs(), buildSageMakerTrainingJobs, client.TestOptions{})
}
