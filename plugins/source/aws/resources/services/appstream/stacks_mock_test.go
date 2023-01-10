package appstream

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/appstream"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildAppstreamStacksMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockAppstreamClient(ctrl)
	stack := types.Stack{}
	err := faker.FakeObject(&stack)
	if err != nil {
		t.Fatal(err)
	}

	stackEntitlements := types.Entitlement{}
	if faker.FakeObject(&stackEntitlements) != nil {
		t.Fatal(err)
	}

	stackUserAssociations := types.UserStackAssociation{}
	if faker.FakeObject(&stackUserAssociations) != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeStacks(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&appstream.DescribeStacksOutput{
			Stacks: []types.Stack{stack},
		}, nil)

	m.EXPECT().DescribeEntitlements(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&appstream.DescribeEntitlementsOutput{
			Entitlements: []types.Entitlement{stackEntitlements},
		}, nil)

	m.EXPECT().DescribeUserStackAssociations(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&appstream.DescribeUserStackAssociationsOutput{
			UserStackAssociations: []types.UserStackAssociation{stackUserAssociations},
		}, nil)

	return client.Services{
		Appstream: m,
	}
}

func TestAppstreamStacks(t *testing.T) {
	client.AwsMockTestHelper(t, Stacks(), buildAppstreamStacksMock, client.TestOptions{})
}
