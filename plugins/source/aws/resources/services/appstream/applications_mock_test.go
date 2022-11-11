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

func buildAppstreamApplicationsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockAppstreamClient(ctrl)
	application := types.Application{}
	err := faker.FakeObject(&application)
	if err != nil {
		t.Fatal(err)
	}

	applicationFleetAssocition := types.ApplicationFleetAssociation{}
	if faker.FakeObject(&applicationFleetAssocition) != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeApplications(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&appstream.DescribeApplicationsOutput{
			Applications: []types.Application{application},
		}, nil)

	m.EXPECT().DescribeApplicationFleetAssociations(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&appstream.DescribeApplicationFleetAssociationsOutput{
			ApplicationFleetAssociations: []types.ApplicationFleetAssociation{applicationFleetAssocition},
		}, nil)

	return client.Services{
		Appstream: m,
	}
}

func TestAppstreamApplications(t *testing.T) {
	client.AwsMockTestHelper(t, Applications(), buildAppstreamApplicationsMock, client.TestOptions{})
}
