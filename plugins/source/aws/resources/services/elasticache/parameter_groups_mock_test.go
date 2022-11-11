package elasticache

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildElasticacheParameterGroups(t *testing.T, ctrl *gomock.Controller) client.Services {
	mockElasticache := mocks.NewMockElasticacheClient(ctrl)
	parameterGroupsOutput := elasticache.DescribeCacheParameterGroupsOutput{}
	err := faker.FakeObject(&parameterGroupsOutput)
	parameterGroupsOutput.Marker = nil
	if err != nil {
		t.Fatal(err)
	}

	parametersOutput := elasticache.DescribeCacheParametersOutput{}
	err = faker.FakeObject(&parametersOutput)
	parametersOutput.Marker = nil
	if err != nil {
		t.Fatal(err)
	}

	//expectedInput := elasticache.DescribeCacheParametersInput{
	//	CacheParameterGroupName: parameterGroupsOutput.CacheParameterGroups[0].CacheParameterGroupName}

	mockElasticache.EXPECT().DescribeCacheParameterGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(&parameterGroupsOutput, nil)
	// mockElasticache.EXPECT().DescribeCacheParameters(gomock.Any(), &expectedInput, gomock.Any()).Return(&parametersOutput, nil)

	return client.Services{
		Elasticache: mockElasticache,
	}
}

func TestElasticacheParameterGroups(t *testing.T) {
	client.AwsMockTestHelper(t, ParameterGroups(), buildElasticacheParameterGroups, client.TestOptions{})
}
