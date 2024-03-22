## Apache Arrow type conversion

The ClickHouse destination plugin supports most of [Apache Arrow](https://arrow.apache.org/docs/index.html) types.
It uses the same approach as documented
in [ClickHouse reference](https://clickhouse.com/docs/en/sql-reference/formats#data-format-arrow).
The following table shows the supported types and how they are mapped
to [ClickHouse data types](https://clickhouse.com/docs/en/sql-reference/data-types).

:::callout{type="info"}
Unsupported types are always mapped to [String](https://clickhouse.com/docs/en/sql-reference/data-types/string).
:::

<!-- vale off -->

| Arrow Column Type           | ClickHouse Type                                                                    |
|-----------------------------|------------------------------------------------------------------------------------|
| Binary                      | [String](https://clickhouse.com/docs/en/sql-reference/data-types/string)           |
| Binary View                 | [String](https://clickhouse.com/docs/en/sql-reference/data-types/string)           |
| Boolean                     | [Bool](https://clickhouse.com/docs/en/sql-reference/data-types/boolean)            |
| Date32                      | [Date32](https://clickhouse.com/docs/en/sql-reference/data-types/date32)           |
| Date64                      | [DateTime](https://clickhouse.com/docs/en/sql-reference/data-types/datetime)       |
| Decimal128 (Decimal)        | [Decimal](https://clickhouse.com/docs/en/sql-reference/data-types/decimal)         |
| Decimal256                  | [Decimal](https://clickhouse.com/docs/en/sql-reference/data-types/decimal)         |
| Fixed Size Binary           | [FixedString](https://clickhouse.com/docs/en/sql-reference/data-types/fixedstring) |
| Fixed Size List             | [Array](https://clickhouse.com/docs/en/sql-reference/data-types/array)             |
| Float16                     | [Float32](https://clickhouse.com/docs/en/sql-reference/data-types/float)           |
| Float32                     | [Float32](https://clickhouse.com/docs/en/sql-reference/data-types/float)           |
| Float64                     | [Float64](https://clickhouse.com/docs/en/sql-reference/data-types/float)           |
| Int8                        | [Int8](https://clickhouse.com/docs/en/sql-reference/data-types/int-uint)           |
| Int16                       | [Int16](https://clickhouse.com/docs/en/sql-reference/data-types/int-uint)          |
| Int32                       | [Int32](https://clickhouse.com/docs/en/sql-reference/data-types/int-uint)          |
| Int64                       | [Int64](https://clickhouse.com/docs/en/sql-reference/data-types/int-uint)          |
| Large Binary                | [String](https://clickhouse.com/docs/en/sql-reference/data-types/string)           |
| Large List                  | [Array](https://clickhouse.com/docs/en/sql-reference/data-types/array)             |
| Large String                | [String](https://clickhouse.com/docs/en/sql-reference/data-types/string)           |
| List                        | [Array](https://clickhouse.com/docs/en/sql-reference/data-types/array)             |
| Map                         | [Map](https://clickhouse.com/docs/en/sql-reference/data-types/map)                 |
| String                      | [String](https://clickhouse.com/docs/en/sql-reference/data-types/string)           |
| String View                 | [String](https://clickhouse.com/docs/en/sql-reference/data-types/string)           |
| Struct                      | [Tuple](https://clickhouse.com/docs/en/sql-reference/data-types/tuple)             |
| Time32                      | [DateTime64](https://clickhouse.com/docs/en/sql-reference/data-types/datetime64)   |
| Time64                      | [DateTime64](https://clickhouse.com/docs/en/sql-reference/data-types/datetime64)   |
| Timestamp                   | [DateTime64](https://clickhouse.com/docs/en/sql-reference/data-types/datetime64)   |
| UUID (CloudQuery extension) | [UUID](https://clickhouse.com/docs/en/sql-reference/data-types/uuid)               |
| Uint8                       | [UInt8](https://clickhouse.com/docs/en/sql-reference/data-types/int-uint)          |
| Uint16                      | [UInt16](https://clickhouse.com/docs/en/sql-reference/data-types/int-uint)         |
| Uint32                      | [UInt32](https://clickhouse.com/docs/en/sql-reference/data-types/int-uint)         |
| Uint64                      | [UInt64](https://clickhouse.com/docs/en/sql-reference/data-types/int-uint)         |

<!-- vale on -->

:::callout{type="info"}
[Nested](https://clickhouse.com/docs/en/sql-reference/data-types/nested-data-structures/nested) ClickHouse types have their values converted according to the aforementioned rules.
:::
