{{template "mock_test_base.go.tpl" .}}

func create{{ .AzureSubService }}Mock(t *testing.T, ctrl *gomock.Controller) services.Services {
	{{template "mock_test_setup.go.tpl" .}}

	pager := runtime.NewPager(runtime.PagingHandler[{{ .AzurePackageName }}.{{ or .MockListResult .AzureSubService }}ClientListResponse]{
		More: func(page {{ .AzurePackageName }}.{{ or .MockListResult .AzureSubService }}ClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *{{ .AzurePackageName }}.{{ or .MockListResult .AzureSubService }}ClientListResponse) ({{ .AzurePackageName }}.{{ or .MockListResult .AzureSubService }}ClientListResponse, error) {
			return {{ .AzurePackageName }}.{{ or .MockListResult .AzureSubService }}ClientListResponse{
				{{ .AzureSubService | ToSingular }}ListResult: {{ .AzurePackageName }}.{{ .AzureSubService | ToSingular }}ListResult{
					NextLink: nil,
					Value:    []*{{ .AzurePackageName }}.{{ .AzureStructName }}{&data},
				},
			}, nil
		},
	})

	mockClient.EXPECT().NewListPager(gomock.Any()).Return(
		pager,
	)
	return s
}