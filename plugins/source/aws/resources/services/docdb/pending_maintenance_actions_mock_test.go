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

func buildPendingMaintenanceActionsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDocdbClient(ctrl)
	services := client.Services{
		Docdb: m,
	}
	var ev docdb.DescribePendingMaintenanceActionsOutput
	require.NoError(t, faker.FakeObject(&ev))

	ev.Marker = nil
	m.EXPECT().DescribePendingMaintenanceActions(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ev,
		nil,
	)
	return services
}

func TestPendingMaintenanceActions(t *testing.T) {
	client.AwsMockTestHelper(t, PendingMaintenanceActions(), buildPendingMaintenanceActionsMock, client.TestOptions{})
}
