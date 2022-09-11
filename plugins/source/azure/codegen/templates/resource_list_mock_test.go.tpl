{{template "mock_test_base.go.tpl" .}}

func create{{ .AzureSubService }}Mock(t *testing.T, ctrl *gomock.Controller) services.Services {
	{{template "mock_test_setup.go.tpl" .}}

	{{if .MockListResult }}
	{{if eq .MockListResult "CQ_CODEGEN_DIRECT_RESPONSE" }}
	result := data
	{{else}}
    result := {{ .AzurePackageName }}.New{{ .MockListResult }}Page({{ .AzurePackageName }}.{{ .MockListResult }}{Value: &[]{{ .AzurePackageName }}.{{ or .MockValueType .AzureStructName }}{data}}, func(ctx context.Context, result {{ .AzurePackageName }}.{{ .MockListResult }}) ({{ .AzurePackageName }}.{{ .MockListResult }}, error) {
		return {{ .AzurePackageName }}.{{ .MockListResult }}{}, nil
	})
	{{end}}
	{{else}}
	result := {{ .AzurePackageName }}.New{{ .AzureStructName }}{{ "ListResult" }}Page({{ .AzurePackageName }}.{{ .AzureStructName }}{{ "ListResult" }}{Value: &[]{{ .AzurePackageName }}.{{ or .MockValueType .AzureStructName }}{data}}, func(ctx context.Context, result {{ .AzurePackageName }}.{{ .AzureStructName }}{{ "ListResult" }}) ({{ .AzurePackageName }}.{{ .AzureStructName }}{{ "ListResult" }}, error) {
		return {{ .AzurePackageName }}.{{ .AzureStructName }}{{ "ListResult" }}{}, nil
	})
	{{end}}

	{{template "mock_test_assert.go.tpl" .}}
}