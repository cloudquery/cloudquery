package kafka

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/kafka"
	"github.com/aws/aws-sdk-go-v2/service/kafka/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildKafkaNodesMock(t *testing.T, m *mocks.MockKafkaClient) {
	object := types.NodeInfo{}
	err := faker.FakeObject(&object)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListNodes(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&kafka.ListNodesOutput{
			NodeInfoList: []types.NodeInfo{object},
		}, nil)
}
