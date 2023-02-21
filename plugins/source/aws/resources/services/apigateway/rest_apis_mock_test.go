package apigateway

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildApigatewayRestApis(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockApigatewayClient(ctrl)

	r := types.RestApi{}
	err := faker.FakeObject(&r)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetRestApis(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigateway.GetRestApisOutput{
			Items: []types.RestApi{r},
		}, nil)

	a := types.Authorizer{}
	err = faker.FakeObject(&a)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetAuthorizers(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigateway.GetAuthorizersOutput{
			Items: []types.Authorizer{a},
		}, nil)

	d := types.Deployment{}
	err = faker.FakeObject(&d)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetDeployments(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigateway.GetDeploymentsOutput{
			Items: []types.Deployment{d},
		}, nil)

	dp := types.DocumentationPart{}
	err = faker.FakeObject(&dp)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetDocumentationParts(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigateway.GetDocumentationPartsOutput{
			Items: []types.DocumentationPart{dp},
		}, nil)

	dv := types.DocumentationVersion{}
	err = faker.FakeObject(&dv)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetDocumentationVersions(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigateway.GetDocumentationVersionsOutput{
			Items: []types.DocumentationVersion{dv},
		}, nil)

	gr := types.GatewayResponse{}
	err = faker.FakeObject(&gr)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetGatewayResponses(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigateway.GetGatewayResponsesOutput{
			Items: []types.GatewayResponse{gr},
		}, nil)

	am := types.Model{}
	err = faker.FakeObject(&am)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetModels(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigateway.GetModelsOutput{
			Items: []types.Model{am},
		}, nil)

	mt := apigateway.GetModelTemplateOutput{}
	err = faker.FakeObject(&mt)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetModelTemplate(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&mt, nil)

	rv := types.RequestValidator{}
	err = faker.FakeObject(&rv)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetRequestValidators(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigateway.GetRequestValidatorsOutput{
			Items: []types.RequestValidator{rv},
		}, nil)

	ar := types.Resource{}
	err = faker.FakeObject(&ar)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetResources(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigateway.GetResourcesOutput{
			Items: []types.Resource{ar},
		}, nil)

	method := apigateway.GetMethodOutput{}
	err = faker.FakeObject(&method)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetMethod(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&method, nil)

	integration := apigateway.GetIntegrationOutput{}
	err = faker.FakeObject(&integration)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetIntegration(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&integration, nil)

	s := types.Stage{}
	err = faker.FakeObject(&s)
	if err != nil {
		t.Fatal(err)
	}
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
