package client

// Spec defines DigitalOcean source plugin Spec
type Spec struct {
	Token string `json:"token,omitempty"`
	// SpacesRegions is a list of DO regions to fetch spaces from, if not given we execute on all regions
	SpacesRegions []string `json:"spaces_regions,omitempty"`
	// SpacesAccessKey is the secret access token generated in DO control panel
	SpacesAccessKey string `json:"spaces_access_key,omitempty"`
	// SpacesAccessKeyId is the unique identifier of the access key generated in the DO control panel
	SpacesAccessKeyId string `json:"spaces_access_key_id,omitempty"`
	// SpacesDebugLogging allows enabling AWS S3 request logging on spaces requests
	SpacesDebugLogging bool `json:"spaces_debug_logging,omitempty"`
}
