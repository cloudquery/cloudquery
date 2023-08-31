package apigateway

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildApigatewayRestApis(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockApigatewayClient(ctrl)

	r := types.RestApi{}
	require.NoError(t, faker.FakeObject(&r))

	m.EXPECT().GetRestApis(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigateway.GetRestApisOutput{
			Items: []types.RestApi{r},
		}, nil)

	a := types.Authorizer{}
	require.NoError(t, faker.FakeObject(&a))

	m.EXPECT().GetAuthorizers(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigateway.GetAuthorizersOutput{
			Items: []types.Authorizer{a},
		}, nil)

	d := types.Deployment{}
	require.NoError(t, faker.FakeObject(&d))

	m.EXPECT().GetDeployments(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigateway.GetDeploymentsOutput{
			Items: []types.Deployment{d},
		}, nil)

	dp := types.DocumentationPart{}
	require.NoError(t, faker.FakeObject(&dp))

	m.EXPECT().GetDocumentationParts(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigateway.GetDocumentationPartsOutput{
			Items: []types.DocumentationPart{dp},
		}, nil)

	dv := types.DocumentationVersion{}
	require.NoError(t, faker.FakeObject(&dv))

	m.EXPECT().GetDocumentationVersions(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigateway.GetDocumentationVersionsOutput{
			Items: []types.DocumentationVersion{dv},
		}, nil)

	gr := types.GatewayResponse{}
	require.NoError(t, faker.FakeObject(&gr))

	m.EXPECT().GetGatewayResponses(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigateway.GetGatewayResponsesOutput{
			Items: []types.GatewayResponse{gr},
		}, nil)

	am := types.Model{}
	require.NoError(t, faker.FakeObject(&am))

	m.EXPECT().GetModels(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigateway.GetModelsOutput{
			Items: []types.Model{am},
		}, nil)

	mt := apigateway.GetModelTemplateOutput{}
	require.NoError(t, faker.FakeObject(&mt))

	m.EXPECT().GetModelTemplate(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&mt, nil)

	rv := types.RequestValidator{}
	require.NoError(t, faker.FakeObject(&rv))

	m.EXPECT().GetRequestValidators(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigateway.GetRequestValidatorsOutput{
			Items: []types.RequestValidator{rv},
		}, nil)

	ar := types.Resource{}
	require.NoError(t, faker.FakeObject(&ar))

	m.EXPECT().GetResources(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigateway.GetResourcesOutput{
			Items: []types.Resource{ar},
		}, nil)

	method := apigateway.GetMethodOutput{}
	require.NoError(t, faker.FakeObject(&method))

	m.EXPECT().GetMethod(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&method, nil)

	integration := apigateway.GetIntegrationOutput{}
	require.NoError(t, faker.FakeObject(&integration))

	m.EXPECT().GetIntegration(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&integration, nil)

	s := types.Stage{}
	require.NoError(t, faker.FakeObject(&s))

	m.EXPECT().GetStages(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigateway.GetStagesOutput{
			Item: []types.Stage{s},
		}, nil)

	return client.Services{
		Apigateway: m,
	}
}

func TestApigatewayRestApis(t *testing.T) {
	client.AwsMockTestHelper(t, RestApis(), buildApigatewayRestApis, client.TestOptions{})
}
