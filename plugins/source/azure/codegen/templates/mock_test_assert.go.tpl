    {{ range or .MockListFunctionArgsInit .ListFunctionArgsInit }}
	{{.}}{{ end }}
	mockClient.EXPECT().{{ or .ListFunction "ListAll" }}(gomock.Any(){{ range or .MockListFunctionArgs .ListFunctionArgs }}, {{.}}{{ end }}).Return(result, nil)
	return s