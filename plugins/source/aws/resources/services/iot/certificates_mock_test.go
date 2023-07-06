package iot

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildIotCertificatesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIotClient(ctrl)

	certs := iot.ListCertificatesOutput{}
	require.NoError(t, faker.FakeObject(&certs))
	certs.NextMarker = nil
	m.EXPECT().ListCertificates(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&certs, nil)

	cd := iot.DescribeCertificateOutput{}
	require.NoError(t, faker.FakeObject(&cd))
	m.EXPECT().DescribeCertificate(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&cd, nil)

	p := iot.ListAttachedPoliciesOutput{}
	require.NoError(t, faker.FakeObject(&p))
	p.NextMarker = nil
	m.EXPECT().ListAttachedPolicies(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&p, nil)

	return client.Services{
		Iot: m,
	}
}

func TestIotCertificates(t *testing.T) {
	client.AwsMockTestHelper(t, Certificates(), buildIotCertificatesMock, client.TestOptions{})
}
