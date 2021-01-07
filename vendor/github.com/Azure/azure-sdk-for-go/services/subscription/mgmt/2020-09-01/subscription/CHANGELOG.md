Generated from https://github.com/Azure/azure-rest-api-specs/tree/3c764635e7d442b3e74caf593029fcd440b3ef82

Code generator @microsoft.azure/autorest.go@~2.1.161

## Breaking Changes

- Function `NewListResultPage` parameter(s) have been changed from `(func(context.Context, ListResult) (ListResult, error))` to `(ListResult, func(context.Context, ListResult) (ListResult, error))`
- Function `NewTenantListResultPage` parameter(s) have been changed from `(func(context.Context, TenantListResult) (TenantListResult, error))` to `(TenantListResult, func(context.Context, TenantListResult) (TenantListResult, error))`

## New Content

- New field `ResellerID` in struct `PutAliasRequestProperties`
