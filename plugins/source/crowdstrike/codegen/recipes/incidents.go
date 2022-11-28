package recipes

import (
	"github.com/crowdstrike/gofalcon/falcon/models"
)

func CrowdScore() []*Resource {
	resources := []*Resource{
		{
			Service:    "incidents",
			SubService: "crowdscore",
			Struct:     &models.APIMsaEnvironmentScoreResponse{},
		},
	}
	return resources
}
