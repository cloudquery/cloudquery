// +build mock

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

func buildSageMakerNotebookInstances(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockSageMakerClient(ctrl)

	summ := types.NotebookInstanceSummary{}
	if err := faker.FakeData(&summ); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListNotebookInstances(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&sagemaker.ListNotebookInstancesOutput{NotebookInstances: []types.NotebookInstanceSummary{summ}},
		nil,
	)

	note := sagemaker.DescribeNotebookInstanceOutput{}
	if err := faker.FakeData(&note); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeNotebookInstance(gomock.Any(), gomock.Any(), gomock.Any()).Return(
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

func TestSageMakerNotebookInstances(t *testing.T) {
	client.AwsMockTestHelper(t, SagemakerNotebookInstances(), buildSageMakerNotebookInstances, client.TestOptions{})
}
