package kafka

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/kafka"
	"github.com/aws/aws-sdk-go-v2/service/kafka/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildKafkaConfigurationsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockKafkaClient(ctrl)
	object := types.Configuration{}
	require.NoError(t, faker.FakeObject(&object))

	m.EXPECT().ListConfigurations(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&kafka.ListConfigurationsOutput{
			Configurations: []types.Configuration{object},
		}, nil)

	tagsOutput := kafka.ListTagsForResourceOutput{}
	require.NoError(t, faker.FakeObject(&tagsOutput))
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tagsOutput, nil).AnyTimes()
	return client.Services{
		Kafka: m,
	}
}
func TestKafkaConfigurations(t *testing.T) {
	client.AwsMockTestHelper(t, Configurations(), buildKafkaConfigurationsMock, client.TestOptions{})
}
