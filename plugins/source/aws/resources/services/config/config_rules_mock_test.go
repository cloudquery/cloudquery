package config

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildConfigRules(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockConfigserviceClient(ctrl)
	l := types.ConfigRule{}
	if err := faker.FakeObject(&l); err != nil {
		t.Fatal(err)
	}
	sl := types.ComplianceByConfigRule{}
	if err := faker.FakeObject(&sl); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeConfigRules(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&configservice.DescribeConfigRulesOutput{
			ConfigRules: []types.ConfigRule{l},
		}, nil)
	m.EXPECT().DescribeComplianceByConfigRule(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&configservice.DescribeComplianceByConfigRuleOutput{
			ComplianceByConfigRules: []types.ComplianceByConfigRule{sl},
		}, nil)
	return client.Services{
		Configservice: m,
	}
}

func TestConfigRules(t *testing.T) {
	client.AwsMockTestHelper(t, ConfigRules(), buildConfigRules, client.TestOptions{})
}
