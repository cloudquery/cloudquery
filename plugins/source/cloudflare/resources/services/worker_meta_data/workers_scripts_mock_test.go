package worker_meta_data

import (
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildWorkerMetaData(t *testing.T, ctrl *gomock.Controller) client.Clients {
	mock := mocks.NewMockApi(ctrl)

	var workerScript cloudflare.WorkerMetaData
	if err := faker.FakeObject(&workerScript); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListWorkerScripts(
		gomock.Any(),
	).Return(
		cloudflare.WorkerListResponse{
			WorkerList: []cloudflare.WorkerMetaData{workerScript},
		},
		nil,
	)

	var workerCronTrigger cloudflare.WorkerCronTrigger
	if err := faker.FakeObject(&workerCronTrigger); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListWorkerCronTriggers(
		gomock.Any(),
		client.TestAccountID,
		workerScript.ID,
	).Return(
		[]cloudflare.WorkerCronTrigger{workerCronTrigger},
		nil,
	)

	var workerSecret cloudflare.WorkersSecret
	if err := faker.FakeObject(&workerSecret); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListWorkersSecrets(
		gomock.Any(),
		workerScript.ID,
	).Return(
		cloudflare.WorkersListSecretsResponse{
			Result: []cloudflare.WorkersSecret{workerSecret},
		},
		nil,
	)

	return client.Clients{
		client.TestAccountID: mock,
	}
}

func TestWorkerMetaData(t *testing.T) {
	client.MockTestHelper(t, WorkerMetaData(), buildWorkerMetaData)
}
