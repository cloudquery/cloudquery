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

func buildSageMakerModels(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockSageMakerClient(ctrl)

	summ := types.ModelSummary{}
	if err := faker.FakeData(&summ); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListModels(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&sagemaker.ListModelsOutput{Models: []types.ModelSummary{summ}},
		nil,
	)

	model := sagemaker.DescribeModelOutput{}
	if err := faker.FakeData(&model); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeModel(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&model,
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

func TestSageMakerModels(t *testing.T) {
	client.AwsMockTestHelper(t, SagemakerModels(), buildSageMakerModels, client.TestOptions{})
}
