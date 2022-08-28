    mockClient := mocks.NewMock{{ .AzureService }}{{ .AzureSubService }}Client(ctrl)
	s := services.Services{
		{{ .AzureService }}: services.{{ .AzureService }}Client{
			{{ .AzureSubService }}: mockClient,
		},
	}

	data := {{ .AzurePackageName }}.{{ .AzureStructName }}{}
	fieldsToIgnore := []string{{"{"}}{{ range .MockFieldsToIgnore }}"{{.}}",{{ end }}{{"}"}}
	require.Nil(t, faker.FakeData(&data, options.WithIgnoreInterface(true), options.WithFieldsToIgnore(fieldsToIgnore...), options.WithRandomMapAndSliceMinSize(1), options.WithRandomMapAndSliceMaxSize(1)))