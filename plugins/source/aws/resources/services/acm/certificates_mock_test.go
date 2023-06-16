package acm

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/aws/aws-sdk-go-v2/service/acm/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildACMCertificates(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockAcmClient(ctrl)

	var cs types.CertificateSummary
	require.NoError(t, faker.FakeObject(&cs))

	mock.EXPECT().ListCertificates(
		gomock.Any(),
		&acm.ListCertificatesInput{
			CertificateStatuses: types.CertificateStatus("").Values(),
			Includes: &types.Filters{
				ExtendedKeyUsage: []types.ExtendedKeyUsageName{types.ExtendedKeyUsageNameAny},
				KeyTypes:         types.KeyAlgorithm("").Values(),
				KeyUsage:         []types.KeyUsageName{types.KeyUsageNameAny},
			},
		},
		gomock.Any(),
	).Return(
		&acm.ListCertificatesOutput{CertificateSummaryList: []types.CertificateSummary{cs}},
		nil,
	)

	var cert types.CertificateDetail
	require.NoError(t, faker.FakeObject(&cert))

	cert.CertificateArn = cs.CertificateArn
	mock.EXPECT().DescribeCertificate(
		gomock.Any(),
		&acm.DescribeCertificateInput{CertificateArn: cs.CertificateArn},
		gomock.Any(),
	).Return(
		&acm.DescribeCertificateOutput{Certificate: &cert},
		nil,
	)

	mock.EXPECT().ListTagsForCertificate(
		gomock.Any(),
		&acm.ListTagsForCertificateInput{CertificateArn: cert.CertificateArn},
		gomock.Any(),
	).Return(
		&acm.ListTagsForCertificateOutput{
			Tags: []types.Tag{
				{Key: aws.String("key"), Value: aws.String("value")},
			},
		},
		nil,
	)
	return client.Services{Acm: mock}
}

func TestACMCertificates(t *testing.T) {
	client.AwsMockTestHelper(t, Certificates(), buildACMCertificates, client.TestOptions{})
}
