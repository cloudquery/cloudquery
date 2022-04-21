package config

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildConfigConformancePack(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockConfigServiceClient(ctrl)

	var cpd types.ConformancePackDetail
	if err := faker.FakeData(&cpd); err != nil {
		t.Fatal(err)
	}
	var cprc types.ConformancePackRuleCompliance
	if err := faker.FakeData(&cprc); err != nil {
		t.Fatal(err)
	}
	var cpre types.ConformancePackEvaluationResult
	if err := faker.FakeData(&cpre); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeConformancePacks(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&configservice.DescribeConformancePacksOutput{
			ConformancePackDetails: []types.ConformancePackDetail{cpd},
		}, nil)
	m.EXPECT().DescribeConformancePackCompliance(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&configservice.DescribeConformancePackComplianceOutput{
			ConformancePackRuleComplianceList: []types.ConformancePackRuleCompliance{cprc},
		}, nil)
	m.EXPECT().GetConformancePackComplianceDetails(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&configservice.GetConformancePackComplianceDetailsOutput{
			ConformancePackRuleEvaluationResults: []types.ConformancePackEvaluationResult{cpre},
		}, nil)

	return client.Services{
		ConfigService: m,
	}
}

func TestConfigConformancePack(t *testing.T) {
	client.AwsMockTestHelper(t, ConfigConformancePack(), buildConfigConformancePack, client.TestOptions{})
}
