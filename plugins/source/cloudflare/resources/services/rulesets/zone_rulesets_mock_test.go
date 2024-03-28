package rulesets

import (
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildZoneRulesets(t *testing.T, ctrl *gomock.Controller) client.Clients {
	mock := mocks.NewMockApi(ctrl)

	p := cloudflare.PolishOff
	s := cloudflare.SecurityLevelEssentiallyOff
	ssl := cloudflare.SSLFull
	var ruleset cloudflare.Ruleset
	require.NoError(t, faker.FakeObject(&ruleset))
	ruleset.Rules[0].ActionParameters.Polish = &p
	ruleset.Rules[0].ActionParameters.SecurityLevel = &s
	ruleset.Rules[0].ActionParameters.SSL = &ssl
	mock.EXPECT().ListRulesets(
		gomock.Any(),
		cloudflare.ZoneIdentifier(client.TestZoneID),
		gomock.Any(),
	).Return(
		[]cloudflare.Ruleset{ruleset},
		nil,
	)

	return client.Clients{
		client.TestAccountID: mock,
	}
}

func TestZoneRuleSets(t *testing.T) {
	client.MockTestHelper(t, ZoneRulesets(), buildZoneRulesets)
}
