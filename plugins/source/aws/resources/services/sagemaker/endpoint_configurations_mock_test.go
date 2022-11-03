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

func buildSageMakerEndpointConfigs(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockSagemakerClient(ctrl)

	summ := types.EndpointConfigSummary{}
	if err := faker.FakeObject(&summ); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListEndpointConfigs(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&sagemaker.ListEndpointConfigsOutput{EndpointConfigs: []types.EndpointConfigSummary{summ}},
		nil,
	)

	endpointConfig := sagemaker.DescribeEndpointConfigOutput{}
	if err := faker.FakeObject(&endpointConfig); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeEndpointConfig(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&endpointConfig,
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

func TestSageMakerEndpointConfigurations(t *testing.T) {
	client.AwsMockTestHelper(t, EndpointConfigurations(), buildSageMakerEndpointConfigs, client.TestOptions{})
}
