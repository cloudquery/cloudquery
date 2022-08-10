package athena

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildDataCatalogs(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockAthenaClient(ctrl)

	catalogs := athena.ListDataCatalogsOutput{}
	err := faker.FakeData(&catalogs)
	if err != nil {
		t.Fatal(err)
	}
	catalogs.NextToken = nil
	m.EXPECT().ListDataCatalogs(gomock.Any(), gomock.Any(), gomock.Any()).Return(&catalogs, nil)

	catalog := athena.GetDataCatalogOutput{}
	err = faker.FakeData(&catalog)
	if err != nil {
		t.Fatal(err)
	}
	catalogs.NextToken = nil
	m.EXPECT().GetDataCatalog(gomock.Any(), gomock.Any(), gomock.Any()).Return(&catalog, nil)

	databases := athena.ListDatabasesOutput{}
	err = faker.FakeData(&databases)
	if err != nil {
		t.Fatal(err)
	}
	databases.NextToken = nil
	m.EXPECT().ListDatabases(gomock.Any(), gomock.Any(), gomock.Any()).Return(&databases, nil)

	tags := athena.ListTagsForResourceOutput{}
	err = faker.FakeData(&tags)
	if err != nil {
		t.Fatal(err)
	}
	tags.NextToken = nil
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tags, nil)

	tables := athena.ListTableMetadataOutput{}
	err = faker.FakeData(&tables)
	if err != nil {
		t.Fatal(err)
	}
	tables.NextToken = nil
	m.EXPECT().ListTableMetadata(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tables, nil)

	return client.Services{
		Athena: m,
	}
}

func TestDataCatalogs(t *testing.T) {
	client.AwsMockTestHelper(t, DataCatalogs(), buildDataCatalogs, client.TestOptions{})
}
