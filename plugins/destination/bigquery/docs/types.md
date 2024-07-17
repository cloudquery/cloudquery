# BigQuery Types

The BigQuery destination (`v3.0.0` and later) supports most [Apache Arrow](https://arrow.apache.org/docs/index.html)
types. The following table shows the supported types and how they are mapped
to [BigQuery data types](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types).

| Arrow Column Type      | Supported? | BigQuery Type                                            |
|------------------------|------------|----------------------------------------------------------|
| Binary                 | ✅ Yes      | `BYTES`                                                  |
| Boolean                | ✅ Yes      | `BOOL`                                                   |
| Date32                 | ✅ Yes      | `DATE`                                                   |
| Date64                 | ✅ Yes      | `DATE`                                                   |
| Decimal                | ✅ Yes      | `BIGNUMERIC`                                             |
| Dense Union            | ❌ No       |                                                          |
| Dictionary             | ❌ No       |                                                          |
| Duration               | ✅ Yes      | `INT64`                                                  |
| Fixed Size List        | ✅ Yes      | (Repeated column) †                                      |
| Float16                | ✅ Yes      | `FLOAT64`                                                |
| Float32                | ✅ Yes      | `FLOAT64`                                                |
| Float64                | ✅ Yes      | `FLOAT64`                                                |
| Inet                   | ✅ Yes      | `STRING`                                                 |
| Int8                   | ✅ Yes      | `INT64`                                                  |
| Int16                  | ✅ Yes      | `INT64`                                                  |
| Int32                  | ✅ Yes      | `INT64`                                                  |
| Int64                  | ✅ Yes      | `INT64`                                                  |
| Interval[DayTime]      | ✅ Yes      | `RECORD<days: INT64, milliseconds: INT64>`               |
| Interval[MonthDayNano] | ✅ Yes      | `RECORD<months: INT64, days: int64, nanoseconds: int64>` |
| Interval[Month]        | ✅ Yes      | `RECORD<months: INT64>`                                  |
| JSON                   | ✅ Yes      | `JSON`                                                   |
| Large Binary           | ✅ Yes      | `BYTES`                                                  |
| Large List             | ✅ Yes      | (Repeated column) †                                      |
| Large String           | ✅ Yes      | `STRING`                                                 |
| List                   | ✅ Yes      | (Repeated column) †                                      |
| MAC                    | ✅ Yes      | `STRING`                                                 |
| Map                    | ❌ No       |                                                          |
| String                 | ✅ Yes      | `STRING`                                                 |
| Struct                 | ✅ Yes      | `RECORD`                                                 |
| Timestamp              | ✅ Yes      | `TIMESTAMP`                                              |
| UUID                   | ✅ Yes      | `STRING`                                                 |
| Uint8                  | ✅ Yes      | `INT64`                                                  |
| Uint16                 | ✅ Yes      | `INT64`                                                  |
| Uint32                 | ✅ Yes      | `INT64`                                                  |
| Uint64                 | ✅ Yes      | `NUMERIC`                                                |
| Union                  | ❌ No       |                                                          |

## Notes

† Repeated columns in BigQuery do not support null values. Right now, if an array contains null values, these null
values will be dropped when writing to BigQuery. Also, because we use `REPEATED` columns to represent lists, lists of
lists are not supported right now.
