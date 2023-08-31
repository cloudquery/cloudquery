package glue

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildRegistriesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockGlueClient(ctrl)

	var r types.RegistryListItem
	require.NoError(t, faker.FakeObject(&r))
	m.EXPECT().ListRegistries(
		gomock.Any(),
		&glue.ListRegistriesInput{MaxResults: aws.Int32(100)},
		gomock.Any(),
	).Return(
		&glue.ListRegistriesOutput{Registries: []types.RegistryListItem{r}},
		nil,
	)

	m.EXPECT().GetTags(
		gomock.Any(),
		&glue.GetTagsInput{ResourceArn: r.RegistryArn},
		gomock.Any(),
	).Return(
		&glue.GetTagsOutput{Tags: map[string]string{"tag": "value"}},
		nil,
	)

	var s glue.GetSchemaOutput
	require.NoError(t, faker.FakeObject(&s))
	m.EXPECT().ListSchemas(
		gomock.Any(),
		&glue.ListSchemasInput{
			RegistryId: &types.RegistryId{RegistryArn: r.RegistryArn},
			MaxResults: aws.Int32(100),
		},
		gomock.Any(),
	).Return(
		&glue.ListSchemasOutput{Schemas: []types.SchemaListItem{{SchemaArn: s.SchemaArn}}},
		nil,
	)

	m.EXPECT().GetSchema(
		gomock.Any(),
		&glue.GetSchemaInput{SchemaId: &types.SchemaId{SchemaArn: s.SchemaArn}},
		gomock.Any(),
	).Return(&s, nil)

	m.EXPECT().GetTags(
		gomock.Any(),
		&glue.GetTagsInput{ResourceArn: s.SchemaArn},
		gomock.Any(),
	).Return(
		&glue.GetTagsOutput{Tags: map[string]string{"tag": "value"}},
		nil,
	)

	var lsv glue.ListSchemaVersionsOutput
	require.NoError(t, faker.FakeObject(&lsv))
	lsv.NextToken = nil
	m.EXPECT().ListSchemaVersions(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(&lsv, nil)

	var sv glue.GetSchemaVersionOutput
	require.NoError(t, faker.FakeObject(&sv))
	m.EXPECT().GetSchemaVersion(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(&sv, nil)

	var sm glue.QuerySchemaVersionMetadataOutput
	require.NoError(t, faker.FakeObject(&sm))
	sm.NextToken = nil
	m.EXPECT().QuerySchemaVersionMetadata(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(&sm, nil)

	return client.Services{
		Glue: m,
	}
}

func TestRegistries(t *testing.T) {
	client.AwsMockTestHelper(t, Registries(), buildRegistriesMock, client.TestOptions{})
}
