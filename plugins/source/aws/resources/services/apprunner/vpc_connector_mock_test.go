package apprunner

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/apprunner"
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildApprunnerVpcConnectorsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockApprunnerClient(ctrl)
	vc := types.VpcConnector{}
	err := faker.FakeObject(&vc)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListVpcConnectors(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apprunner.ListVpcConnectorsOutput{
			VpcConnectors: []types.VpcConnector{vc},
		}, nil)
	tags := types.Tag{}
	err = faker.FakeObject(&tags)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apprunner.ListTagsForResourceOutput{Tags: []types.Tag{tags}}, nil)

	return client.Services{
		Apprunner: m,
	}
}

func TestApprunnerVpcConnector(t *testing.T) {
	client.AwsMockTestHelper(t, VpcConnectors(), buildApprunnerVpcConnectorsMock, client.TestOptions{})
}
