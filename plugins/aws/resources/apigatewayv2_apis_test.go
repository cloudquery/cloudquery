package resources

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"

	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildApigatewayv2Apis(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockApigatewayv2Client(ctrl)

	a := types.Api{}
	err := faker.FakeData(&a)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetApis(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigatewayv2.GetApisOutput{
			Items: []types.Api{a},
		}, nil)

	aa := types.Authorizer{}
	err = faker.FakeData(&aa)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetAuthorizers(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigatewayv2.GetAuthorizersOutput{
			Items: []types.Authorizer{aa},
		}, nil)

	d := types.Deployment{}
	err = faker.FakeData(&d)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetDeployments(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigatewayv2.GetDeploymentsOutput{
			Items: []types.Deployment{d},
		}, nil)

	i := types.Integration{}
	err = faker.FakeData(&i)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetIntegrations(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigatewayv2.GetIntegrationsOutput{
			Items: []types.Integration{i},
		}, nil)

	ir := types.IntegrationResponse{}
	err = faker.FakeData(&ir)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetIntegrationResponses(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigatewayv2.GetIntegrationResponsesOutput{
			Items: []types.IntegrationResponse{ir},
		}, nil)

	am := types.Model{}
	err = faker.FakeData(&am)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetModels(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigatewayv2.GetModelsOutput{
			Items: []types.Model{am},
		}, nil)

	mt := apigatewayv2.GetModelTemplateOutput{}
	err = faker.FakeData(&mt)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetModelTemplate(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&mt, nil)

	r := types.Route{}
	err = faker.FakeData(&r)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetRoutes(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigatewayv2.GetRoutesOutput{
			Items: []types.Route{r},
		}, nil)

	s := types.Stage{}
	err = faker.FakeData(&s)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetStages(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigatewayv2.GetStagesOutput{
			Items: []types.Stage{s},
		}, nil)

	rr := types.RouteResponse{}
	err = faker.FakeData(&rr)
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
	awsTestHelper(t, Apigatewayv2Apis(), buildApigatewayv2Apis)
}
