{{template "mock_test_base.go.tpl" .}}

func create{{ .AzureSubService }}Mock(t *testing.T, ctrl *gomock.Controller) services.Services {
	{{template "mock_test_setup.go.tpl" .}}

	pager := runtime.NewPager(runtime.PagingHandler[{{ .AzurePackageName }}.{{ .MockListResult }}]{
		More: func(page {{ .AzurePackageName }}.{{ .MockListResult }}) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *{{ .AzurePackageName }}.{{ .MockListResult }}) ({{ .AzurePackageName }}.{{ .MockListResult }}, error) {
			return {{ .AzurePackageName }}.{{ .MockListResult }}{
				{{ .AzureSubService | ToSingular }}ListResult: {{ .AzurePackageName }}.{{ .AzureSubService | ToSingular }}ListResult{
					Value:    []*{{ .AzurePackageName }}.{{ .AzureStructName }}{&data},
				},
			}, nil
		},
	})

	{{ range or .MockListFunctionArgsInit .ListFunctionArgsInit }}
	{{.}}{{ end }}
	mockClient.EXPECT().{{ or .ListFunction "NewListPager" }}(gomock.Any(){{ range or .MockListFunctionArgs .ListFunctionArgs }}, {{.}}{{ end }}).Return(
		pager,
	)
	return s
}