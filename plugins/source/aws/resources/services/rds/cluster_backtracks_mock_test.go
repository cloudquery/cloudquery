package rds

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildRdsClusterBacktracks(t *testing.T, mockRds *mocks.MockRdsClient) {
	var d types.DBClusterBacktrack
	require.NoError(t, faker.FakeObject(&d))

	mockRds.EXPECT().DescribeDBClusterBacktracks(gomock.Any(), gomock.Any(), gomock.Any()).Return(&rds.DescribeDBClusterBacktracksOutput{
		DBClusterBacktracks: []types.DBClusterBacktrack{d},
	}, nil)
}
