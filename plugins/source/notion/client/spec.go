package client

import _ "embed"

type Spec struct {
	BearerToken   string `json:"bearer_token,omitempty" jsonschema:"required,minLength=1"`
	NotionVersion string `json:"notion_version,omitempty" jsonschema:"minLength=1,default=2022-02-22"`
}

//go:embed schema.json
var JSONSchema string

func (s *Spec) SetDefaults() {
	if len(s.NotionVersion) < 1 {
		s.NotionVersion = "2022-02-22"
	}
}
