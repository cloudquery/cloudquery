package module

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersionErrors(t *testing.T) {
	const (
		newerErr      = "provider P seems to support a newer version of M, which is incompatible with your cloudquery version"
		olderErr      = "provider P seems to support an older version of M, which is incompatible with your cloudquery version"
		noSupportErr  = "provider P doesn't support M yet"
		noSupportErr2 = "providers P1, P2 don't support M yet"
	)

	for _, tc := range []struct {
		Mod           []uint32
		Prov          map[string][]uint32
		ExpectedError string
	}{
		{
			Mod: []uint32{1, 2},
			Prov: map[string][]uint32{
				"P": {3},
			},
			ExpectedError: newerErr,
		},
		{
			Mod: []uint32{2},
			Prov: map[string][]uint32{
				"P": {1},
			},
			ExpectedError: olderErr,
		},
		{
			Mod: []uint32{2},
			Prov: map[string][]uint32{
				"P": nil,
			},
			ExpectedError: noSupportErr,
		},
		{
			Mod: []uint32{1},
			Prov: map[string][]uint32{
				"P1": nil,
				"P2": nil,
			},
			ExpectedError: noSupportErr2,
		},
	} {
		err := versionError("M", tc.Mod, tc.Prov)
		assert.Error(t, err)
		if tc.ExpectedError != "" {
			assert.EqualError(t, err, tc.ExpectedError)
		}
	}
}
