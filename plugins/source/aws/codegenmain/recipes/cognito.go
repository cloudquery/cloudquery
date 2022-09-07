package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/cloudquery/plugin-sdk/codegen"
)

func init() {
	add(combine(&Resource{
		DefaultColumns:             []codegen.ColumnDefinition{AccountIdColumn, RegionColumn},
		AWSStruct:                  &cognitoidentity.DescribeIdentityPoolOutput{},
		AWSService:                 "Cognito",
		AWSServiceClient:           "CognitoIdentityPools",
		Template:                   "resource_list_describe",
		MultiplexerServiceOverride: "cognito-identity",
		PaginatorStruct:            &cognitoidentity.ListIdentityPoolsOutput{},
		PaginatorGetStruct:         &cognitoidentity.DescribeIdentityPoolInput{},
		ItemsStruct:                &cognitoidentity.DescribeIdentityPoolOutput{},
		PrimaryKeys:                []string{"account_id", "id"},
		SkipFields:                 []string{"ResultMetadata"},
		CustomInputs: []string{
			`MaxResults: 60, // we want max results to reduce List calls as much as possible, services limited to less than or equal to 60`,
		},
		Imports: []string{
			`cognito "github.com/aws/aws-sdk-go-v2/service/cognitoidentity"`,
		},
		MockImports: []string{
			`cognito "github.com/aws/aws-sdk-go-v2/service/cognitoidentity"`,
			`"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"`,
			`github.com/aws/aws-sdk-go-v2/service/cognitoidentity/types`,
		},
	},
		parentize(&Resource{
			DefaultColumns:             []codegen.ColumnDefinition{AccountIdColumn, RegionColumn},
			AWSService:                 "Cognito",
			AWSServiceClient:           "CognitoUserPools",
			AWSStruct:                  &types.UserPoolType{},
			Template:                   "resource_list_describe",
			MultiplexerServiceOverride: "cognito-idp",
			PaginatorStruct:            &cognitoidentityprovider.ListUserPoolsOutput{},
			PaginatorGetStruct:         &cognitoidentityprovider.DescribeUserPoolInput{},
			ItemsStruct:                &cognitoidentityprovider.DescribeUserPoolOutput{},
			PrimaryKeys:                []string{"account_id", "id"},
			CustomInputs: []string{
				`MaxResults: 60, // we want max results to reduce List calls as much as possible, services limited to less than or equal to 60`,
			},
			Imports: []string{
				`cognito "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"`,
			},
			MockImports: []string{
				`cognito "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"`,
				`github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types`,
			},
			SkipMainImport: true,
			ItemName:       "UserPool",
		},
			&Resource{
				AWSStruct:          &types.IdentityProviderType{},
				Template:           "resource_list_describe",
				PaginatorStruct:    &cognitoidentityprovider.ListIdentityProvidersOutput{},
				PaginatorGetStruct: &cognitoidentityprovider.DescribeIdentityProviderInput{},
				ItemsStruct:        &cognitoidentityprovider.DescribeIdentityProviderOutput{},
				ParentFieldName:    "Id",
				ChildFieldName:     "UserPoolId",
				Imports: []string{
					`cognito "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"`,
				},
				MockImports: []string{
					`cognito "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"`,
					`github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types`,
				},
				SkipMainImport: true,
				ItemName:       "IdentityProvider",
			},
		),
	)...)
}
