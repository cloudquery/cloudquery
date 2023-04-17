package elasticbeanstalk

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	elasticbeanstalkTypes "github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v2/faker"
	"github.com/golang/mock/gomock"
)

func buildElasticbeanstalkApplications(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockElasticbeanstalkClient(ctrl)

	la := elasticbeanstalkTypes.ApplicationDescription{}
	err := faker.FakeObject(&la)
	if err != nil {
		t.Fatal(err)
	}

	tag := elasticbeanstalkTypes.Tag{}
	err = faker.FakeObject(&tag)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeApplications(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&elasticbeanstalk.DescribeApplicationsOutput{
			Applications: []elasticbeanstalkTypes.ApplicationDescription{la},
		}, nil)

	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&elasticbeanstalk.ListTagsForResourceOutput{
			ResourceTags: []elasticbeanstalkTypes.Tag{tag},
			ResourceArn:  la.ApplicationArn,
		}, nil)

	return client.Services{
		Elasticbeanstalk: m,
	}
}

func TestElasticbeanstalkApplications(t *testing.T) {
	client.AwsMockTestHelper(t, Applications(), buildElasticbeanstalkApplications, client.TestOptions{})
}
