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

	page := {{ .AzurePackageName }}.New{{ .AzureStructName }}{{ or .MockListResult "ListResult" }}Page({{ .AzurePackageName }}.{{ .AzureStructName }}{{ or .MockListResult "ListResult" }}{Value: &[]{{ .AzurePackageName }}.{{ .AzureStructName }}{data}}, func(ctx context.Context, result {{ .AzurePackageName }}.{{ .AzureStructName }}{{ or .MockListResult "ListResult" }}) ({{ .AzurePackageName }}.{{ .AzureStructName }}{{ or .MockListResult "ListResult" }}, error) {
		return {{ .AzurePackageName }}.{{ .AzureStructName }}{{ or .MockListResult "ListResult" }}{}, nil
	})

	{{ range .ListFunctionArgsInit }}
	{{.}}{{ end }}
	mockClient.EXPECT().{{ or .ListFunction "ListAll" }}(gomock.Any(){{ range or .MockListFunctionArgs .ListFunctionArgs }}, {{.}}{{ end }}).Return(page, nil)
	return s
}