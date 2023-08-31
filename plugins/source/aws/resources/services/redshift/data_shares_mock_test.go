package redshift

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildDataSharesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRedshiftClient(ctrl)
	ds := types.DataShare{}
	require.NoError(t, faker.FakeObject(&ds))

	m.EXPECT().DescribeDataShares(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&redshift.DescribeDataSharesOutput{
			DataShares: []types.DataShare{ds},
		}, nil)

	return client.Services{
		Redshift: m,
	}
}

func TestRedshiftDataShares(t *testing.T) {
	client.AwsMockTestHelper(t, DataShares(), buildDataSharesMock, client.TestOptions{})
}
