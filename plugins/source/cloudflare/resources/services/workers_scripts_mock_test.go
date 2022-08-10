package services

import (
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cq-provider-cloudflare/client"
	"github.com/cloudquery/cq-provider-cloudflare/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildWorkersScripts(t *testing.T, ctrl *gomock.Controller) client.Clients {
	mock := mocks.NewMockApi(ctrl)

	var workerScript cloudflare.WorkerMetaData
	if err := faker.FakeData(&workerScript); err != nil {
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
	if err := faker.FakeData(&workerCronTrigger); err != nil {
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
	if err := faker.FakeData(&workerSecret); err != nil {
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

func TestWorkersScripts(t *testing.T) {
	client.CFMockTestHelper(t, WorkersScripts(), buildWorkersScripts)
}
