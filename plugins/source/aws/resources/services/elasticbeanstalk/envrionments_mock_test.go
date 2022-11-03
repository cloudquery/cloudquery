package elasticbeanstalk

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	elasticbeanstalkTypes "github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildElasticbeanstalkEnvironments(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockElasticbeanstalkClient(ctrl)

	la := elasticbeanstalkTypes.ApplicationDescription{}
	err := faker.FakeObject(&la)
	if err != nil {
		t.Fatal(err)
	}

	l := elasticbeanstalkTypes.EnvironmentDescription{
		ApplicationName: la.ApplicationName,
	}
	err = faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeEnvironments(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&elasticbeanstalk.DescribeEnvironmentsOutput{
			Environments: []elasticbeanstalkTypes.EnvironmentDescription{l},
		}, nil)

	tags := elasticbeanstalk.ListTagsForResourceOutput{}
	err = faker.FakeObject(&tags)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&tags, nil)

	configSettingsOutput := elasticbeanstalk.DescribeConfigurationSettingsOutput{}
	err = faker.FakeObject(&configSettingsOutput)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeConfigurationSettings(gomock.Any(), gomock.Any(), gomock.Any()).Return(&configSettingsOutput, nil)

	configOptsOutput := elasticbeanstalk.DescribeConfigurationOptionsOutput{}
	err = faker.FakeObject(&configOptsOutput)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeConfigurationOptions(gomock.Any(), gomock.Any(), gomock.Any()).Return(&configOptsOutput, nil)

	return client.Services{
		Elasticbeanstalk: m,
	}
}

func TestElasticbeanstalkEnvironments(t *testing.T) {
	client.AwsMockTestHelper(t, Environments(), buildElasticbeanstalkEnvironments, client.TestOptions{})
}
