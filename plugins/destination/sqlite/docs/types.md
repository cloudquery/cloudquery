---
hub-title: Supported Types
hub-order: 5
---

The SQLite destination supports most [Apache Arrow](https://arrow.apache.org/docs/index.html)
types. The following table shows the supported types and how they are mapped
to [SQLite data types](https://www.sqlite.org/datatype3.html).

:::callout{type="info"}
Unsupported types are always mapped to `text`.
:::

| Arrow Column Type      | Supported? | SQLite Type |
| ---------------------- | ---------- | ----------- |
| Binary                 | ✅ Yes      | `blob`      |
| Boolean                | ✅ Yes     | `boolean`   |
| Date32                 | ✅ Yes     | `text`      |
| Date64                 | ✅ Yes     | `text`      |
| Decimal                | ✅ Yes     | `text`      |
| Dense Union            | ✅ Yes     | `text`      |
| Dictionary             | ✅ Yes     | `text`      |
| Duration               | ✅ Yes     | `text`      |
| Fixed Size List        | ✅ Yes     | `text`      |
| Float16                | ✅ Yes     | `real`      |
| Float32                | ✅ Yes     | `real`      |
| Float64                | ✅ Yes     | `real`      |
| Inet                   | ✅ Yes     | `text`      |
| Int8                   | ✅ Yes     | `integer`   |
| Int16                  | ✅ Yes     | `integer`   |
| Int32                  | ✅ Yes     | `integer`   |
| Int64                  | ✅ Yes     | `integer`   |
| Interval[DayTime]      | ✅ Yes     | `text`      |
| Interval[MonthDayNano] | ✅ Yes     | `text`      |
| Interval[Month]        | ✅ Yes     | `text`      |
| JSON                   | ✅ Yes     | `text`      |
| Large Binary           | ✅ Yes     | `blob`      |
| Large List             | ✅ Yes     | `text`      |
| Large String           | ✅ Yes     | `text`      |
| List                   | ✅ Yes     | `text`      |
| MAC                    | ✅ Yes     | `text`      |
| Map                    | ✅ Yes     | `text`      |
| String                 | ✅ Yes     | `text`      |
| Struct                 | ✅ Yes     | `text`      |
| Time32                 | ✅ Yes     | `text`      |
| Time64                 | ✅ Yes     | `text`      |
| Timestamp              | ✅ Yes     | `timestamp` |
| UUID                   | ✅ Yes     | `text`      |
| Uint8                  | ✅ Yes     | `integer`   |
| Uint16                 | ✅ Yes     | `integer`   |
| Uint32                 | ✅ Yes     | `integer`   |
| Uint64                 | ✅ Yes     | `integer`   |
| Union                  | ✅ Yes     | `text`      |

## Notes

- SQLite has a simplified type system with only 5 storage classes: NULL, INTEGER, REAL, TEXT, and BLOB
- All integer types (signed and unsigned, 8-bit to 64-bit) are stored as SQLite `integer`
- All floating-point types (Float16, Float32, Float64) are stored as SQLite `real`
- Complex data types like JSON, List, and Struct are serialized and stored as `text`
- Binary data uses SQLite's `blob` storage class
- SQLite's dynamic typing system allows flexible data storage regardless of declared column type
