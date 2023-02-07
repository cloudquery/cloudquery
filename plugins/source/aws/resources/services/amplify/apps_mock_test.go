package amplify

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/amplify"
	"github.com/aws/aws-sdk-go-v2/service/amplify/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildApps(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockAmplifyClient(ctrl)

	var app types.App
	if err := faker.FakeObject(&app); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListApps(gomock.Any(), gomock.Any()).Return(
		&amplify.ListAppsOutput{
			Apps: []types.App{app},
		},
		nil,
	)

	return client.Services{Amplify: m}
}

func TestApps(t *testing.T) {
	client.AwsMockTestHelper(t, Apps(), buildApps, client.TestOptions{})
}
