package lightsail

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildDatabasesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockLightsailClient(ctrl)

	b := lightsail.GetRelationalDatabasesOutput{}
	err := faker.FakeObject(&b)
	if err != nil {
		t.Fatal(err)
	}
	b.NextPageToken = nil
	m.EXPECT().GetRelationalDatabases(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&b, nil)

	ac := lightsail.GetRelationalDatabaseParametersOutput{}
	err = faker.FakeObject(&ac)
	if err != nil {
		t.Fatal(err)
	}
	ac.NextPageToken = nil

	m.EXPECT().GetRelationalDatabaseParameters(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ac, nil)

	e := lightsail.GetRelationalDatabaseEventsOutput{}
	err = faker.FakeObject(&e)
	if err != nil {
		t.Fatal(err)
	}
	e.NextPageToken = nil

	m.EXPECT().GetRelationalDatabaseEvents(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&e, nil)
	ls := lightsail.GetRelationalDatabaseLogStreamsOutput{}
	err = faker.FakeObject(&ls)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetRelationalDatabaseLogStreams(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ls, nil)

	le := lightsail.GetRelationalDatabaseLogEventsOutput{}
	err = faker.FakeObject(&le)
	if err != nil {
		t.Fatal(err)
	}
	le.NextForwardToken = nil
	m.EXPECT().GetRelationalDatabaseLogEvents(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&le, nil)

	return client.Services{
		Lightsail: m,
	}
}

func TestDatabases(t *testing.T) {
	client.AwsMockTestHelper(t, Databases(), buildDatabasesMock, client.TestOptions{})
}
