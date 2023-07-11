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

func buildSageMakerModels(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockSagemakerClient(ctrl)

	summ := types.ModelSummary{}
	require.NoError(t, faker.FakeObject(&summ))

	m.EXPECT().ListModels(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&sagemaker.ListModelsOutput{Models: []types.ModelSummary{summ}},
		nil,
	)

	model := sagemaker.DescribeModelOutput{}
	require.NoError(t, faker.FakeObject(&model))

	m.EXPECT().DescribeModel(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&model,
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

func TestSageMakerModels(t *testing.T) {
	client.AwsMockTestHelper(t, Models(), buildSageMakerModels, client.TestOptions{})
}
