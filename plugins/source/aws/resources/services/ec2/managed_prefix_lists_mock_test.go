package ec2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildEc2ManagedPrefixList(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	l := types.ManagedPrefixList{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeManagedPrefixLists(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeManagedPrefixListsOutput{
			PrefixLists: []types.ManagedPrefixList{l},
		}, nil)
	return client.Services{
		Ec2: m,
	}
}

func TestEc2ManagedPrefixList(t *testing.T) {
	client.AwsMockTestHelper(t, ManagedPrefixLists(), buildEc2ManagedPrefixList, client.TestOptions{})
}
