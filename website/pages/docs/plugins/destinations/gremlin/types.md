# Gremlin Types

The Gremlin destination (`v2.0.0` and later) supports most [Apache Arrow](https://arrow.apache.org/docs/index.html) types. The following table shows the supported types and how they are mapped to Gremlin data types.

| Arrow Column Type      | Supported? | Gremlin Type          |
|------------------------|------------|-----------------------|
| Binary                 | ✅ Yes      | `Bytes`               |
| Boolean                | ✅ Yes      | `Boolean`             |
| Date32                 | ✅ Yes      | `String`              |
| Date64                 | ✅ Yes      | `String`              |
| Decimal                | ✅ Yes      | `String`              |
| Dense Union            | ✅ Yes      | `String`              |
| Dictionary             | ✅ Yes      | `String`              |
| Duration[ms]           | ✅ Yes      | `String`              |
| Duration[ns]           | ✅ Yes      | `String`              |
| Duration[s]            | ✅ Yes      | `String`              |
| Duration[us]           | ✅ Yes      | `String`              |
| Fixed Size List        | ✅ Yes      | `String`              |
| Float16                | ✅ Yes      | `String`              |
| Float32                | ✅ Yes      | `Float`               |
| Float64                | ✅ Yes      | `Float`               |
| Inet                   | ✅ Yes      | `String`              |
| Int8                   | ✅ Yes      | `Integer`             |
| Int16                  | ✅ Yes      | `Integer`             |
| Int32                  | ✅ Yes      | `Integer`             |
| Int64                  | ✅ Yes      | `Integer`             |
| Interval[DayTime]      | ✅ Yes      | `String`              |
| Interval[MonthDayNano] | ✅ Yes      | `String`              |
| Interval[Month]        | ✅ Yes      | `String`              |
| JSON                   | ✅ Yes      | `String`              |
| Large Binary           | ✅ Yes      | `Bytes`               |
| Large List             | ✅ Yes      | `String`              |
| Large String           | ✅ Yes      | `String`              |
| List                   | ✅ Yes      | `String` or `List` †  |
| MAC                    | ✅ Yes      | `String`              |
| Map                    | ✅ Yes      | `String`              |
| String                 | ✅ Yes      | `String`              |
| Struct                 | ✅ Yes      | `String`              |
| Timestamp[ms]          | ✅ Yes      | `String` <sup>*</sup> |
| Timestamp[ns]          | ✅ Yes      | `String`              |
| Timestamp[s]           | ✅ Yes      | `String`              |
| Timestamp[us]          | ✅ Yes      | `String`              |
| UUID                   | ✅ Yes      | `String`              |
| Uint8                  | ✅ Yes      | `String`              |
| Uint16                 | ✅ Yes      | `Integer`             |
| Uint32                 | ✅ Yes      | `Integer`             |
| Uint64                 | ✅ Yes      | `Integer`             |
| Union                  | ✅ Yes      | `String`              |

## Notes

<sup>*</sup> Timestamps are converted to strings in the format `yyyy-MM-dd HH:mm:ss.SSSSSSSSS` (UTC timezone) (e.g. `2021-01-01 00:00:00.000000000`). `_cq_sync_time` column is persisted in native `Timestamp` type.

† List types are persisted as-is only if `complete_types` option is enabled. Otherwise, they are converted to strings.

`NUL` bytes are stripped from strings.
