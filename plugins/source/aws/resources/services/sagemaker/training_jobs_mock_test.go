package sagemaker

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	types "github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildSageMakerTrainingJobs(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockSagemakerClient(ctrl)

	summ := types.TrainingJobSummary{}
	require.NoError(t, faker.FakeObject(&summ))

	m.EXPECT().ListTrainingJobs(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&sagemaker.ListTrainingJobsOutput{TrainingJobSummaries: []types.TrainingJobSummary{summ}},
		nil,
	)

	note := sagemaker.DescribeTrainingJobOutput{}
	require.NoError(t, faker.FakeObject(&note))

	m.EXPECT().DescribeTrainingJob(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&note,
		nil,
	)

	var tagsOut sagemaker.ListTagsOutput
	require.NoError(t, faker.FakeObject(&tagsOut))

	tagsOut.NextToken = nil
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
