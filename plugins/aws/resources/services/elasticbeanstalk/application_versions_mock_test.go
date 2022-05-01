package elasticbeanstalk

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	elasticbeanstalkTypes "github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildElasticbeanstalkApplicationVersions(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockElasticbeanstalkClient(ctrl)

	la := elasticbeanstalkTypes.ApplicationVersionDescription{}
	err := faker.FakeData(&la)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeApplicationVersions(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&elasticbeanstalk.DescribeApplicationVersionsOutput{
			ApplicationVersions: []elasticbeanstalkTypes.ApplicationVersionDescription{la},
		}, nil)

	return client.Services{
		ElasticBeanstalk: m,
	}
}

func TestElasticbeanstalkApplicationVersions(t *testing.T) {
	client.AwsMockTestHelper(t, ApplicationVersions(), buildElasticbeanstalkApplicationVersions, client.TestOptions{})
}
