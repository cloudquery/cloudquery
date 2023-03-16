package rds

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildRdsClusterBacktracks(t *testing.T, mockRds *mocks.MockRdsClient) {
	var d types.DBClusterBacktrack
	err := faker.FakeObject(&d)
	if err != nil {
		t.Fatal(err)
	}

	mockRds.EXPECT().DescribeDBClusterBacktracks(gomock.Any(), gomock.Any(), gomock.Any()).Return(&rds.DescribeDBClusterBacktracksOutput{
		DBClusterBacktracks: []types.DBClusterBacktrack{d},
	}, nil)
}
