// +build mock

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

func buildElasticbeanstalkApplications(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockElasticbeanstalkClient(ctrl)

	la := elasticbeanstalkTypes.ApplicationDescription{}
	err := faker.FakeData(&la)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeApplications(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&elasticbeanstalk.DescribeApplicationsOutput{
			Applications: []elasticbeanstalkTypes.ApplicationDescription{la},
		}, nil)

	return client.Services{
		ElasticBeanstalk: m,
	}
}

func TestElasticbeanstalkApplications(t *testing.T) {
	client.AwsMockTestHelper(t, ElasticbeanstalkApplications(), buildElasticbeanstalkApplications, client.TestOptions{})
}
