package spaces

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
)

func createSpaces(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockSpacesService(ctrl)

	var data *s3.ListBucketsOutput
	if err := faker.FakeObject(&data); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListBuckets(gomock.Any(), gomock.Any(), gomock.Any()).Return(data, nil)

	var cors *s3.GetBucketCorsOutput
	if err := faker.FakeObject(&cors); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetBucketCors(gomock.Any(), gomock.Any(), gomock.Any()).Return(cors, nil)

	var acl *s3.GetBucketAclOutput
	if err := faker.FakeObject(&acl); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetBucketAcl(gomock.Any(), gomock.Any(), gomock.Any()).Return(acl, nil)

	return client.Services{
		Spaces: m,
	}
}

func TestSpaces(t *testing.T) {
	client.MockTestHelper(t, Spaces(), createSpaces, client.TestOptions{})
}
