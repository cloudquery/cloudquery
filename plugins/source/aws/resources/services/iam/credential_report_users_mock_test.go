package iam

import (
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/gocarina/gocsv"
	"github.com/golang/mock/gomock"
	"testing"
	"time"
)

func buildCredentialReportUsers(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIamClient(ctrl)

	ru := CredentialReportUser{}
	err := faker.FakeData(&ru)
	if err != nil {
		t.Fatal(err)
	}
	ru.Arn = "arn123"
	ru.PasswordStatus = "true"
	ru.PasswordNextRotation = time.Now().Format(time.RFC3339)
	ru.PasswordLastChanged = time.Now().Format(time.RFC3339)
	ru.AccessKey1LastRotated = time.Now().Format(time.RFC3339)
	ru.AccessKey2LastRotated = time.Now().Format(time.RFC3339)
	ru.Cert1LastRotated = time.Now().Format(time.RFC3339)
	ru.Cert2LastRotated = time.Now().Format(time.RFC3339)
	content, err := gocsv.MarshalBytes([]CredentialReportUser{ru})
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().GetCredentialReport(gomock.Any(), gomock.Any()).Return(
		&iam.GetCredentialReportOutput{
			Content: content,
		}, nil)

	return client.Services{
		IAM: m,
	}
}

func TestCredentialReportUsers(t *testing.T) {
	client.AwsMockTestHelper(t, CredentialReportUsers(), buildCredentialReportUsers, client.TestOptions{})
}
