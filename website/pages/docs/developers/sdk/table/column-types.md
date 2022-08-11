# Column Types

Available types for resource columns are defined in [column.go](https://github.com/cloudquery/cq-provider-sdk/blob/main/provider/schema/column.go).

The SDK types in `schema.ValueType` are tied to both Go and database types:

- Columns defined as a certain ValueType expect certain Go types as values (e.g. A column in the type `schema.TypeBigInt` expects values of the Go types `int`, `uint32` or `int64`)
- `schema.DBTypeFromType` (in [dialect.go](https://github.com/cloudquery/cq-provider-sdk/blob/main/provider/schema/dialect.go)) is used to determine which datatype is used in the database table definition.

| ValueType        | Go Type(s)             |
| ---------------- | ---------------------- |
| TypeBigInt       | int, uint32, int64     |
| TypeBool         | bool                   |
| TypeFloat        | float32, float64       |
| TypeInt          | uint16, int32          |
| TypeSmallInt     | int8, uint8, int16     |
| TypeString       | string                 |
|                  |                        |
| TypeCIDR         | net.IPNet              |
| TypeInet         | net.IP, net.IPAddr     |
| TypeJSON         | struct, slice, map     |
| TypeMacAddr      | net.HardwareAddr       |
| TypeTimestamp    | time.Time              |
| TypeUUID         | uuid.UUID              |
|                  |                        |
| TypeByteArray    | []byte                 |
| TypeStringArray  | []string               |
| TypeCIDRArray    | []net.IPNet            |
| TypeInetArray    | []net.IP, []net.IPAddr |
| TypeIntArray     | []uint16, []int32      |
| TypeMacAddrArray | []net.HardwareAddr     |
| TypeUUIDArray    | []uuid.UUID            |
