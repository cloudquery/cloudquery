{{template "base_mock_test.go.tpl" .}}

func create{{ .AzureService }}{{ .AzureSubService }}Mock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMock{{ .AzureSubService }}Client(ctrl)
	s := services.Services{
		{{ .AzureService }}: services.{{ .AzureService }}Client{
			{{ .AzureSubService }}: mockClient,
		},
	}

	data := {{ .AzurePackageName }}.{{ .AzureStructName }}{}
	require.Nil(t, faker.FakeData(&data))

	{{if .MockListResult}}
    page := {{ .AzurePackageName }}.New{{ .MockListResult }}Page({{ .AzurePackageName }}.{{ .MockListResult }}{Value: &[]{{ .AzurePackageName }}.{{ .AzureStructName }}{data}}, func(ctx context.Context, result {{ .AzurePackageName }}.{{ .MockListResult }}) ({{ .AzurePackageName }}.{{ .MockListResult }}, error) {
		return {{ .AzurePackageName }}.{{ .MockListResult }}{}, nil
	})
	{{else}}
	page := {{ .AzurePackageName }}.New{{ .AzureStructName }}{{ "ListResult" }}Page({{ .AzurePackageName }}.{{ .AzureStructName }}{{ "ListResult" }}{Value: &[]{{ .AzurePackageName }}.{{ .AzureStructName }}{data}}, func(ctx context.Context, result {{ .AzurePackageName }}.{{ .AzureStructName }}{{ "ListResult" }}) ({{ .AzurePackageName }}.{{ .AzureStructName }}{{ "ListResult" }}, error) {
		return {{ .AzurePackageName }}.{{ .AzureStructName }}{{ "ListResult" }}{}, nil
	})
	{{end}}

	{{ range or .MockListFunctionArgsInit .ListFunctionArgsInit }}
	{{.}}{{ end }}
	mockClient.EXPECT().{{ or .ListFunction "ListAll" }}(gomock.Any(){{ range or .MockListFunctionArgs .ListFunctionArgs }}, {{.}}{{ end }}).Return(page, nil)
	return s
}