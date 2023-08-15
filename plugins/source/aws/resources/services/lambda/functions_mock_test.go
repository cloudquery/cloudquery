package lambda

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildLambdaFunctionsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockLambdaClient(ctrl)

	lastModified := "1994-11-05T08:15:30.000+0500"

	f := lambda.GetFunctionOutput{}
	require.NoError(t, faker.FakeObject(&f))
	f.Configuration.LastModified = &lastModified
	m.EXPECT().GetFunction(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&f, nil)

	fc := types.FunctionConfiguration{}
	require.NoError(t, faker.FakeObject(&fc))
	m.EXPECT().ListFunctions(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&lambda.ListFunctionsOutput{
			Functions: []types.FunctionConfiguration{fc},
		}, nil)

	a := types.AliasConfiguration{}
	require.NoError(t, faker.FakeObject(&a))
	m.EXPECT().ListAliases(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&lambda.ListAliasesOutput{
			Aliases: []types.AliasConfiguration{a},
		}, nil)

	i := types.FunctionEventInvokeConfig{}
	require.NoError(t, faker.FakeObject(&i))
	m.EXPECT().ListFunctionEventInvokeConfigs(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&lambda.ListFunctionEventInvokeConfigsOutput{
			FunctionEventInvokeConfigs: []types.FunctionEventInvokeConfig{i},
		}, nil)

	cc := types.ProvisionedConcurrencyConfigListItem{}
	require.NoError(t, faker.FakeObject(&cc))
	cc.LastModified = &lastModified
	m.EXPECT().ListProvisionedConcurrencyConfigs(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&lambda.ListProvisionedConcurrencyConfigsOutput{
			ProvisionedConcurrencyConfigs: []types.ProvisionedConcurrencyConfigListItem{cc},
		}, nil)

	esm := types.EventSourceMappingConfiguration{}
	require.NoError(t, faker.FakeObject(&esm))
	m.EXPECT().ListEventSourceMappings(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&lambda.ListEventSourceMappingsOutput{
			EventSourceMappings: []types.EventSourceMappingConfiguration{esm},
		}, nil)

	fp := lambda.GetPolicyOutput{}
	require.NoError(t, faker.FakeObject(&fp))
	document := "{\"test\":1}"
	fp.Policy = &document
	m.EXPECT().GetPolicy(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&fp, nil)

	csco := lambda.GetFunctionCodeSigningConfigOutput{}
	require.NoError(t, faker.FakeObject(&csco))
	m.EXPECT().GetFunctionCodeSigningConfig(gomock.Any(), gomock.Any(), gomock.Any()).Return(&csco, nil)

	csc := types.CodeSigningConfig{}
	require.NoError(t, faker.FakeObject(&csc))
	isoDate := "2011-10-05T14:48:00.000Z"
	csc.LastModified = &isoDate
	m.EXPECT().GetCodeSigningConfig(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&lambda.GetCodeSigningConfigOutput{
			CodeSigningConfig: &csc,
		}, nil)
	fc.LastModified = &lastModified
	m.EXPECT().ListVersionsByFunction(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&lambda.ListVersionsByFunctionOutput{
			Versions: []types.FunctionConfiguration{fc},
		}, nil)

	urlConfig := lambda.GetFunctionUrlConfigOutput{}
	require.NoError(t, faker.FakeObject(&urlConfig))
	urlConfig.CreationTime = aws.String("2012-07-14T01:00:00+01:00")
	urlConfig.LastModifiedTime = aws.String("2012-07-14T01:00:00+01:00")
	m.EXPECT().GetFunctionUrlConfig(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&urlConfig, nil)

	runtimeManagementConfig := lambda.GetRuntimeManagementConfigOutput{}
	require.NoError(t, faker.FakeObject(&runtimeManagementConfig))
	m.EXPECT().GetRuntimeManagementConfig(gomock.Any(), gomock.Any(), gomock.Any()).Return(&runtimeManagementConfig, nil)

	return client.Services{
		Lambda: m,
	}
}

func TestLambdaFunctions(t *testing.T) {
	client.AwsMockTestHelper(t, Functions(), buildLambdaFunctionsMock, client.TestOptions{})
}
