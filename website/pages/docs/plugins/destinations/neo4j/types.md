# Neo4j Types

The Neo4j destination (`v3.0.0` and later) supports most [Apache Arrow](https://arrow.apache.org/docs/index.html) types. The following table shows the supported types and how they are mapped to [Neo4j data types](https://neo4j.com/docs/graphql-manual/current/type-definitions/types/).

| Arrow Column Type      | Supported?  | Neo4j Type |
|------------------------|-------------|------------|
| Binary                 | ✅ Yes      | `Bytes`    |
| Boolean                | ✅ Yes      | `Boolean`  |
| Date32                 | ✅ Yes      | `String`   |
| Date64                 | ✅ Yes      | `String`   |
| Decimal                | ✅ Yes      | `String`   |
| Dense Union            | ✅ Yes      | `String`   |
| Dictionary             | ✅ Yes      | `String`   |
| Duration[ms]           | ✅ Yes      | `String`   |
| Duration[ns]           | ✅ Yes      | `String`   |
| Duration[s]            | ✅ Yes      | `String`   |
| Duration[us]           | ✅ Yes      | `String`   |
| Fixed Size List        | ✅ Yes      | `String`   |
| Float16                | ✅ Yes      | `String`   |
| Float32                | ✅ Yes      | `Float`    |
| Float64                | ✅ Yes      | `Float`    |
| Inet                   | ✅ Yes      | `String`   |
| Int8                   | ✅ Yes      | `BigInt`   |
| Int16                  | ✅ Yes      | `BigInt`   |
| Int32                  | ✅ Yes      | `BigInt`   |
| Int64                  | ✅ Yes      | `BigInt`   |
| Interval[DayTime]      | ✅ Yes      | `String`   |
| Interval[MonthDayNano] | ✅ Yes      | `String`   |
| Interval[Month]        | ✅ Yes      | `String`   |
| JSON                   | ✅ Yes      | `String`   |
| Large Binary           | ✅ Yes      | `Bytes`    |
| Large List             | ✅ Yes      | `List`     |
| Large String           | ✅ Yes      | `String`   |
| List                   | ✅ Yes      | `List`     |
| MAC                    | ✅ Yes      | `String`   |
| Map                    | ✅ Yes      | `String`   |
| String                 | ✅ Yes      | `String`   |
| Struct                 | ✅ Yes      | `String`   |
| Timestamp[ms]          | ✅ Yes      | `DateTime` |
| Timestamp[ns]          | ✅ Yes      | `DateTime` |
| Timestamp[s]           | ✅ Yes      | `DateTime` |
| Timestamp[us]          | ✅ Yes      | `DateTime` |
| UUID                   | ✅ Yes      | `String`   |
| Uint8                  | ✅ Yes      | `BigInt`   |
| Uint16                 | ✅ Yes      | `BigInt`   |
| Uint32                 | ✅ Yes      | `BigInt`   |
| Uint64                 | ✅ Yes      | `BigInt`   |
| Union                  | ✅ Yes      | `String`   |

String-persisted data types are encoded according to the [Arrow String Representation](/docs/advanced-topics/arrow-string-representation) specification.
