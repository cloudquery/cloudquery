package appsync

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/appsync"
	"github.com/aws/aws-sdk-go-v2/service/appsync/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildAppsyncGraphqlApisMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockAppsyncClient(ctrl)
	l := types.GraphqlApi{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListGraphqlApis(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&appsync.ListGraphqlApisOutput{
			GraphqlApis: []types.GraphqlApi{l},
		}, nil)

	return client.Services{
		Appsync: m,
	}
}

func TestAppSyncGraphqlApis(t *testing.T) {
	client.AwsMockTestHelper(t, GraphqlApis(), buildAppsyncGraphqlApisMock, client.TestOptions{})
}
