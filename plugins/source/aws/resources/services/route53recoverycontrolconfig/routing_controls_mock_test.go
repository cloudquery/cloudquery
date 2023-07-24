package route53recoverycontrolconfig

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/route53recoverycontrolconfig"
	"github.com/aws/aws-sdk-go-v2/service/route53recoverycontrolconfig/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildRoutingControls(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRoute53recoverycontrolconfigClient(ctrl)

	var rc types.RoutingControl
	require.NoError(t, faker.FakeObject(&rc))

	m.EXPECT().ListRoutingControls(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(
		&route53recoverycontrolconfig.ListRoutingControlsOutput{
			RoutingControls: []types.RoutingControl{rc},
		},
		nil,
	)

	return client.Services{
		Route53recoverycontrolconfig: m,
	}
}

func TestRoutingControls(t *testing.T) {
	client.AwsMockTestHelper(t, RoutingControls(), buildRoutingControls, client.TestOptions{Region: "us-west-2"})
}
