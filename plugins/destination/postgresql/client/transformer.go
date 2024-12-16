package client

import (
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/jackc/pgx/v5/pgtype"
)

func stripNulls(s string) string {
	return strings.ReplaceAll(s, "\x00", "")
}

// Best effort to strip nulls from json. In case of an error, we fallback to returning the input.
// We only strip the nulls from the values (i.e. never from the keys).
// In some far-fetched case a JSON like this may happen, but we ignore this for now: `{"key\u0000": "v1", "key\u0000\u0000": "v2"}`
func stripNullsFromMarshalledJson(inputJsonMarshaled []byte) []byte {
	var m any
	err := json.Unmarshal(inputJsonMarshaled, &m)
	if err != nil {
		return inputJsonMarshaled
	}

	m = stripNullsFromJsonValue(m)
	outputJsonMarshaled, err := json.Marshal(m)
	if err != nil {
		return inputJsonMarshaled
	}

	return outputJsonMarshaled
}

func stripNullsFromJsonValue(val any) any {
	switch concreteVal := val.(type) {
	case string:
		return stripNulls(concreteVal)
	case map[string]any:
		for k := range concreteVal {
			concreteVal[k] = stripNullsFromJsonValue(concreteVal[k])
		}
		return concreteVal
	case []any:
		for i := range concreteVal {
			concreteVal[i] = stripNullsFromJsonValue(concreteVal[i])
		}
		return concreteVal
	default:
		return val
	}
}

func (c *Client) transformArr(arr arrow.Array) []any {
	pgArr := make([]any, arr.Len())
	for i := 0; i < arr.Len(); i++ {
		if arr.IsNull(i) || !arr.IsValid(i) {
			pgArr[i] = nil
			continue
		}
		switch a := arr.(type) {
		case *array.Boolean:
			pgArr[i] = pgtype.Bool{
				Bool:  a.Value(i),
				Valid: a.IsValid(i),
			}
		case *array.Int8:
			pgArr[i] = pgtype.Int2{
				Int16: int16(a.Value(i)),
				Valid: a.IsValid(i),
			}
		case *array.Int16:
			pgArr[i] = pgtype.Int2{
				Int16: a.Value(i),
				Valid: a.IsValid(i),
			}
		case *array.Int32:
			pgArr[i] = pgtype.Int4{
				Int32: a.Value(i),
				Valid: a.IsValid(i),
			}
		case *array.Int64:
			pgArr[i] = pgtype.Int8{
				Int64: a.Value(i),
				Valid: a.IsValid(i),
			}
		case *array.Uint8:
			pgArr[i] = pgtype.Int2{
				Int16: int16(a.Value(i)),
				Valid: a.IsValid(i),
			}
		case *array.Uint16:
			pgArr[i] = pgtype.Int4{
				Int32: int32(a.Value(i)),
				Valid: a.IsValid(i),
			}
		case *array.Uint32:
			pgArr[i] = pgtype.Int8{
				Int64: int64(a.Value(i)),
				Valid: a.IsValid(i),
			}
		case *array.Uint64:
			if c.pgType == pgTypeCrateDB {
				pgArr[i] = a.ValueStr(i)
			} else {
				pgArr[i] = a.Value(i)
			}
		case *array.Float32:
			pgArr[i] = pgtype.Float4{
				Float32: a.Value(i),
				Valid:   a.IsValid(i),
			}
		case *array.Float64:
			pgArr[i] = pgtype.Float8{
				Float64: a.Value(i),
				Valid:   a.IsValid(i),
			}
		case *array.Binary:
			if c.pgType == pgTypeCrateDB {
				pgArr[i] = base64.StdEncoding.EncodeToString(a.Value(i))
			} else {
				pgArr[i] = a.Value(i)
			}
		case *array.LargeBinary:
			if c.pgType == pgTypeCrateDB {
				pgArr[i] = base64.StdEncoding.EncodeToString(a.Value(i))
			} else {
				pgArr[i] = a.Value(i)
			}
		case *array.String:
			pgArr[i] = pgtype.Text{
				String: stripNulls(a.Value(i)),
				Valid:  a.IsValid(i),
			}
		case *array.LargeString:
			pgArr[i] = pgtype.Text{
				String: stripNulls(a.Value(i)),
				Valid:  a.IsValid(i),
			}
		case *array.Timestamp:
			pgArr[i] = pgtype.Timestamptz{
				Time:  a.Value(i).ToTime(a.DataType().(*arrow.TimestampType).Unit).UTC(),
				Valid: a.IsValid(i),
			}
		case *array.Time32:
			pgArr[i] = pgtype.Time{
				Microseconds: a.Value(i).ToTime(a.DataType().(*arrow.Time32Type).Unit).UTC().UnixMicro(),
				Valid:        a.IsValid(i),
			}
		case *array.Time64:
			pgArr[i] = pgtype.Time{
				Microseconds: a.Value(i).ToTime(a.DataType().(*arrow.Time64Type).Unit).UTC().UnixMicro(),
				Valid:        a.IsValid(i),
			}
		case *array.Date32:
			pgArr[i] = pgtype.Date{
				Time:  a.Value(i).ToTime().UTC(),
				Valid: a.IsValid(i),
			}
		case *array.Date64:
			pgArr[i] = pgtype.Date{
				Time:  a.Value(i).ToTime().UTC(),
				Valid: a.IsValid(i),
			}
		case *types.UUIDArray:
			bUUID, err := a.Value(i).MarshalBinary()
			if err != nil {
				panic(err)
			}
			pgUUID := pgtype.UUID{
				Valid: a.IsValid(i),
			}
			copy(pgUUID.Bytes[:], bUUID)
			pgArr[i] = pgUUID
		case *array.Map:
			pgArr[i] = stripNulls(arr.ValueStr(i))
		case array.ListLike:
			start, end := a.ValueOffsets(i)
			nested := array.NewSlice(a.ListValues(), start, end)
			pgArr[i] = c.transformArr(nested)
		case *types.JSONArray:
			pgArr[i] = stripNullsFromMarshalledJson(a.Storage().(*array.Binary).Value(i))
		default:
			pgArr[i] = stripNulls(arr.ValueStr(i))
		}
	}

	return pgArr
}

func (c *Client) transformValues(r arrow.Record) [][]any {
	results := make([][]any, r.NumRows())

	for i := range results {
		results[i] = make([]any, r.NumCols())
	}

	for i := 0; i < int(r.NumCols()); i++ {
		col := r.Column(i)
		transformed := c.transformArr(col)
		for l := 0; l < col.Len(); l++ {
			results[l][i] = transformed[l]
		}
	}
	return results
}
