    var response api.{{ .StructName }}
	require.NoError(t, faker.FakeObject(&response))
{{- if .ValueField }}
    // Use correct Azure ID format
	const id = "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	response.{{ .ValueField }}[0].ID = to.Ptr(id)
{{- end }}

    mock{{ .Client }}.EXPECT().{{ .Method }}({{- range .Params }}gomock.Any(), {{ end }}gomock.Any()).
        Return(response, nil).MinTimes(1)