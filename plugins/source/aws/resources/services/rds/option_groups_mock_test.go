package rds

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildOptionGroups(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockRdsClient(ctrl)
	var s types.OptionGroup
	require.NoError(t, faker.FakeObject(&s))

	mock.EXPECT().DescribeOptionGroups(gomock.Any(), &rds.DescribeOptionGroupsInput{}, gomock.Any()).Return(
		&rds.DescribeOptionGroupsOutput{OptionGroupsList: []types.OptionGroup{s}},
		nil,
	)
	return client.Services{Rds: mock}
}

func TestRDSOptionGroups(t *testing.T) {
	client.AwsMockTestHelper(t, OptionGroups(), buildOptionGroups, client.TestOptions{})
}
