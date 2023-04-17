package glue

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v2/faker"
	"github.com/golang/mock/gomock"
)

func buildClassifiers(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockGlueClient(ctrl)

	var c glue.GetClassifiersOutput
	if err := faker.FakeObject(&c); err != nil {
		t.Fatal(err)
	}
	c.NextToken = nil
	m.EXPECT().GetClassifiers(gomock.Any(), gomock.Any()).Return(&c, nil)

	return client.Services{
		Glue: m,
	}
}

func TestClassifiers(t *testing.T) {
	client.AwsMockTestHelper(t, Classifiers(), buildClassifiers, client.TestOptions{})
}
