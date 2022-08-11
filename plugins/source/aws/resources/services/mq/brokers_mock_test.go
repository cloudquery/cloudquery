package mq

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mq"
	"github.com/aws/aws-sdk-go-v2/service/mq/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildMqBrokers(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockMQClient(ctrl)

	bs := types.BrokerSummary{}
	if err := faker.FakeData(&bs); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListBrokers(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&mq.ListBrokersOutput{
			BrokerSummaries: []types.BrokerSummary{bs},
		}, nil)

	bo := mq.DescribeBrokerOutput{}
	if err := faker.FakeData(&bo); err != nil {
		t.Fatal(err)
	}
	bo.BrokerId = bs.BrokerId
	username := "test_username"
	bo.Users = []types.UserSummary{{Username: &username}}
	var cfgID types.ConfigurationId
	if err := faker.FakeData(&cfgID); err != nil {
		t.Fatal(err)
	}
	bo.Configurations.Current = &cfgID
	bo.Configurations.History = []types.ConfigurationId{cfgID}
	m.EXPECT().DescribeBroker(gomock.Any(), &mq.DescribeBrokerInput{BrokerId: bs.BrokerId}, gomock.Any()).Return(&bo, nil)

	uo := mq.DescribeUserOutput{}
	if err := faker.FakeData(&uo); err != nil {
		t.Fatal(err)
	}
	uo.Username = &username
	uo.BrokerId = bo.BrokerId
	m.EXPECT().DescribeUser(gomock.Any(), &mq.DescribeUserInput{BrokerId: bo.BrokerId, Username: &username}, gomock.Any()).Return(&uo, nil)

	var co mq.DescribeConfigurationOutput
	if err := faker.FakeData(&co); err != nil {
		t.Fatal(err)
	}
	co.Id = cfgID.Id
	m.EXPECT().DescribeConfiguration(gomock.Any(), &mq.DescribeConfigurationInput{ConfigurationId: cfgID.Id}, gomock.Any()).Return(&co, nil)

	revisions := mq.ListConfigurationRevisionsOutput{}
	if err := faker.FakeData(&revisions); err != nil {
		t.Fatal(err)
	}
	revisions.NextToken = nil
	m.EXPECT().ListConfigurationRevisions(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&revisions, nil)

	revision := mq.DescribeConfigurationRevisionOutput{}
	if err := faker.FakeData(&revision); err != nil {
		t.Fatal(err)
	}
	revision.Data = aws.String("PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiPz48aGVsbG8+d29ybGQ8L2hlbGxvPg==")
	m.EXPECT().DescribeConfigurationRevision(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&revision, nil)

	return client.Services{MQ: m}
}

func TestMqBrokers(t *testing.T) {
	client.AwsMockTestHelper(t, Brokers(), buildMqBrokers, client.TestOptions{})
}
