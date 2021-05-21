package resources

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildLambdaLayersMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockLambdaClient(ctrl)

	l := types.LayersListItem{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListLayers(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&lambda.ListLayersOutput{
			Layers: []types.LayersListItem{l},
		}, nil)

	lv := types.LayerVersionsListItem{}
	err = faker.FakeData(&lv)
	if err != nil {
		t.Fatal(err)
	}
	arn := "arn:aws:s3:::my_corporate_bucket/test:exampleobject.png:1"
	lv.LayerVersionArn = &arn
	m.EXPECT().ListLayerVersions(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&lambda.ListLayerVersionsOutput{
			LayerVersions: []types.LayerVersionsListItem{lv},
		}, nil)

	lvp := lambda.GetLayerVersionPolicyOutput{}
	err = faker.FakeData(&lvp)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetLayerVersionPolicy(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&lvp, nil)

	return client.Services{
		Lambda: m,
	}
}

func TestLambdaLayers(t *testing.T) {
	awsTestHelper(t, LambdaLayers(), buildLambdaLayersMock)
}
