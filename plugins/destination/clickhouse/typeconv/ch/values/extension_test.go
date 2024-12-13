package values

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func Test_extensionUUID(t *testing.T) {
	const amount = 100

	values := make([]uuid.UUID, amount)
	for i := range values {
		values[i] = uuid.New()
	}

	builder := types.NewUUIDBuilder(array.NewExtensionBuilder(memory.DefaultAllocator, types.NewUUIDType()))
	for _, uid := range values {
		builder.Append(uid)
	}

	data := extensionValue(builder.NewArray().(array.ExtensionArray))
	uidArr := data.([]*uuid.UUID)

	require.Equal(t, amount, len(uidArr))
	for i, uid := range uidArr {
		require.NotNil(t, uid)
		require.Exactly(t, values[i], *uid)
	}
}
