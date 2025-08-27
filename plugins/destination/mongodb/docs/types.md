# MongoDB Types

The MongoDB destination supports most [Apache Arrow](https://arrow.apache.org/docs/index.html)
types. The following table shows the supported types and how they are mapped
to [BSON data types](https://www.mongodb.com/docs/manual/reference/bson-types/).

:::callout{type="info"}
Unsupported types are converted to strings using their string representation.
:::

| Arrow Column Type | Supported? | BSON Type |
|-------------------|------------|-----------|
| Binary            | ✅ Yes      | Binary    |
| Boolean           | ✅ Yes      | Boolean   |
| Date32            | ✅ Yes      | String    |
| Date64            | ✅ Yes      | String    |
| Decimal           | ✅ Yes      | String    |
| Dense Union       | ✅ Yes      | String    |
| Dictionary        | ✅ Yes      | String    |
| Duration          | ✅ Yes      | String    |
| Fixed Size List   | ✅ Yes      | String    |
| Float16           | ✅ Yes      | String    |
| Float32           | ✅ Yes      | Double    |
| Float64           | ✅ Yes      | Double    |
| Inet              | ✅ Yes      | String    |
| Int8              | ✅ Yes      | Int32     |
| Int16             | ✅ Yes      | Int32     |
| Int32             | ✅ Yes      | Int32     |
| Int64             | ✅ Yes      | Int64     |
| Interval[DayTime] | ✅ Yes      | String    |
| Interval[MonthDayNano] | ✅ Yes | String    |
| Interval[Month]   | ✅ Yes      | String    |
| JSON              | ✅ Yes      | Document  |
| Large Binary      | ✅ Yes      | Binary    |
| Large List        | ✅ Yes      | Array     |
| Large String      | ✅ Yes      | String    |
| List              | ✅ Yes      | Array     |
| MAC               | ✅ Yes      | String    |
| Map               | ✅ Yes      | String    |
| String            | ✅ Yes      | String    |
| Struct            | ✅ Yes      | Document  |
| Timestamp         | ✅ Yes      | Date      |
| UUID              | ✅ Yes      | String    |
| Uint8             | ✅ Yes      | Int32     |
| Uint16            | ✅ Yes      | Int32     |
| Uint32            | ✅ Yes      | Int32     |
| Uint64            | ✅ Yes      | Custom † |
| Union             | ✅ Yes      | String    |

## Notes

- Timestamps are converted to MongoDB Date objects using the appropriate time unit
- JSON and Struct types are unmarshaled and stored as native BSON documents
- List types are recursively transformed and stored as BSON arrays
- All integer types smaller than Int32 are promoted to Int32 (BSON's native integer type)
- Float32 values are stored as Double (BSON's native floating-point type)

† Uint64 values are stored as BSON Int64 using signed conversion. Values exceeding 9,223,372,036,854,775,807 (max int64) will overflow and be stored as negative numbers, potentially causing data loss.