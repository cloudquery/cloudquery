package client

import (
	"reflect"
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/google/uuid"
	mssql "github.com/microsoft/go-mssqldb"
	"github.com/stretchr/testify/require"
)

func TestTVPValues(t *testing.T) {
	columns := schema.ColumnList{
		schema.CqIDColumn,
		schema.CqParentIDColumn,
		schema.CqSourceNameColumn,
		schema.CqSyncTimeColumn,
		schema.Column{Name: "tst", Type: schema.TypeFloat},
	}

	tf := tableTransformer(columns)
	uid1, _ := mssql.UniqueIdentifier(uuid.New()).Value()
	uid2, _ := mssql.UniqueIdentifier(uuid.New()).Value()
	row := tf([][]any{{uid1, uid2, "source_name_val", time.Now(), 3.5}})

	checkTVP(t, row)
}

func checkTVP(t *testing.T, val any) {
	valueOf := reflect.ValueOf(val)
	require.Equal(t, reflect.Slice, valueOf.Kind())
	require.False(t, valueOf.IsNil())
	typeOf := reflect.TypeOf(val)
	require.Equal(t, reflect.Struct, typeOf.Elem().Kind())

	require.NotZero(t, valueOf.Len())
	for i := 0; i < valueOf.Len(); i++ {
		el := valueOf.Index(i)
		switch el.Type().Kind() {
		case reflect.Pointer:
			require.False(t, el.IsNil())
		default:
			require.False(t, el.IsZero())
		}
	}
}
