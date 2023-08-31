package lambda

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildLambdaLayersMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockLambdaClient(ctrl)

	creationDate := "1994-11-05T08:15:30.000+0500"

	l := types.LayersListItem{}
	require.NoError(t, faker.FakeObject(&l))
	l.LatestMatchingVersion.CreatedDate = &creationDate
	m.EXPECT().ListLayers(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&lambda.ListLayersOutput{
			Layers: []types.LayersListItem{l},
		}, nil)

	lv := types.LayerVersionsListItem{}
	require.NoError(t, faker.FakeObject(&lv))
	arn := "arn:aws:s3:::my_corporate_bucket/test:exampleobject.png:1"
	lv.LayerVersionArn = &arn
	lv.CreatedDate = &creationDate
	m.EXPECT().ListLayerVersions(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&lambda.ListLayerVersionsOutput{
			LayerVersions: []types.LayerVersionsListItem{lv},
		}, nil)

	lvp := lambda.GetLayerVersionPolicyOutput{}
	require.NoError(t, faker.FakeObject(&lvp))
	m.EXPECT().GetLayerVersionPolicy(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&lvp, nil)

	return client.Services{
		Lambda: m,
	}
}

func TestLambdaLayers(t *testing.T) {
	client.AwsMockTestHelper(t, Layers(), buildLambdaLayersMock, client.TestOptions{})
}
