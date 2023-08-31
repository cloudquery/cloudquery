package acmpca

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/acmpca"
	"github.com/aws/aws-sdk-go-v2/service/acmpca/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildACMPCACertificateAuthorities(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockAcmpcaClient(ctrl)

	var ca types.CertificateAuthority
	require.NoError(t, faker.FakeObject(&ca))

	mock.EXPECT().ListCertificateAuthorities(
		gomock.Any(),
		&acmpca.ListCertificateAuthoritiesInput{},
		gomock.Any(),
	).Return(
		&acmpca.ListCertificateAuthoritiesOutput{CertificateAuthorities: []types.CertificateAuthority{ca}},
		nil,
	)

	mock.EXPECT().ListTags(
		gomock.Any(),
		&acmpca.ListTagsInput{CertificateAuthorityArn: ca.Arn},
		gomock.Any(),
	).Return(
		&acmpca.ListTagsOutput{
			Tags: []types.Tag{
				{Key: aws.String("key"), Value: aws.String("value")},
			},
		},
		nil,
	)
	return client.Services{Acmpca: mock}
}

func TestACMPCACertificateAuthorities(t *testing.T) {
	client.AwsMockTestHelper(t, CertificateAuthorities(), buildACMPCACertificateAuthorities, client.TestOptions{})
}
