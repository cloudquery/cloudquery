package repositories

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/cloudquery/plugins/source/github/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v49/github"
)

func buildDependabot(t *testing.T, ctrl *gomock.Controller) client.DependabotService {
	mock := mocks.NewMockDependabotService(ctrl)

	var alert github.DependabotAlert
	if err := faker.FakeObject(&alert); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListRepoAlerts(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(
		[]*github.DependabotAlert{&alert}, &github.Response{}, nil)

	var secret github.Secret
	if err := faker.FakeObject(&secret); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListRepoSecrets(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&github.Secrets{TotalCount: 1, Secrets: []*github.Secret{&secret}}, &github.Response{}, nil)

	return mock
}
