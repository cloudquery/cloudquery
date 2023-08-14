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

func buildControlPanels(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRoute53recoverycontrolconfigClient(ctrl)

	var c types.ControlPanel
	require.NoError(t, faker.FakeObject(&c))

	m.EXPECT().ListControlPanels(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(
		&route53recoverycontrolconfig.ListControlPanelsOutput{
			ControlPanels: []types.ControlPanel{c},
		},
		nil,
	)

	var ruleAssert types.Rule
	require.NoError(t, faker.FakeObject(&ruleAssert))
	ruleAssert.GATING = nil
	var ruleGaiting types.Rule
	require.NoError(t, faker.FakeObject(&ruleGaiting))
	ruleGaiting.ASSERTION = nil

	m.EXPECT().ListSafetyRules(
		gomock.Any(),
		&route53recoverycontrolconfig.ListSafetyRulesInput{
			ControlPanelArn: c.ControlPanelArn,
		},
		gomock.Any(),
	).Return(
		&route53recoverycontrolconfig.ListSafetyRulesOutput{
			SafetyRules: []types.Rule{ruleAssert, ruleGaiting},
		},
		nil,
	)

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

func TestControlPanels(t *testing.T) {
	client.AwsMockTestHelper(t, ControlPanels(), buildControlPanels, client.TestOptions{Region: "us-west-2"})
}
