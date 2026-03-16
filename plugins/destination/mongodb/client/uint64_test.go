package client

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func marshalWithRegistry(reg *bson.Registry, val any) ([]byte, error) {
	var buf bytes.Buffer
	vw := bson.NewDocumentWriter(&buf)
	enc := bson.NewEncoder(vw)
	enc.SetRegistry(reg)
	if err := enc.Encode(val); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func TestGetRegistry_CustomUint64Encoding(t *testing.T) {
	reg := getRegistry()
	require.NotNil(t, reg)

	val := CustomUnit64(12345)
	doc := bson.M{"value": val}

	encoded, err := marshalWithRegistry(reg, doc)
	require.NoError(t, err)

	var result bson.M
	err = bson.Unmarshal(encoded, &result)
	require.NoError(t, err)

	require.Equal(t, int64(12345), result["value"])
}

func TestGetRegistry_LargeUint64(t *testing.T) {
	reg := getRegistry()

	// Value that exceeds int64 max when unsigned
	val := CustomUnit64(1<<63 + 100)
	doc := bson.M{"value": val}

	encoded, err := marshalWithRegistry(reg, doc)
	require.NoError(t, err)

	var result bson.M
	err = bson.Unmarshal(encoded, &result)
	require.NoError(t, err)

	// Stored as int64 (negative), but roundtrips back to uint64
	got := uint64(result["value"].(int64))
	require.Equal(t, uint64(val), got)
}

func TestGetRegistry_RegularTypes(t *testing.T) {
	reg := getRegistry()

	doc := bson.M{
		"str":   "hello",
		"int":   int32(42),
		"float": 3.14,
		"bool":  true,
	}

	encoded, err := marshalWithRegistry(reg, doc)
	require.NoError(t, err)

	var result bson.M
	err = bson.Unmarshal(encoded, &result)
	require.NoError(t, err)

	require.Equal(t, "hello", result["str"])
	require.Equal(t, int32(42), result["int"])
	require.InDelta(t, 3.14, result["float"], 0.001)
	require.Equal(t, true, result["bool"])
}
