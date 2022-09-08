package services

import (
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildWorkersRoutes(t *testing.T, ctrl *gomock.Controller) client.Clients {
	mock := mocks.NewMockApi(ctrl)

	var workerRoute cloudflare.WorkerRoute
	if err := faker.FakeData(&workerRoute); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListWorkerRoutes(
		gomock.Any(),
		client.TestZoneID,
	).Return(
		cloudflare.WorkerRoutesResponse{
			Routes: []cloudflare.WorkerRoute{workerRoute},
		},
		nil,
	)

	return client.Clients{
		client.TestAccountID: mock,
	}
}

func TestWorkersRoutes(t *testing.T) {
	client.MockTestHelper(t, WorkersRoutes(), buildWorkersRoutes)
}
