The MySQL destination (MySQL 8.0 and later) supports most [Apache Arrow](https://arrow.apache.org/docs/index.html)
types. The following table shows the supported types and how they are mapped
to [MySQL data types](https://dev.mysql.com/doc/refman/8.0/en/data-types.html).

:::callout{type="info"}
Unsupported types are always mapped to `text`.
:::

| Arrow Column Type | Supported? | MySQL Type        |
|-------------------|------------|-------------------|
| Binary            | ✅ Yes      | `blob`            |
| Boolean           | ✅ Yes      | `tinyint(1)`      |
| Date32            | ✅ Yes      | `text`            |
| Date64            | ✅ Yes      | `text`            |
| Decimal           | ✅ Yes      | `text`            |
| Dense Union       | ✅ Yes      | `text`            |
| Dictionary        | ✅ Yes      | `text`            |
| Duration          | ✅ Yes      | `text`            |
| Fixed Size List   | ✅ Yes      | `text`            |
| Float16           | ✅ Yes      | `text`            |
| Float32           | ✅ Yes      | `float`           |
| Float64           | ✅ Yes      | `double`          |
| Inet              | ✅ Yes      | `text`            |
| Int8              | ✅ Yes      | `tinyint`         |
| Int16             | ✅ Yes      | `smallint`        |
| Int32             | ✅ Yes      | `int`             |
| Int64             | ✅ Yes      | `bigint`          |
| Interval[DayTime] | ✅ Yes      | `text`            |
| Interval[MonthDayNano] | ✅ Yes | `text`            |
| Interval[Month]   | ✅ Yes      | `text`            |
| JSON              | ✅ Yes      | `json`            |
| Large Binary      | ✅ Yes      | `blob`            |
| Large List        | ✅ Yes      | `text`            |
| Large String      | ✅ Yes      | `text`            |
| List              | ✅ Yes      | `json`            |
| MAC               | ✅ Yes      | `text`            |
| Map               | ✅ Yes      | `text`            |
| String            | ✅ Yes      | `text`            |
| Struct            | ✅ Yes      | `json`            |
| Timestamp         | ✅ Yes      | `datetime(6)`     |
| UUID              | ✅ Yes      | `binary(16)`      |
| Uint8             | ✅ Yes      | `tinyint unsigned`|
| Uint16            | ✅ Yes      | `smallint unsigned`|
| Uint32            | ✅ Yes      | `int unsigned`    |
| Uint64            | ✅ Yes      | `bigint unsigned` |
| Union             | ✅ Yes      | `text`            |

## Notes

- Boolean values are stored as `tinyint(1)` to align with MySQL's information schema representation
- Timestamps include microsecond precision with `datetime(6)`
- Complex data types like Struct and List are stored as JSON when supported natively
- All other unsupported types fall back to `text` for maximum compatibility
