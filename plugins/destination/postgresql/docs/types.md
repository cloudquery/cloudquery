# PostgreSQL Types

The PostgreSQL destination supports most [Apache Arrow](https://arrow.apache.org/docs/index.html)
types. The following table shows the supported types and how they are mapped
to [PostgreSQL data types](https://www.postgresql.org/docs/current/datatype.html).

:::callout{type="info"}
Unsupported types are converted to text using their string representation.
:::

| Arrow Column Type      | Supported? | PostgreSQL Type |
|------------------------|------------|-----------------|
| Binary                 | ✅ Yes      | `bytea`         |
| Boolean                | ✅ Yes      | `boolean`       |
| Date32                 | ✅ Yes      | `date`          |
| Date64                 | ✅ Yes      | `date`          |
| Decimal                | ✅ Yes      | `text`          |
| Dense Union            | ✅ Yes      | `text`          |
| Dictionary             | ✅ Yes      | `text`          |
| Duration               | ✅ Yes      | `text`          |
| Fixed Size List        | ✅ Yes      | `text`          |
| Float16                | ✅ Yes      | `text`          |
| Float32                | ✅ Yes      | `real`          |
| Float64                | ✅ Yes      | `double precision` |
| Inet                   | ✅ Yes      | `text`          |
| Int8                   | ✅ Yes      | `smallint`      |
| Int16                  | ✅ Yes      | `smallint`      |
| Int32                  | ✅ Yes      | `integer`       |
| Int64                  | ✅ Yes      | `bigint`        |
| Interval[DayTime]      | ✅ Yes      | `text`          |
| Interval[MonthDayNano] | ✅ Yes      | `text`          |
| Interval[Month]        | ✅ Yes      | `text`          |
| JSON                   | ✅ Yes      | `jsonb`         |
| Large Binary           | ✅ Yes      | `bytea`         |
| Large List             | ✅ Yes      | Array           |
| Large String           | ✅ Yes      | `text`          |
| List                   | ✅ Yes      | Array           |
| MAC                    | ✅ Yes      | `text`          |
| Map                    | ✅ Yes      | `text`          |
| String                 | ✅ Yes      | `text`          |
| Struct                 | ✅ Yes      | `text`          |
| Time32                 | ✅ Yes      | `time`          |
| Time64                 | ✅ Yes      | `time`          |
| Timestamp              | ✅ Yes      | `timestamptz`   |
| UUID                   | ✅ Yes      | `uuid`          |
| Uint8                  | ✅ Yes      | `smallint`      |
| Uint16                 | ✅ Yes      | `integer`       |
| Uint32                 | ✅ Yes      | `bigint`        |
| Uint64                 | ✅ Yes      | `numeric` †     |
| Union                  | ✅ Yes      | `text`          |

## Notes

- Null characters (`\x00`) are automatically stripped from string values for PostgreSQL compatibility
- JSON data is stored as `jsonb` with null characters stripped from string values  
- List types are converted to PostgreSQL arrays with recursive transformation
- Unsigned integer types are promoted to larger signed types to prevent overflow
- Time values are stored with microsecond precision
- Timestamps are stored as `timestamptz` (timestamp with timezone) in UTC

† For CrateDB compatibility, Uint64 values are stored as strings instead of numeric