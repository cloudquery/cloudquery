package client

// Provider Configuration

type Config struct {
	Token string `hcl:"token"`
	// SpacesRegions is a list of DO regions to fetch spaces from, if not given we execute on all regions
	SpacesRegions []string `hcl:"spaces_regions,optional"`
	// SpacesAccessKey is the secret access token generated in DO control panel
	SpacesAccessKey string `hcl:"spaces_access_key,optional"`
	// SpacesAccessKeyId is the unique identifier of the access key generated in the DO control panel
	SpacesAccessKeyId string `hcl:"spaces_access_key_id,optional"`
	// SpacesDebugLogging allows enabling AWS S3 request logging on spaces requests
	SpacesDebugLogging bool `hcl:"spaces_debug_logging,optional"`
}

func (c Config) Example() string {
	return `
		configuration {
			// API Token to access DigialOcean resources 
			// See https://docs.digitalocean.com/reference/api/api-reference/#section/Authentication
			token = "<YOUR_API_TOKEN_HERE>"
			// List of regions to fetch spaces from, if not given all regions are assumed
			// spaces_regions = ["nyc3", "sfo3", "ams3", "sgp1", "fra1"]
			// Spaces Access Key generated at https://cloud.digitalocean.com/settings/api/tokens
			// spaces_access_key = "<YOUR_SPACES_ACCESS_KEY>"
			// Spaces Access Key Id generated at https://cloud.digitalocean.com/settings/api/tokens
			// spaces_access_key_id = "<YOUR_SPACES_ACCESS_KEY_ID>"
			// SpacesDebugLogging allows enabling AWS S3 request logging on spaces requests
			// spaces_debug_logging = false
		}
`
}
