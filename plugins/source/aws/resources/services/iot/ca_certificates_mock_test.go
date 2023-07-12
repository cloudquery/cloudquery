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

func buildIotCaCertificatesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIotClient(ctrl)

	ca := iot.ListCACertificatesOutput{}
	require.NoError(t, faker.FakeObject(&ca))
	ca.NextMarker = nil
	m.EXPECT().ListCACertificates(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ca, nil)

	cd := iot.DescribeCACertificateOutput{}
	require.NoError(t, faker.FakeObject(&cd))
	m.EXPECT().DescribeCACertificate(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&cd, nil)

	ct := iot.ListCertificatesByCAOutput{}
	require.NoError(t, faker.FakeObject(&ct))
	ct.NextMarker = nil
	m.EXPECT().ListCertificatesByCA(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ct, nil)

	return client.Services{
		Iot: m,
	}
}

func TestIotCaCertificates(t *testing.T) {
	client.AwsMockTestHelper(t, CaCertificates(), buildIotCaCertificatesMock, client.TestOptions{})
}
