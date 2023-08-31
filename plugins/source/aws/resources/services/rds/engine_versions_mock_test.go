package rds

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildEngineVersionsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRdsClient(ctrl)
	services := client.Services{
		Rds: m,
	}
	var ev rds.DescribeDBEngineVersionsOutput
	require.NoError(t, faker.FakeObject(&ev))

	ev.Marker = nil
	var aurora types.DBEngineVersion
	require.NoError(t, faker.FakeObject(&aurora))

	aurora.DBParameterGroupFamily = aws.String("aurora-mysql5.7")
	ev.DBEngineVersions = append(ev.DBEngineVersions, aurora)
	m.EXPECT().DescribeDBEngineVersions(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ev,
		nil,
	)

	var parameters rds.DescribeEngineDefaultClusterParametersOutput
	require.NoError(t, faker.FakeObject(&parameters))

	m.EXPECT().DescribeEngineDefaultClusterParameters(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&parameters,
		nil,
	)

	return services
}

func TestEngineVersions(t *testing.T) {
	client.AwsMockTestHelper(t, EngineVersions(), buildEngineVersionsMock, client.TestOptions{})
}
