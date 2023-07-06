package ssm

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildSSMDocuments(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockSsmClient(ctrl)

	docName := "testDocName"
	mock.EXPECT().ListDocuments(
		gomock.Any(),
		&ssm.ListDocumentsInput{Filters: []types.DocumentKeyValuesFilter{{Key: aws.String("Owner"), Values: []string{"Self"}}}},
		gomock.Any(),
	).Return(
		&ssm.ListDocumentsOutput{DocumentIdentifiers: []types.DocumentIdentifier{{Name: &docName}}},
		nil,
	)

	var d types.DocumentDescription
	require.NoError(t, faker.FakeObject(&d))

	d.Name = &docName
	mock.EXPECT().DescribeDocument(
		gomock.Any(),
		&ssm.DescribeDocumentInput{Name: &docName},
		gomock.Any(),
	).Return(
		&ssm.DescribeDocumentOutput{Document: &d},
		nil,
	)

	mock.EXPECT().DescribeDocumentPermission(
		gomock.Any(),
		&ssm.DescribeDocumentPermissionInput{
			Name:           &docName,
			PermissionType: types.DocumentPermissionTypeShare,
		},
		gomock.Any(),
	).Return(
		&ssm.DescribeDocumentPermissionOutput{
			AccountIds:             []string{"some"},
			AccountSharingInfoList: []types.AccountSharingInfo{{AccountId: aws.String("other"), SharedDocumentVersion: aws.String("version")}},
		},
		nil,
	)

	var v types.DocumentVersionInfo
	require.NoError(t, faker.FakeObject(&v))

	v.Name = &docName
	mock.EXPECT().ListDocumentVersions(
		gomock.Any(),
		&ssm.ListDocumentVersionsInput{Name: &docName},
		gomock.Any(),
	).Return(
		&ssm.ListDocumentVersionsOutput{
			DocumentVersions: []types.DocumentVersionInfo{v},
		},
		nil,
	)

	return client.Services{Ssm: mock}
}

func TestSSMDocuments(t *testing.T) {
	client.AwsMockTestHelper(t, Documents(), buildSSMDocuments, client.TestOptions{})
}
