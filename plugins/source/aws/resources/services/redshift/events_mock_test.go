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

func buildEventsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRedshiftClient(ctrl)
	ev := types.Event{}
	require.NoError(t, faker.FakeObject(&ev))

	m.EXPECT().DescribeEvents(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&redshift.DescribeEventsOutput{
			Events: []types.Event{ev},
		}, nil)

	return client.Services{
		Redshift: m,
	}
}

func TestRedshiftEvents(t *testing.T) {
	client.AwsMockTestHelper(t, Events(), buildEventsMock, client.TestOptions{})
}
