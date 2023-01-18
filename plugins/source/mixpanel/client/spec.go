package client

type Spec struct {
	Username string `json:"username,omitempty"`
	Secret   string `json:"secret,omitempty"`

	ProjectID   int64 `json:"project_id,omitempty"`
	WorkspaceID int64 `json:"workspace_id,omitempty"`

	// Optional
	BaseURL    string `json:"base_url,omitempty"`
	Timeout    int64  `json:"timeout_secs,omitempty"`
	MaxRetries int64  `json:"max_retries,omitempty"`
	PageSize   int64  `json:"page_size,omitempty"`
}
