    {{ range or .MockListFunctionArgsInit .ListFunctionArgsInit }}
	{{.}}{{ end }}
	mockClient.EXPECT().{{ or .ListFunction "ListAll" }}(gomock.Any(){{ range or .MockListFunctionArgs .ListFunctionArgs }}, {{.}}{{ end }}).Return(result, nil)
	{{if .GetFunction}}
	mockClient.EXPECT().{{ or .GetFunction "Get" }}(gomock.Any(){{ range or .MockGetFunctionArgs .GetFunctionArgs }}, {{.}}{{ end }}).Return(getData, nil)
	{{end}}return s