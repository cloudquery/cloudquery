package recordupdater

import (
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/google/uuid"
)

func cloneMultipliedColumn(column arrow.Array, j, multiplier int, isPk bool) (arrow.Array, error) {
	switch column.DataType().String() {
	case "utf8":
		values := make([]*string, column.Len()*multiplier)
		for i := 0; i < column.Len(); i++ {
			value := column.GetOneForMarshal(i)
			if value == nil {
				values[i] = nil
				continue
			}
			if isPk {
				values[i] = toPtr(fmt.Sprintf("%s_cloned_%d", value.(string), j))
			} else {
				values[i] = toPtr(value.(string))
			}
		}
		return buildUTF8Column(values), nil
	case "json":
		values := make([]*any, column.Len()*multiplier)
		for i := 0; i < column.Len(); i++ {
			value := column.GetOneForMarshal(i)
			if value == nil {
				values[i] = nil
				continue
			}
			rawJSON := value.(json.RawMessage)
			if isPk {
				values[i] = toPtr(any(fmt.Sprintf("%s_cloned_%d", string(rawJSON), j)))
			} else {
				values[i] = toPtr(any(string(rawJSON)))
			}
		}
		return buildJSONColumn(values), nil
	case "inet":
		values := make([]*any, column.Len()*multiplier)
		for i := 0; i < column.Len(); i++ {
			value := column.GetOneForMarshal(i)
			if value == nil {
				values[i] = nil
				continue
			}
			if isPk {
				values[i] = toPtr(any(fmt.Sprintf("%s_cloned_%d", value.(string), j)))
			} else {
				values[i] = toPtr(any(value.(string)))
			}
		}
		return buildInetColumn(values), nil
	case "uuid":
		values := make([]uuid.UUID, column.Len()*multiplier)
		for i := 0; i < column.Len(); i++ {
			randomUUID, _ := uuid.NewRandom()
			if isPk {
				values[i] = randomUUID
			} else {
				values[i] = randomUUID
			}
		}
		return buildUUIDColumn(values), nil
	case "int64":
		values := make([]*int64, column.Len()*multiplier)
		constant := int64(12345678)
		for i := 0; i < column.Len(); i++ {
			value := column.GetOneForMarshal(i)
			if value == nil {
				values[i] = nil
				continue
			}
			if isPk {
				values[i] = toPtr(value.(int64) + constant + int64(i))
			} else {
				values[i] = toPtr(value.(int64))
			}
		}
		return buildInt64Column(values), nil
	case "int32":
		values := make([]*int32, column.Len()*multiplier)
		constant := int32(12345678)
		for i := 0; i < column.Len(); i++ {
			value := column.GetOneForMarshal(i)
			if value == nil {
				values[i] = nil
				continue
			}
			if isPk {
				values[i] = toPtr(value.(int32) + constant + int32(i))
			} else {
				values[i] = toPtr(value.(int32))
			}
		}
		return buildInt32Column(values), nil
	case "float64":
		values := make([]*float64, column.Len()*multiplier)
		constant := float64(12345678)
		for i := 0; i < column.Len(); i++ {
			value := column.GetOneForMarshal(i)
			if value == nil {
				values[i] = nil
				continue
			}
			if isPk {
				values[i] = toPtr(value.(float64) + constant + float64(i))
			} else {
				values[i] = toPtr(value.(float64))
			}
		}
		return buildFloat64Column(values), nil
	case "bool":
		values := make([]*bool, column.Len()*multiplier)
		for i := 0; i < column.Len(); i++ {
			if val := column.GetOneForMarshal(i); val != nil {
				if isPk {
					values[i] = toPtr(!val.(bool))
				} else {
					values[i] = toPtr(val.(bool))
				}
			} else {
				values[i] = nil
			}
		}
		return buildBoolColumn(values), nil
	case "timestamp[us, tz=UTC]":
		values := make([]*time.Time, column.Len()*multiplier)
		constantDuration := 100 * 24 * 365 * time.Hour
		for i := 0; i < column.Len(); i++ {
			val := column.GetOneForMarshal(i)
			if val == nil {
				values[i] = nil
				continue
			}
			timeVal, err := time.Parse("2006-01-02 15:04:05.999999Z", val.(string))
			if err != nil {
				return nil, err
			}
			if isPk {
				values[i] = toPtr(timeVal.Add(constantDuration + time.Duration(i)*time.Second))
			} else {
				values[i] = toPtr(timeVal)
			}
		}
		return buildTimestampColumn(values)
	case "list<item: uuid, nullable>":
		values := make([][]uuid.UUID, column.Len()*multiplier)
		for i := 0; i < column.Len(); i++ {
			randomUUID, _ := uuid.NewRandom()
			if isPk {
				values[i] = []uuid.UUID{randomUUID}
			} else {
				values[i] = []uuid.UUID{randomUUID}
			}
		}
		return buildUUIDListColumn(values)
	case "list<item: utf8, nullable>":
		values := make([][]string, column.Len()*multiplier)
		for i := 0; i < column.Len(); i++ {
			if val := column.GetOneForMarshal(i); val != nil {
				if isPk {
					values[i] = []string{fmt.Sprintf("%s_cloned_%d", string(val.(json.RawMessage)), j)}
				} else {
					values[i] = []string{string(val.(json.RawMessage))}
				}
			} else {
				values[i] = []string{}
			}
		}
		return buildUTF8ListColumn(values)
	case "list<item: int64, nullable>":
		values := make([][]int64, column.Len()*multiplier)
		for i := 0; i < column.Len(); i++ {
			if isPk {
				values[i] = []int64{column.GetOneForMarshal(i).(int64) + int64(i)}
			} else {
				values[i] = []int64{column.GetOneForMarshal(i).(int64)}
			}
		}
		return buildInt64ListColumn(values)
	case "list<item: json, nullable>":
		values := make([][]json.RawMessage, column.Len()*multiplier)
		for i := 0; i < column.Len(); i++ {
			if val := column.GetOneForMarshal(i); val != nil {
				if isPk {
					values[i] = []json.RawMessage{val.(json.RawMessage)}
				} else {
					values[i] = []json.RawMessage{val.(json.RawMessage)}
				}
			} else {
				values[i] = []json.RawMessage{}
			}
		}
		return buildJSONListColumn(values), nil
	case "list<item: inet, nullable>":
		values := make([][]net.IPNet, column.Len()*multiplier)
		for i := 0; i < column.Len(); i++ {
			if isPk {
				values[i] = []net.IPNet{{}}
			} else {
				values[i] = []net.IPNet{{}}
			}
		}
		return buildInetListColumn(values), nil
	case "binary":
		values := make([][]byte, column.Len()*multiplier)
		for i := 0; i < column.Len(); i++ {
			if isPk {
				values[i] = []byte(fmt.Sprintf("%s_cloned_%d", string(column.GetOneForMarshal(i).([]uint8)), j))
			} else {
				values[i] = []byte(string(column.GetOneForMarshal(i).([]uint8)))
			}
		}
		return buildBinaryColumn(values)
	default:
		return nil, fmt.Errorf("unsupported data type: %s, %s", column.DataType().ID(), column.DataType().String())
	}
}

func toPtr[T any](value T) *T {
	return &value
}
