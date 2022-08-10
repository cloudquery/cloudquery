package iot

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildIotCaCertificatesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIOTClient(ctrl)

	ca := iot.ListCACertificatesOutput{}
	err := faker.FakeData(&ca)
	if err != nil {
		t.Fatal(err)
	}
	ca.NextMarker = nil
	m.EXPECT().ListCACertificates(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ca, nil)

	cd := iot.DescribeCACertificateOutput{}
	err = faker.FakeData(&cd)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeCACertificate(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&cd, nil)

	ct := iot.ListCertificatesByCAOutput{}
	err = faker.FakeData(&ct)
	if err != nil {
		t.Fatal(err)
	}
	ct.NextMarker = nil
	m.EXPECT().ListCertificatesByCA(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ct, nil)

	return client.Services{
		IOT: m,
	}
}

func TestIotCaCertificates(t *testing.T) {
	client.AwsMockTestHelper(t, IotCaCertificates(), buildIotCaCertificatesMock, client.TestOptions{})
}
