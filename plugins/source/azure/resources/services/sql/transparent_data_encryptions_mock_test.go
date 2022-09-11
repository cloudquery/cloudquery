// Auto generated code - DO NOT EDIT.

package sql

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"
)

func createTransparentDataEncryptionsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSQLTransparentDataEncryptionsClient(ctrl)
	s := services.Services{
		SQL: services.SQLClient{
			TransparentDataEncryptions: mockClient,
		},
	}

	data := sql.TransparentDataEncryption{}
	require.Nil(t, faker.FakeObject(&data))

	result := sql.NewTransparentDataEncryptionListResultPage(sql.TransparentDataEncryptionListResult{Value: &[]sql.TransparentDataEncryption{data}}, func(ctx context.Context, result sql.TransparentDataEncryptionListResult) (sql.TransparentDataEncryptionListResult, error) {
		return sql.TransparentDataEncryptionListResult{}, nil
	})

	mockClient.EXPECT().Get(gomock.Any(), "test", "test").Return(result, nil)
	return s
}
