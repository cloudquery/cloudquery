# DuckDB Types

The DuckDB destination supports most [Apache Arrow](https://arrow.apache.org/docs/index.html)
types. The following table shows the supported types and how they are mapped
to [DuckDB data types](https://duckdb.org/docs/sql/data_types/overview).

:::callout{type="info"}
Unsupported types are always mapped to `varchar`.
:::

| Arrow Column Type | Supported? | DuckDB Type |
|-------------------|------------|-------------|
| Binary            | ✅ Yes      | `blob`      |
| Boolean           | ✅ Yes      | `boolean`   |
| Date32            | ✅ Yes      | `timestamp` |
| Date64            | ✅ Yes      | `timestamp` |
| Decimal           | ✅ Yes      | `varchar`   |
| Dense Union       | ✅ Yes      | `varchar`   |
| Dictionary        | ✅ Yes      | `varchar`   |
| Duration          | ✅ Yes      | `varchar`   |
| Fixed Size List   | ✅ Yes      | `varchar`   |
| Float16           | ✅ Yes      | `varchar`   |
| Float32           | ✅ Yes      | `float`     |
| Float64           | ✅ Yes      | `double`    |
| Inet              | ✅ Yes      | `varchar`   |
| Int8              | ✅ Yes      | `tinyint`   |
| Int16             | ✅ Yes      | `smallint`  |
| Int32             | ✅ Yes      | `integer`   |
| Int64             | ✅ Yes      | `bigint`    |
| Interval[DayTime] | ✅ Yes      | `varchar`   |
| Interval[MonthDayNano] | ✅ Yes | `varchar`   |
| Interval[Month]   | ✅ Yes      | `varchar`   |
| JSON              | ✅ Yes      | `json`      |
| Large Binary      | ✅ Yes      | `blob`      |
| Large List        | ✅ Yes      | Array of element type |
| Large String      | ✅ Yes      | `varchar`   |
| List              | ✅ Yes      | Array of element type |
| MAC               | ✅ Yes      | `varchar`   |
| Map               | ✅ Yes      | `varchar` † |
| String            | ✅ Yes      | `varchar`   |
| Struct            | ✅ Yes      | `varchar` † |
| Timestamp         | ✅ Yes      | `timestamp` |
| UUID              | ✅ Yes      | `uuid`      |
| Uint8             | ✅ Yes      | `uinteger`  |
| Uint16            | ✅ Yes      | `uinteger`  |
| Uint32            | ✅ Yes      | `uinteger`  |
| Uint64            | ✅ Yes      | `ubigint`   |
| Union             | ✅ Yes      | `varchar`   |

## Notes

- DuckDB supports native unsigned integer types (`uinteger`, `ubigint`)
- List types are converted to DuckDB arrays with the appropriate element type (e.g., `integer[]`)
- Date32 and Date64 types are both mapped to `timestamp` for compatibility
- UUID and JSON types have native DuckDB support

† Complex types like Struct and Map are converted to `varchar` for storage, losing their structured nature