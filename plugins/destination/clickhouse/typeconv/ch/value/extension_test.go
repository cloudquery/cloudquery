package value

import (
	"testing"

	"github.com/apache/arrow/go/v12/arrow/array"
	"github.com/apache/arrow/go/v12/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v2/types"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func Test_extensionUUID(t *testing.T) {
	const N = 100

	values := make([]uuid.UUID, N)
	for i := range values {
		values[i] = uuid.New()
	}

	bld := types.NewUUIDBuilder(array.NewExtensionBuilder(memory.DefaultAllocator, types.NewUUIDType()))
	for _, uid := range values {
		bld.Append(uid)
	}

	data := extensionValue(bld.NewArray().(array.ExtensionArray))
	uidArr := data.([]*uuid.UUID)

	require.Equal(t, N, len(uidArr))
	for i, uid := range uidArr {
		require.NotNil(t, uid)
		require.Exactly(t, values[i], *uid)
	}
}
