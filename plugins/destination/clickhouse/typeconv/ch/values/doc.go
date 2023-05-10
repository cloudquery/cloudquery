// Package values allows conversion from Apache Arrow array to the type accepted by ClickHouse SDK
// The conversion can become tricky for some nested types, specifically, maps.
// Nullable(item) -> *scanType(item). However, non-nullable are also saved as pointers for ease of coding.
// Decimal(precision,scale) -> decimal.Decimal
// UUID -> uuid.UUID
// Tuple(??) -> map[string]any (arrow supports only tuples with fields)
// Map(keys, values) -> map[scanType(keys)]scanType(values)
// Array(items) -> []scanType(items)
//
// example:
// Map(String, Map(UUID, Map(String, Tuple(`uint8` Nullable(UInt8), `uuid` UUID)))) -> *map[string]*map[uuid.UUID]*map[string]*map[string]any
package values
