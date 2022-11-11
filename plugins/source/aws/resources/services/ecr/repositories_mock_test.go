package ecr

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildEcrRepositoriesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEcrClient(ctrl)
	l := types.Repository{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}
	i := types.ImageDetail{}
	err = faker.FakeObject(&i)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeRepositories(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ecr.DescribeRepositoriesOutput{
			Repositories: []types.Repository{l},
		}, nil)

	m.EXPECT().DescribeImages(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ecr.DescribeImagesOutput{
			ImageDetails: []types.ImageDetail{i},
		}, nil)

	tagResponse := ecr.ListTagsForResourceOutput{}
	err = faker.FakeObject(&tagResponse)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tagResponse, nil)

	iF := ecr.DescribeImageScanFindingsOutput{}
	err = faker.FakeObject(&iF)
	if err != nil {
		t.Fatal(err)
	}

	iF.NextToken = nil
	m.EXPECT().DescribeImageScanFindings(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iF, nil)

	repoResponse := ecr.GetRepositoryPolicyOutput{}
	err = faker.FakeObject(&repoResponse)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetRepositoryPolicy(gomock.Any(), gomock.Any(), gomock.Any()).Return(&repoResponse, nil)

	return client.Services{
		Ecr: m,
	}
}

func TestEcrRepositories(t *testing.T) {
	client.AwsMockTestHelper(t, Repositories(), buildEcrRepositoriesMock, client.TestOptions{})
}
