package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemovePII(t *testing.T) {
	tests := []struct {
		input      string
		sensitives []string
		expected   string
	}{
		{
			input: `table resolver ignored error: googleapi: Error 403: Cloud Functions API has not been used in project 12345678 before or it is disabled. Enable it by visiting https://console.developers.google.com/apis/api/cloudfunctions.googleapis.com/overview?project=12345678 then retry. If you enabled this API recently, wait a few minutes for the action to propagate to our systems and retry.
Details:
[
  {
    "@type": "type.googleapis.com/google.rpc.Help",
    "links": [
      {
        "description": "Google developers console API activation",
        "url": "https://console.developers.google.com/apis/api/cloudfunctions.googleapis.com/overview?project=12345678"
      }
    ]
  },
  {
    "@type": "type.googleapis.com/google.rpc.ErrorInfo",
    "domain": "googleapis.com",
    "metadata": {
      "consumer": "projects/12345678",
      "service": "cloudfunctions.googleapis.com"
    },
    "reason": "SERVICE_DISABLED"
  }
]
, accessNotConfigured`,
			sensitives: []string{"12345678"},
			expected: `table resolver ignored error: googleapi: Error 403: Cloud Functions API has not been used in project xxxx before or it is disabled. Enable it by visiting https://console.developers.google.com/apis/api/cloudfunctions.googleapis.com/overview?project=xxxx then retry. If you enabled this API recently, wait a few minutes for the action to propagate to our systems and retry.
Details:
[
  {
    "@type": "type.googleapis.com/google.rpc.Help",
    "links": [
      {
        "description": "Google developers console API activation",
        "url": "https://console.developers.google.com/apis/api/cloudfunctions.googleapis.com/overview?project=xxxx"
      }
    ]
  },
  {
    "@type": "type.googleapis.com/google.rpc.ErrorInfo",
    "domain": "googleapis.com",
    "metadata": {
      "consumer": "projects/xxxx",
      "service": "cloudfunctions.googleapis.com"
    },
    "reason": "SERVICE_DISABLED"
  }
]
, accessNotConfigured`,
		},
		{
			input:      `table resolver ignored error: googleapi: Error 403: user@company.iam.gserviceaccount.com does not have storage.buckets.list access to the Google Cloud project., forbidden`,
			sensitives: []string{"user@company"},
			expected:   `table resolver ignored error: googleapi: Error 403: xxxx@xxxx.iam.gserviceaccount.com does not have storage.buckets.list access to the Google Cloud project., forbidden`,
		},
		{
			input:      `compute.disk_types: table resolve error: googleapi: Error 503: Internal error. Please try again or contact Google Support. (Code: '0A0B0CD0000EF.A00BC00.0D0E0FAB'), backendError`,
			sensitives: []string{"0A0B0CD0000EF.A00BC00.0D0E0FAB"},
			expected:   `compute.disk_types: table resolve error: googleapi: Error 503: Internal error. Please try again or contact Google Support. (Code: 'xxxx'), backendError`,
		},
	}
	for _, tc := range tests {
		out := removePII(nil, tc.input)
		for _, s := range tc.sensitives {
			assert.NotContains(t, out, s)
		}
		if tc.expected != "" {
			assert.Equal(t, tc.expected, out)
		}
	}
}
