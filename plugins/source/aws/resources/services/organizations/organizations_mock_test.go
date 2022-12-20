package organizations

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildOrganizations(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockOrganizationsClient(ctrl)

	o := organizations.DescribeOrganizationOutput{}
	err := faker.FakeObject(&o)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeOrganization(gomock.Any(), gomock.Any()).Return(&o, nil)

	return client.Services{
		Organizations: m,
	}
}

func TestOrganizations(t *testing.T) {
	client.AwsMockTestHelper(t, Organizations(), buildOrganizations, client.TestOptions{})
}
