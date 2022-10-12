package cloudhsmv2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildHSMClusters(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockCloudHSMV2Client(ctrl)

	var clusters []types.Cluster
	if err := faker.FakeData(&clusters); err != nil {
		t.Fatal(err)
	}

	mock.EXPECT().DescribeClusters(
		gomock.Any(),
		&cloudhsmv2.DescribeClustersInput{},
		gomock.Any(),
	).Return(
		&cloudhsmv2.DescribeClustersOutput{Clusters: clusters},
		nil,
	)

	return client.Services{CloudHSMV2: mock}
}

func TestClusters(t *testing.T) {
	client.AwsMockTestHelper(t, Clusters(), buildHSMClusters, client.TestOptions{})
}
