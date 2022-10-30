package apigatewayv2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildApigatewayv2Apis(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockApigatewayv2Client(ctrl)

	a := types.Api{}
	err := faker.FakeObject(&a)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetApis(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigatewayv2.GetApisOutput{
			Items: []types.Api{a},
		}, nil)

	aa := types.Authorizer{}
	err = faker.FakeObject(&aa)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetAuthorizers(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigatewayv2.GetAuthorizersOutput{
			Items: []types.Authorizer{aa},
		}, nil)

	d := types.Deployment{}
	err = faker.FakeObject(&d)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetDeployments(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigatewayv2.GetDeploymentsOutput{
			Items: []types.Deployment{d},
		}, nil)

	i := types.Integration{}
	err = faker.FakeObject(&i)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetIntegrations(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigatewayv2.GetIntegrationsOutput{
			Items: []types.Integration{i},
		}, nil)

	ir := types.IntegrationResponse{}
	err = faker.FakeObject(&ir)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetIntegrationResponses(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigatewayv2.GetIntegrationResponsesOutput{
			Items: []types.IntegrationResponse{ir},
		}, nil)

	am := types.Model{}
	err = faker.FakeObject(&am)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetModels(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigatewayv2.GetModelsOutput{
			Items: []types.Model{am},
		}, nil)

	mt := apigatewayv2.GetModelTemplateOutput{}
	err = faker.FakeObject(&mt)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetModelTemplate(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&mt, nil)

	r := types.Route{}
	err = faker.FakeObject(&r)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetRoutes(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigatewayv2.GetRoutesOutput{
			Items: []types.Route{r},
		}, nil)

	s := types.Stage{}
	err = faker.FakeObject(&s)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetStages(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigatewayv2.GetStagesOutput{
			Items: []types.Stage{s},
		}, nil)

	rr := types.RouteResponse{}
	err = faker.FakeObject(&rr)
	if err != nil {
		t.Fatal(err)
	}
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
