package lambda

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/aws/smithy-go"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func buildLambdaFunctionsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockLambdaClient(ctrl)

	lastModified := "1994-11-05T08:15:30.000+0500"

	fc := types.FunctionConfiguration{}
	require.NoError(t, faker.FakeObject(&fc))

	fc2 := types.FunctionConfiguration{}
	require.NoError(t, faker.FakeObject(&fc2))
	fc2.FunctionArn = aws.String("arn:aws:lambda:us-east-1:123456789012:function:my-function:2")
	m.EXPECT().ListFunctions(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(&lambda.ListFunctionsOutput{Functions: []types.FunctionConfiguration{fc, fc2}}, nil)

	f := lambda.GetFunctionOutput{}
	require.NoError(t, faker.FakeObject(&f))
	f.Configuration.LastModified = &lastModified
	err := smithy.GenericAPIError{Code: "AccessDenied", Message: "This is an error message"}

	// There are 2 calls to GetFunction, one succeeds and the other fails
	gomock.InOrder(
		m.EXPECT().GetFunction(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(&f, nil),

		m.EXPECT().GetFunction(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil, &err),
	)
	a := types.AliasConfiguration{}
	require.NoError(t, faker.FakeObject(&a))
	m.EXPECT().ListAliases(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(&lambda.ListAliasesOutput{Aliases: []types.AliasConfiguration{a}}, nil).AnyTimes()

	urlConfig := types.FunctionUrlConfig{}
	require.NoError(t, faker.FakeObject(&urlConfig))
	urlConfig.CreationTime = aws.String("2012-07-14T01:00:00+01:00")
	urlConfig.LastModifiedTime = aws.String("2012-07-14T01:00:00+01:00")
	m.EXPECT().ListFunctionUrlConfigs(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(&lambda.ListFunctionUrlConfigsOutput{FunctionUrlConfigs: []types.FunctionUrlConfig{urlConfig}}, nil).AnyTimes()

	gfco := lambda.GetFunctionConcurrencyOutput{}
	require.NoError(t, faker.FakeObject(&gfco))
	m.EXPECT().GetFunctionConcurrency(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(&gfco, nil).AnyTimes()

	lto := lambda.ListTagsOutput{}
	require.NoError(t, faker.FakeObject(&lto))
	m.EXPECT().ListTags(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(&lto, nil).AnyTimes()

	i := types.FunctionEventInvokeConfig{}
	require.NoError(t, faker.FakeObject(&i))
	m.EXPECT().ListFunctionEventInvokeConfigs(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(&lambda.ListFunctionEventInvokeConfigsOutput{FunctionEventInvokeConfigs: []types.FunctionEventInvokeConfig{i}}, nil).AnyTimes()

	cc := types.ProvisionedConcurrencyConfigListItem{}
	require.NoError(t, faker.FakeObject(&cc))
	cc.LastModified = &lastModified
	m.EXPECT().ListProvisionedConcurrencyConfigs(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(&lambda.ListProvisionedConcurrencyConfigsOutput{ProvisionedConcurrencyConfigs: []types.ProvisionedConcurrencyConfigListItem{cc}}, nil).AnyTimes()

	esm := types.EventSourceMappingConfiguration{}
	require.NoError(t, faker.FakeObject(&esm))
	esm.UUID = aws.String(uuid.NewString())
	m.EXPECT().ListEventSourceMappings(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(&lambda.ListEventSourceMappingsOutput{EventSourceMappings: []types.EventSourceMappingConfiguration{esm}}, nil).AnyTimes()

	fp := lambda.GetPolicyOutput{}
	require.NoError(t, faker.FakeObject(&fp))
	document := "{\"test\":1}"
	fp.Policy = &document
	m.EXPECT().GetPolicy(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(&fp, nil).AnyTimes()

	csco := lambda.GetFunctionCodeSigningConfigOutput{}
	require.NoError(t, faker.FakeObject(&csco))
	m.EXPECT().GetFunctionCodeSigningConfig(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(&csco, nil).AnyTimes()

	csc := types.CodeSigningConfig{}
	require.NoError(t, faker.FakeObject(&csc))
	isoDate := "2011-10-05T14:48:00.000Z"
	csc.LastModified = &isoDate
	m.EXPECT().GetCodeSigningConfig(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(&lambda.GetCodeSigningConfigOutput{CodeSigningConfig: &csc}, nil).AnyTimes()

	fc.LastModified = &lastModified
	m.EXPECT().ListVersionsByFunction(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(&lambda.ListVersionsByFunctionOutput{Versions: []types.FunctionConfiguration{fc}}, nil).AnyTimes()

	runtimeManagementConfig := lambda.GetRuntimeManagementConfigOutput{}
	require.NoError(t, faker.FakeObject(&runtimeManagementConfig))
	m.EXPECT().GetRuntimeManagementConfig(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(&runtimeManagementConfig, nil).AnyTimes()

	return client.Services{
		Lambda: m,
	}
}

func TestLambdaFunctions(t *testing.T) {
	client.AwsMockTestHelper(t, Functions(), buildLambdaFunctionsMock, client.TestOptions{})
}
