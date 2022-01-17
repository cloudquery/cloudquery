// +build mock

package iot

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildIotTopicRules(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIOTClient(ctrl)

	faker.SetIgnoreInterface(true)
	lp := iot.ListTopicRulesOutput{}
	err := faker.FakeData(&lp)
	if err != nil {
		t.Fatal(err)
	}
	lp.NextToken = nil
	m.EXPECT().ListTopicRules(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&lp, nil)

	p, err := buildRule()
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetTopicRule(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		p, nil)

	tags := iot.ListTagsForResourceOutput{}
	err = faker.FakeData(&tags)
	if err != nil {
		t.Fatal(err)
	}
	tags.NextToken = nil
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&tags, nil)

	return client.Services{
		IOT: m,
	}
}

func buildRule() (*iot.GetTopicRuleOutput, error) {
	p := types.TopicRule{}
	err := faker.FakeDataSkipFields(&p, []string{"Actions", "ErrorAction", "noSmithyDocumentSerde"})
	if err != nil {
		return nil, err
	}
	a := types.Action{}
	err = faker.FakeDataSkipFields(&a, []string{"IotSiteWise", "noSmithyDocumentSerde"})
	if err != nil {
		return nil, err
	}
	a.IotSiteWise = &types.IotSiteWiseAction{
		RoleArn: aws.String(faker.Word()),
	}
	p.Actions = []types.Action{
		a,
	}
	p.ErrorAction = &a

	return &iot.GetTopicRuleOutput{
		Rule:    &p,
		RuleArn: aws.String(faker.Word()),
	}, nil
}

func TestIotTopicRules(t *testing.T) {
	client.AwsMockTestHelper(t, IotTopicRules(), buildIotTopicRules, client.TestOptions{})
}
