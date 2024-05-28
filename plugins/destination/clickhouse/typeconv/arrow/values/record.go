package values

import (
	"github.com/apache/arrow/go/v16/arrow/array"
)

func AppendToRecordBuilder(builder *array.RecordBuilder, data []any) error {
	for i, builder := range builder.Fields() {
		if err := buildValue(builder, data[i]); err != nil {
			return err
		}
	}
	return nil
}
