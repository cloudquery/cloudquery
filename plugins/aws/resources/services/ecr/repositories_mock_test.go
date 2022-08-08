package ecr

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildEcrRepositoriesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEcrClient(ctrl)
	l := types.Repository{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}
	i := types.ImageDetail{}
	err = faker.FakeData(&i)
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
	err = faker.FakeData(&tagResponse)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tagResponse, nil)

	return client.Services{
		ECR: m,
	}
}

func TestEcrRepositories(t *testing.T) {
	client.AwsMockTestHelper(t, Repositories(), buildEcrRepositoriesMock, client.TestOptions{})
}
