package ses

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildActiveReceiptRuleSets(t *testing.T, ctrl *gomock.Controller) client.Services {
	sesClient := mocks.NewMockSesClient(ctrl)

	data := new(ses.DescribeActiveReceiptRuleSetOutput)
	if err := faker.FakeObject(data); err != nil {
		t.Fatal(err)
	}

	sesClient.EXPECT().DescribeActiveReceiptRuleSet(gomock.Any(), gomock.Any()).Return(data, nil)

	return client.Services{
		Ses: sesClient,
	}
}

func TestActiveReceiptRuleSets(t *testing.T) {
	client.AwsMockTestHelper(t, ActiveReceiptRuleSets(), buildActiveReceiptRuleSets, client.TestOptions{})
}
