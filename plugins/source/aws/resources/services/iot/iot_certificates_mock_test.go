package iot

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildIotCertificatesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIotClient(ctrl)

	certs := iot.ListCertificatesOutput{}
	err := faker.FakeObject(&certs)
	if err != nil {
		t.Fatal(err)
	}
	certs.NextMarker = nil
	m.EXPECT().ListCertificates(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&certs, nil)

	cd := iot.DescribeCertificateOutput{}
	err = faker.FakeObject(&cd)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeCertificate(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&cd, nil)

	p := iot.ListAttachedPoliciesOutput{}
	err = faker.FakeObject(&p)
	if err != nil {
		t.Fatal(err)
	}
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
