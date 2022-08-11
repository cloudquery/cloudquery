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

func buildSageMakerEndpointConfigs(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockSageMakerClient(ctrl)

	summ := types.EndpointConfigSummary{}
	if err := faker.FakeData(&summ); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListEndpointConfigs(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&sagemaker.ListEndpointConfigsOutput{EndpointConfigs: []types.EndpointConfigSummary{summ}},
		nil,
	)

	endpointConfig := sagemaker.DescribeEndpointConfigOutput{}
	if err := faker.FakeData(&endpointConfig); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeEndpointConfig(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&endpointConfig,
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

func TestSageMakerEndpointConfigurations(t *testing.T) {
	client.AwsMockTestHelper(t, SagemakerEndpointConfigurations(), buildSageMakerEndpointConfigs, client.TestOptions{})
}
