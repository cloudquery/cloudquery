package lightsail

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildDatabasesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockLightsailClient(ctrl)

	b := lightsail.GetRelationalDatabasesOutput{}
	require.NoError(t, faker.FakeObject(&b))
	b.NextPageToken = nil
	m.EXPECT().GetRelationalDatabases(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&b, nil)

	ac := lightsail.GetRelationalDatabaseParametersOutput{}
	require.NoError(t, faker.FakeObject(&ac))
	ac.NextPageToken = nil

	m.EXPECT().GetRelationalDatabaseParameters(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ac, nil)

	e := lightsail.GetRelationalDatabaseEventsOutput{}
	require.NoError(t, faker.FakeObject(&e))
	e.NextPageToken = nil

	m.EXPECT().GetRelationalDatabaseEvents(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&e, nil)
	ls := lightsail.GetRelationalDatabaseLogStreamsOutput{}
	require.NoError(t, faker.FakeObject(&ls))
	m.EXPECT().GetRelationalDatabaseLogStreams(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ls, nil)

	le := lightsail.GetRelationalDatabaseLogEventsOutput{}
	require.NoError(t, faker.FakeObject(&le))
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
