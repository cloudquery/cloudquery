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

func buildSageMakerNotebookInstances(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockSagemakerClient(ctrl)

	summ := types.NotebookInstanceSummary{}
	if err := faker.FakeObject(&summ); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListNotebookInstances(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&sagemaker.ListNotebookInstancesOutput{NotebookInstances: []types.NotebookInstanceSummary{summ}},
		nil,
	)

	note := sagemaker.DescribeNotebookInstanceOutput{}
	if err := faker.FakeObject(&note); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeNotebookInstance(gomock.Any(), gomock.Any(), gomock.Any()).Return(
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

func TestSageMakerNotebookInstances(t *testing.T) {
	client.AwsMockTestHelper(t, NotebookInstances(), buildSageMakerNotebookInstances, client.TestOptions{})
}
