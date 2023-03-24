package gaql

import (
	"testing"

	"github.com/shenzhencenter/google-ads-pb/resources"
	"github.com/stretchr/testify/require"
)

func TestFieldName(t *testing.T) {
	type testCase struct {
		name string
		typ  any
	}
	for _, tc := range []testCase{
		{
			name: "ad_group",
			typ:  new(resources.AdGroup),
		},
		{
			name: "ad_group_ad",
			typ:  new(resources.AdGroupAd),
		},
		{
			name: "customer",
			typ:  new(resources.Customer),
		},
	} {
		require.Equal(t, tc.name, FieldName(tc.typ))
	}
}
