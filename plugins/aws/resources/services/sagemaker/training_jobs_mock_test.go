package sagemaker

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	types "github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildSageMakerTrainingJobs(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockSageMakerClient(ctrl)

	summ := types.TrainingJobSummary{}
	if err := faker.FakeData(&summ); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListTrainingJobs(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&sagemaker.ListTrainingJobsOutput{TrainingJobSummaries: []types.TrainingJobSummary{summ}},
		nil,
	)

	note := sagemaker.DescribeTrainingJobOutput{}
	if err := faker.FakeData(&note); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeTrainingJob(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&note,
		nil,
	)

	var tagsOut sagemaker.ListTagsOutput
	if err := faker.FakeData(&tagsOut); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTags(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&tagsOut, nil,
	)

	return client.Services{
		SageMaker: m,
	}
}

func TestSageMakerTrainingJobs(t *testing.T) {
	client.AwsMockTestHelper(t, SagemakerTrainingJobs(), buildSageMakerTrainingJobs, client.TestOptions{})
}
