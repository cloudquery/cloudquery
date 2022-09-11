    mockClient := mocks.NewMock{{ .AzureService }}{{ .AzureSubService }}Client(ctrl)
	s := services.Services{
		{{ .AzureService }}: services.{{ .AzureService }}Client{
			{{ .AzureSubService }}: mockClient,
		},
	}

	data := {{ .AzurePackageName }}.{{ or .MockDefinitionType .AzureStructName }}{}
	fieldsToIgnore := []string{{"{"}}{{ range .MockFieldsToIgnore }}"{{.}}",{{ end }}{{"}"}}
	require.Nil(t, faker.FakeData(&data, fakerOptions.WithIgnoreInterface(true), fakerOptions.WithRecursionMaxDepth(2), fakerOptions.WithFieldsToIgnore(fieldsToIgnore...), fakerOptions.WithRandomMapAndSliceMinSize(1), fakerOptions.WithRandomMapAndSliceMaxSize(1)))