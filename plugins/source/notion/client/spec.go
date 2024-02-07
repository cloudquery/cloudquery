package client

type Spec struct {
	BearerToken   string `json:"bearer_token,omitempty" jsonschema_extras:"x-cq-auth=true"`
	NotionVersion string `json:"notion_version,omitempty"`
}

func (s *Spec) SetDefaults() {
	if len(s.NotionVersion) < 1 {
		s.NotionVersion = "2022-02-22"
	}
}
