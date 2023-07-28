package route53recoverycontrolconfig

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/route53recoverycontrolconfig"
	"github.com/aws/aws-sdk-go-v2/service/route53recoverycontrolconfig/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildClusters(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRoute53recoverycontrolconfigClient(ctrl)

	var c types.Cluster
	require.NoError(t, faker.FakeObject(&c))

	m.EXPECT().ListClusters(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(
		&route53recoverycontrolconfig.ListClustersOutput{
			Clusters: []types.Cluster{c},
		},
		nil,
	)

	return client.Services{
		Route53recoverycontrolconfig: m,
	}
}

func TestClusters(t *testing.T) {
	client.AwsMockTestHelper(t, Clusters(), buildClusters, client.TestOptions{Region: "us-west-2"})
}
