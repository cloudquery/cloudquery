Generated from https://github.com/Azure/azure-rest-api-specs/tree/3c764635e7d442b3e74caf593029fcd440b3ef82

Code generator @microsoft.azure/autorest.go@~2.1.161

## Breaking Changes

- Function `NewProviderListResultPage` parameter(s) have been changed from `(func(context.Context, ProviderListResult) (ProviderListResult, error))` to `(ProviderListResult, func(context.Context, ProviderListResult) (ProviderListResult, error))`
- Function `NewDeploymentOperationsListResultPage` parameter(s) have been changed from `(func(context.Context, DeploymentOperationsListResult) (DeploymentOperationsListResult, error))` to `(DeploymentOperationsListResult, func(context.Context, DeploymentOperationsListResult) (DeploymentOperationsListResult, error))`
- Function `NewDeploymentListResultPage` parameter(s) have been changed from `(func(context.Context, DeploymentListResult) (DeploymentListResult, error))` to `(DeploymentListResult, func(context.Context, DeploymentListResult) (DeploymentListResult, error))`
- Function `NewGroupListResultPage` parameter(s) have been changed from `(func(context.Context, GroupListResult) (GroupListResult, error))` to `(GroupListResult, func(context.Context, GroupListResult) (GroupListResult, error))`
- Function `NewOperationListResultPage` parameter(s) have been changed from `(func(context.Context, OperationListResult) (OperationListResult, error))` to `(OperationListResult, func(context.Context, OperationListResult) (OperationListResult, error))`
- Function `NewTagsListResultPage` parameter(s) have been changed from `(func(context.Context, TagsListResult) (TagsListResult, error))` to `(TagsListResult, func(context.Context, TagsListResult) (TagsListResult, error))`
- Function `NewListResultPage` parameter(s) have been changed from `(func(context.Context, ListResult) (ListResult, error))` to `(ListResult, func(context.Context, ListResult) (ListResult, error))`

## New Content

- New const `ExpressionEvaluationOptionsScopeTypeOuter`
- New const `ExpressionEvaluationOptionsScopeTypeInner`
- New const `ExpressionEvaluationOptionsScopeTypeNotSpecified`
- New function `ProvidersClient.RegisterAtManagementGroupScope(context.Context, string, string) (autorest.Response, error)`
- New function `ProvidersClient.RegisterAtManagementGroupScopeSender(*http.Request) (*http.Response, error)`
- New function `ProvidersClient.RegisterAtManagementGroupScopePreparer(context.Context, string, string) (*http.Request, error)`
- New function `PossibleExpressionEvaluationOptionsScopeTypeValues() []ExpressionEvaluationOptionsScopeType`
- New function `ProvidersClient.RegisterAtManagementGroupScopeResponder(*http.Response) (autorest.Response, error)`
- New struct `ExpressionEvaluationOptions`
- New field `ExpressionEvaluationOptions` in struct `DeploymentWhatIfProperties`
- New field `ExpressionEvaluationOptions` in struct `DeploymentProperties`
