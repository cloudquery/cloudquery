package wafv2

import (
	"math/rand"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	cftypes "github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildWAFV2WebACLMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockWafV2Client(ctrl)
	cfm := mocks.NewMockCloudfrontClient(ctrl)

	tempWebACLSum := types.WebACLSummary{}
	if err := faker.FakeData(&tempWebACLSum); err != nil {
		t.Fatal(err)
	}
	// Faker can't handle recursive nested structs so we have to build some
	// parts from scratch.
	defaultAction := types.DefaultAction{}
	if err := faker.FakeData(&defaultAction); err != nil {
		t.Fatal(err)
	}
	visibilityConfig := types.VisibilityConfig{}
	if err := faker.FakeData(&visibilityConfig); err != nil {
		t.Fatal(err)
	}
	customRespBody := map[string]types.CustomResponseBody{}
	if err := faker.FakeData(&customRespBody); err != nil {
		t.Fatal(err)
	}
	overrideAction := types.OverrideAction{}
	if err := faker.FakeData(&overrideAction); err != nil {
		t.Fatal(err)
	}
	processRuleGroups := types.FirewallManagerRuleGroup{
		FirewallManagerStatement: &types.FirewallManagerStatement{},
		Name:                     aws.String(faker.Word()),
		OverrideAction:           &overrideAction,
		Priority:                 rand.Int31(),
		VisibilityConfig:         &visibilityConfig,
	}
	action := types.RuleAction{}
	if err := faker.FakeData(&action); err != nil {
		t.Fatal(err)
	}
	labels := make([]types.Label, 0)
	if err := faker.FakeData(&labels); err != nil {
		t.Fatal(err)
	}
	var tempResourceArns []string
	if err := faker.FakeData(&tempResourceArns); err != nil {
		t.Fatal(err)
	}
	var tempTags []types.Tag
	if err := faker.FakeData(&tempTags); err != nil {
		t.Fatal(err)
	}
	var loggingConfiguration types.LoggingConfiguration
	if err := faker.FakeData(&loggingConfiguration); err != nil {
		t.Fatal(err)
	}
	rule := types.Rule{
		Name:             aws.String(faker.Name()),
		Priority:         rand.Int31(),
		Statement:        &types.Statement{AndStatement: &types.AndStatement{}},
		VisibilityConfig: &visibilityConfig,
		Action:           &action,
		OverrideAction:   &overrideAction,
		RuleLabels:       labels,
	}
	for _, scope := range []types.Scope{types.ScopeCloudfront, types.ScopeRegional} {
		tempWebACL := types.WebACL{
			ARN:                                  aws.String(faker.UUIDHyphenated()),
			DefaultAction:                        &defaultAction,
			Id:                                   aws.String(faker.UUIDDigit()),
			Name:                                 aws.String(faker.Word()),
			VisibilityConfig:                     &visibilityConfig,
			Capacity:                             rand.Int63(),
			CustomResponseBodies:                 customRespBody,
			Description:                          aws.String(faker.Word()),
			LabelNamespace:                       aws.String(faker.Word()),
			ManagedByFirewallManager:             true,
			PostProcessFirewallManagerRuleGroups: []types.FirewallManagerRuleGroup{processRuleGroups},
			PreProcessFirewallManagerRuleGroups:  []types.FirewallManagerRuleGroup{processRuleGroups},
			Rules:                                []types.Rule{rule},
		}
		m.EXPECT().ListWebACLs(gomock.Any(), &wafv2.ListWebACLsInput{
			Scope: scope,
			Limit: aws.Int32(100),
		}, gomock.Any()).Return(&wafv2.ListWebACLsOutput{
			WebACLs: []types.WebACLSummary{tempWebACLSum},
		}, nil)
		m.EXPECT().GetWebACL(gomock.Any(), &wafv2.GetWebACLInput{
			Id:    tempWebACLSum.Id,
			Name:  tempWebACLSum.Name,
			Scope: scope,
		}, gomock.Any()).Return(&wafv2.GetWebACLOutput{
			WebACL: &tempWebACL,
		}, nil)
		m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(&wafv2.ListTagsForResourceOutput{
			TagInfoForResource: &types.TagInfoForResource{TagList: tempTags},
		}, nil)
		m.EXPECT().GetLoggingConfiguration(gomock.Any(), gomock.Any(), gomock.Any()).Return(&wafv2.GetLoggingConfigurationOutput{
			LoggingConfiguration: &loggingConfiguration,
		}, nil)
	}
	m.EXPECT().ListResourcesForWebACL(gomock.Any(), gomock.Any(), gomock.Any()).Return(&wafv2.ListResourcesForWebACLOutput{
		ResourceArns: tempResourceArns,
	}, nil)

	distributionList := cftypes.DistributionList{}
	if err := faker.FakeData(&distributionList); err != nil {
		t.Fatal(err)
	}
	distributionList.NextMarker = nil
	cfm.EXPECT().ListDistributionsByWebACLId(gomock.Any(), gomock.Any(), gomock.Any()).Return(&cloudfront.ListDistributionsByWebACLIdOutput{
		DistributionList: &distributionList,
	}, nil)

	return client.Services{WafV2: m, Cloudfront: cfm}
}

func TestWafV2WebACL(t *testing.T) {
	client.AwsMockTestHelper(t, Wafv2WebAcls(), buildWAFV2WebACLMock, client.TestOptions{})
}
