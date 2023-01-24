package iam

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	iamTypes "github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildIamInstanceProfiles(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIamClient(ctrl)
	p := iamTypes.InstanceProfile{}
	err := faker.FakeObject(&p)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListInstanceProfiles(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.ListInstanceProfilesOutput{
			InstanceProfiles: []iamTypes.InstanceProfile{p},
		}, nil)

	//get tags
	tag := iamTypes.Tag{}
	err = faker.FakeObject(&tag)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListInstanceProfileTags(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.ListInstanceProfileTagsOutput{
			Tags: []iamTypes.Tag{
				tag,
			},
		}, nil)

	return client.Services{
		Iam: m,
	}
}

func TestIamInstanceProfiles(t *testing.T) {
	client.AwsMockTestHelper(t, InstanceProfiles(), buildIamInstanceProfiles, client.TestOptions{})
}
