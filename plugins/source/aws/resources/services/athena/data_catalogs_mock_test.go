package athena

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildDataCatalogs(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockAthenaClient(ctrl)

	catalogs := athena.ListDataCatalogsOutput{}
	require.NoError(t, faker.FakeObject(&catalogs))

	catalogs.NextToken = nil
	m.EXPECT().ListDataCatalogs(gomock.Any(), gomock.Any(), gomock.Any()).Return(&catalogs, nil)

	catalog := athena.GetDataCatalogOutput{}
	require.NoError(t, faker.FakeObject(&catalog))

	catalogs.NextToken = nil
	m.EXPECT().GetDataCatalog(gomock.Any(), gomock.Any(), gomock.Any()).Return(&catalog, nil)

	databases := athena.ListDatabasesOutput{}
	require.NoError(t, faker.FakeObject(&databases))

	databases.NextToken = nil
	m.EXPECT().ListDatabases(gomock.Any(), gomock.Any(), gomock.Any()).Return(&databases, nil)

	tags := athena.ListTagsForResourceOutput{}
	require.NoError(t, faker.FakeObject(&tags))

	tags.NextToken = nil
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tags, nil)

	tables := athena.ListTableMetadataOutput{}
	require.NoError(t, faker.FakeObject(&tables))

	tables.NextToken = nil
	m.EXPECT().ListTableMetadata(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tables, nil)

	return client.Services{
		Athena: m,
	}
}

func TestDataCatalogs(t *testing.T) {
	client.AwsMockTestHelper(t, DataCatalogs(), buildDataCatalogs, client.TestOptions{})
}
