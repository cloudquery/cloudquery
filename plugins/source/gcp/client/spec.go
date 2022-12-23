package client

// Spec defines GCP source plugin Spec
type Spec struct {
	ProjectIDs            []string `json:"project_ids"`
	ServiceAccountKeyJSON string   `json:"service_account_key_json"`
	FolderIDs             []string `json:"folder_ids"`
	FolderRecursionDepth  *int     `json:"folder_recursion_depth"`
	ProjectFilter         string   `json:"project_filter"`
	BackoffDelay          int      `json:"backoff_delay"`
	BackoffRetries        int      `json:"backoff_retries"`
	EnabledServicesOnly   bool     `json:"enabled_services_only"`
}

func (spec *Spec) setDefaults() {
	var defaultRecursionDepth = 100
	if spec.FolderRecursionDepth == nil {
		spec.FolderRecursionDepth = &defaultRecursionDepth
	}
}
