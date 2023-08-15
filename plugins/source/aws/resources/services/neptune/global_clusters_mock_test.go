package neptune

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/neptune"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildNeptuneGlobalClusters(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockNeptuneClient(ctrl)
	var gc types.GlobalCluster
	require.NoError(t, faker.FakeObject(&gc))

	mock.EXPECT().DescribeGlobalClusters(gomock.Any(), &neptune.DescribeGlobalClustersInput{}, gomock.Any()).Return(
		&neptune.DescribeGlobalClustersOutput{GlobalClusters: []types.GlobalCluster{gc}},
		nil,
	)
	return client.Services{Neptune: mock}
}

func TestNeptuneGlobalCluster(t *testing.T) {
	client.AwsMockTestHelper(t, GlobalClusters(), buildNeptuneGlobalClusters, client.TestOptions{})
}
