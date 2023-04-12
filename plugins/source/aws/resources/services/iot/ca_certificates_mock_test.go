package iot

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildIotCaCertificatesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIotClient(ctrl)

	ca := iot.ListCACertificatesOutput{}
	err := faker.FakeObject(&ca)
	if err != nil {
		t.Fatal(err)
	}
	ca.NextMarker = nil
	m.EXPECT().ListCACertificates(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ca, nil)

	cd := iot.DescribeCACertificateOutput{}
	err = faker.FakeObject(&cd)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeCACertificate(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&cd, nil)

	ct := iot.ListCertificatesByCAOutput{}
	err = faker.FakeObject(&ct)
	if err != nil {
		t.Fatal(err)
	}
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
