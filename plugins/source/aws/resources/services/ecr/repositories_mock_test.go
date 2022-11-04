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

	iF := types.ImageScanFindings{}
	err = faker.FakeObject(&iF)
	if err != nil {
		t.Fatal(err)
	}

	iS := types.ImageScanStatus{}
	err = faker.FakeObject(&iS)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeImageScanFindings(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ecr.DescribeImageScanFindingsOutput{
			ImageScanFindings: &iF,
			ImageScanStatus:   &iS,
		}, nil)

	return client.Services{
		Ecr: m,
	}
}

func TestEcrRepositories(t *testing.T) {
	client.AwsMockTestHelper(t, Repositories(), buildEcrRepositoriesMock, client.TestOptions{})
}
