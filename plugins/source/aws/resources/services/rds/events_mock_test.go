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

func buildRDSEvents(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockRdsClient(ctrl)
	var events []types.Event
	require.NoError(t, faker.FakeObject(&events))

	mock.EXPECT().DescribeEvents(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&rds.DescribeEventsOutput{Events: events},
		nil,
	)
	return client.Services{Rds: mock}
}

func TestRDSEvents(t *testing.T) {
	client.AwsMockTestHelper(t, Events(), buildRDSEvents, client.TestOptions{})
}
