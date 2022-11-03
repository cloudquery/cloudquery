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

func buildSageMakerModels(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockSagemakerClient(ctrl)

	summ := types.ModelSummary{}
	if err := faker.FakeObject(&summ); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListModels(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&sagemaker.ListModelsOutput{Models: []types.ModelSummary{summ}},
		nil,
	)

	model := sagemaker.DescribeModelOutput{}
	if err := faker.FakeObject(&model); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeModel(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&model,
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

func TestSageMakerModels(t *testing.T) {
	client.AwsMockTestHelper(t, Models(), buildSageMakerModels, client.TestOptions{})
}
