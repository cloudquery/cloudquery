package config

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildConfigRules(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockConfigserviceClient(ctrl)
	l := types.ConfigRule{}
	require.NoError(t, faker.FakeObject(&l))

	sl := types.ComplianceByConfigRule{}
	require.NoError(t, faker.FakeObject(&sl))

	m.EXPECT().DescribeConfigRules(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&configservice.DescribeConfigRulesOutput{
			ConfigRules: []types.ConfigRule{l},
		}, nil)
	m.EXPECT().DescribeComplianceByConfigRule(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&configservice.DescribeComplianceByConfigRuleOutput{
			ComplianceByConfigRules: []types.ComplianceByConfigRule{sl},
		}, nil)
	buildRemediationConfigurations(t, m)
	buildComplianceDetails(t, m)
	return client.Services{
		Configservice: m,
	}
}

func TestConfigRules(t *testing.T) {
	client.AwsMockTestHelper(t, ConfigRules(), buildConfigRules, client.TestOptions{})
}
