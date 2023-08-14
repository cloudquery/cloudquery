package docdb

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildEventsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDocdbClient(ctrl)
	services := client.Services{
		Docdb: m,
	}
	var e docdb.DescribeEventsOutput
	require.NoError(t, faker.FakeObject(&e))

	e.Marker = nil
	m.EXPECT().DescribeEvents(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&e,
		nil,
	)

	return services
}

func TestEvents(t *testing.T) {
	client.AwsMockTestHelper(t, Events(), buildEventsMock, client.TestOptions{})
}
