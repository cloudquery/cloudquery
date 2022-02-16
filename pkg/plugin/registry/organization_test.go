package registry

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseProviderName(t *testing.T) {

	testCases := []struct {
		Name         string
		ProviderName string
		ExpectedOrg  string
		ExpectedName string
		ExpectedErr  error
	}{
		{
			Name:         "simple",
			ProviderName: "aws",
			ExpectedOrg:  "cloudquery",
			ExpectedName: "aws",
			ExpectedErr:  nil,
		},
		{
			Name:         "simple caps",
			ProviderName: "AWS",
			ExpectedOrg:  "cloudquery",
			ExpectedName: "aws",
			ExpectedErr:  nil,
		},
		{
			Name:         "with org caps",
			ProviderName: "CLOudQuery/AWS",
			ExpectedOrg:  "cloudquery",
			ExpectedName: "aws",
			ExpectedErr:  nil,
		},
		{
			Name:         "with org name",
			ProviderName: "cloudquery/aws",
			ExpectedOrg:  "cloudquery",
			ExpectedName: "aws",
			ExpectedErr:  nil,
		},
		{
			Name:         "other org name",
			ProviderName: "otherorg/aws",
			ExpectedOrg:  "otherorg",
			ExpectedName: "aws",
			ExpectedErr:  nil,
		},
		{
			Name:         "invalid name",
			ProviderName: "otherorg/invalid/aws",
			ExpectedOrg:  "",
			ExpectedName: "",
			ExpectedErr:  errors.New("invalid provider name otherorg/invalid/aws"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			org, name, err := ParseProviderName(tc.ProviderName)
			assert.Equal(t, tc.ExpectedOrg, org)
			assert.Equal(t, tc.ExpectedName, name)
			assert.Equal(t, tc.ExpectedErr, err)

		})

	}

}
