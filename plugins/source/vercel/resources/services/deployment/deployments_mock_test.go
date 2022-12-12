package deployment

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/vercel/client"
	"github.com/cloudquery/cloudquery/plugins/source/vercel/client/mocks"
	"github.com/cloudquery/cloudquery/plugins/source/vercel/internal/vercel"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildDeployments(t *testing.T, ctrl *gomock.Controller) client.VercelServices {
	mock := mocks.NewMockVercelServices(ctrl)

	var d vercel.Deployment
	if err := faker.FakeObject(&d); err != nil {
		t.Fatal(err)
	}

	mock.EXPECT().ListDeployments(gomock.Any(), gomock.Any()).Return([]vercel.Deployment{d}, &vercel.Paginator{}, nil)

	var c vercel.DeploymentCheck
	if err := faker.FakeObject(&c); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListDeploymentChecks(gomock.Any(), d.UID, gomock.Any()).Return([]vercel.DeploymentCheck{c}, &vercel.Paginator{}, nil)

	return mock
}

func TestDeployments(t *testing.T) {
	client.MockTestHelper(t, Deployments(), buildDeployments)
}
