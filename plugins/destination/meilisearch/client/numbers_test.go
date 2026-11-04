package client

import (
	"testing"

	"github.com/goccy/go-json"
	"github.com/stretchr/testify/require"
)

func TestNarrowIntegersRejectOutOfRange(t *testing.T) {
	_, err := toInt8(json.Number("128"))
	require.Error(t, err)
	_, err = toInt16(json.Number("-32769"))
	require.Error(t, err)
	_, err = toInt32(json.Number("2147483648"))
	require.Error(t, err)
	_, err = toUint8(json.Number("256"))
	require.Error(t, err)
	_, err = toUint16(json.Number("65536"))
	require.Error(t, err)
	_, err = toUint32(json.Number("4294967296"))
	require.Error(t, err)
}

func TestNarrowIntegersAcceptInRange(t *testing.T) {
	i8, err := toInt8(json.Number("-128"))
	require.NoError(t, err)
	require.Equal(t, int8(-128), i8)
	u32, err := toUint32(float64(4294967295))
	require.NoError(t, err)
	require.Equal(t, uint32(4294967295), u32)
}
