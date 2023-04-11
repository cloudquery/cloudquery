package sagemaker

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	types "github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v2/faker"
	"github.com/golang/mock/gomock"
)

func buildSageMakerApps(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockSagemakerClient(ctrl)

	appDets := types.AppDetails{}
	if err := faker.FakeObject(&appDets); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListApps(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&sagemaker.ListAppsOutput{Apps: []types.AppDetails{appDets}},
		nil,
	)

	app := sagemaker.DescribeAppOutput{}
	if err := faker.FakeObject(&app); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeApp(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&app,
		nil,
	)

	var tagsOut sagemaker.ListTagsOutput
	if err := faker.FakeObject(&tagsOut); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTags(gomock.Any(), &sagemaker.ListTagsInput{ResourceArn: app.AppArn}, gomock.Any()).Return(
		&tagsOut, nil,
	)

	return client.Services{
		Sagemaker: m,
	}
}

func TestSageMakerApps(t *testing.T) {
	client.AwsMockTestHelper(t, Apps(), buildSageMakerApps, client.TestOptions{})
}
