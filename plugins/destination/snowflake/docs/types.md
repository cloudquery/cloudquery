The Snowflake destination supports most [Apache Arrow](https://arrow.apache.org/docs/index.html)
types. The following table shows the supported types and how they are mapped
to [Snowflake data types](https://docs.snowflake.com/en/sql-reference/data-types).

:::callout{type="info"}
Unsupported types are always mapped to `text`.
:::

| Arrow Column Type      | Supported? | Snowflake Type |
|------------------------|------------|----------------|
| Binary                 | ✅ Yes      | `binary`       |
| Boolean                | ✅ Yes      | `boolean`      |
| Date32                 | ✅ Yes      | `text`         |
| Date64                 | ✅ Yes      | `text`         |
| Decimal                | ✅ Yes      | `text`         |
| Dense Union            | ✅ Yes      | `text`         |
| Dictionary             | ✅ Yes      | `text`         |
| Duration               | ✅ Yes      | `text`         |
| Fixed Size List        | ✅ Yes      | `array`        |
| Float16                | ✅ Yes      | `text`         |
| Float32                | ✅ Yes      | `float`        |
| Float64                | ✅ Yes      | `float`        |
| Inet                   | ✅ Yes      | `text`         |
| Int8                   | ✅ Yes      | `number`       |
| Int16                  | ✅ Yes      | `number`       |
| Int32                  | ✅ Yes      | `number`       |
| Int64                  | ✅ Yes      | `number`       |
| Interval[DayTime]      | ✅ Yes      | `text`         |
| Interval[MonthDayNano] | ✅ Yes      | `text`         |
| Interval[Month]        | ✅ Yes      | `text`         |
| JSON                   | ✅ Yes      | `variant`      |
| Large Binary           | ✅ Yes      | `binary`       |
| Large List             | ✅ Yes      | `array`        |
| Large String           | ✅ Yes      | `text`         |
| List                   | ✅ Yes      | `array`        |
| MAC                    | ✅ Yes      | `text`         |
| Map                    | ✅ Yes      | `text`         |
| String                 | ✅ Yes      | `text`         |
| Struct                 | ✅ Yes      | `variant`      |
| Time32                 | ✅ Yes      | `text`         |
| Time64                 | ✅ Yes      | `text`         |
| Timestamp              | ✅ Yes      | `timestamp_tz` |
| UUID                   | ✅ Yes      | `text`         |
| Uint8                  | ✅ Yes      | `number`       |
| Uint16                 | ✅ Yes      | `number`       |
| Uint32                 | ✅ Yes      | `number`       |
| Uint64                 | ✅ Yes      | `number`       |
| Union                  | ✅ Yes      | `text`         |

## Notes

- All integer types (signed and unsigned) are mapped to Snowflake's `number` type, which can handle arbitrary precision
- Both Float32 and Float64 are stored as Snowflake's `float` type
- List types are stored as Snowflake `array` type, including both regular and fixed-size lists
- JSON and Struct types use Snowflake's `variant` type for semi-structured data storage
- Timestamps are stored as `timestamp_tz` (timestamp with timezone)
- Complex types not natively supported by Snowflake fall back to `text` representation