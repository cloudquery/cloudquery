package lambda

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildLambdaLayersMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockLambdaClient(ctrl)

	creationDate := "1994-11-05T08:15:30.000+0500"

	l := types.LayersListItem{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}
	l.LatestMatchingVersion.CreatedDate = &creationDate
	m.EXPECT().ListLayers(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&lambda.ListLayersOutput{
			Layers: []types.LayersListItem{l},
		}, nil)

	lv := types.LayerVersionsListItem{}
	err = faker.FakeObject(&lv)
	if err != nil {
		t.Fatal(err)
	}
	arn := "arn:aws:s3:::my_corporate_bucket/test:exampleobject.png:1"
	lv.LayerVersionArn = &arn
	lv.CreatedDate = &creationDate
	m.EXPECT().ListLayerVersions(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&lambda.ListLayerVersionsOutput{
			LayerVersions: []types.LayerVersionsListItem{lv},
		}, nil)

	lvp := lambda.GetLayerVersionPolicyOutput{}
	err = faker.FakeObject(&lvp)
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
	client.AwsMockTestHelper(t, Layers(), buildLambdaLayersMock, client.TestOptions{})
}
