package apigatewayv2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildApigatewayv2Apis(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockApigatewayv2Client(ctrl)

	a := types.Api{}
	require.NoError(t, faker.FakeObject(&a))
	m.EXPECT().GetApis(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigatewayv2.GetApisOutput{
			Items: []types.Api{a},
		}, nil)

	aa := types.Authorizer{}
	require.NoError(t, faker.FakeObject(&aa))
	m.EXPECT().GetAuthorizers(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigatewayv2.GetAuthorizersOutput{
			Items: []types.Authorizer{aa},
		}, nil)

	d := types.Deployment{}
	require.NoError(t, faker.FakeObject(&d))
	m.EXPECT().GetDeployments(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigatewayv2.GetDeploymentsOutput{
			Items: []types.Deployment{d},
		}, nil)

	i := types.Integration{}
	require.NoError(t, faker.FakeObject(&i))
	m.EXPECT().GetIntegrations(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigatewayv2.GetIntegrationsOutput{
			Items: []types.Integration{i},
		}, nil)

	ir := types.IntegrationResponse{}
	require.NoError(t, faker.FakeObject(&ir))
	m.EXPECT().GetIntegrationResponses(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigatewayv2.GetIntegrationResponsesOutput{
			Items: []types.IntegrationResponse{ir},
		}, nil)

	am := types.Model{}
	require.NoError(t, faker.FakeObject(&am))
	m.EXPECT().GetModels(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigatewayv2.GetModelsOutput{
			Items: []types.Model{am},
		}, nil)

	mt := apigatewayv2.GetModelTemplateOutput{}
	require.NoError(t, faker.FakeObject(&mt))
	m.EXPECT().GetModelTemplate(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&mt, nil)

	r := types.Route{}
	require.NoError(t, faker.FakeObject(&r))
	m.EXPECT().GetRoutes(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigatewayv2.GetRoutesOutput{
			Items: []types.Route{r},
		}, nil)

	s := types.Stage{}
	require.NoError(t, faker.FakeObject(&s))
	m.EXPECT().GetStages(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigatewayv2.GetStagesOutput{
			Items: []types.Stage{s},
		}, nil)

	rr := types.RouteResponse{}
	require.NoError(t, faker.FakeObject(&rr))
	m.EXPECT().GetRouteResponses(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigatewayv2.GetRouteResponsesOutput{
			Items: []types.RouteResponse{rr},
		}, nil)

	return client.Services{
		Apigatewayv2: m,
	}
}

func TestApigatewayv2Apis(t *testing.T) {
	client.AwsMockTestHelper(t, Apis(), buildApigatewayv2Apis, client.TestOptions{})
}
