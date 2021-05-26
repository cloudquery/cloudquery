package resources

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	iamTypes "github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildIamServerCerts(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIamClient(ctrl)
	u := iamTypes.ServerCertificateMetadata{}
	err := faker.FakeData(&u)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListServerCertificates(gomock.Any(), gomock.Any()).Return(
		&iam.ListServerCertificatesOutput{
			ServerCertificateMetadataList: []iamTypes.ServerCertificateMetadata{u},
		}, nil)

	return client.Services{
		IAM: m,
	}
}

func TestIamServerCertificates(t *testing.T) {
	awsTestHelper(t, IamServerCertificates(), buildIamServerCerts)
}
