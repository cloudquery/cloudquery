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

func buildSageMakerNotebookInstances(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockSagemakerClient(ctrl)

	summ := types.NotebookInstanceSummary{}
	require.NoError(t, faker.FakeObject(&summ))

	m.EXPECT().ListNotebookInstances(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&sagemaker.ListNotebookInstancesOutput{NotebookInstances: []types.NotebookInstanceSummary{summ}},
		nil,
	)

	note := sagemaker.DescribeNotebookInstanceOutput{}
	require.NoError(t, faker.FakeObject(&note))

	m.EXPECT().DescribeNotebookInstance(gomock.Any(), gomock.Any(), gomock.Any()).Return(
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

func TestSageMakerNotebookInstances(t *testing.T) {
	client.AwsMockTestHelper(t, NotebookInstances(), buildSageMakerNotebookInstances, client.TestOptions{})
}
