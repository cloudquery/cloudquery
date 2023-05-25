# How CloudQuery represents some Arrow types as strings

Some types in Arrow are highly specialized and often there isn't a one to one mapping to a database type. For example, Arrow has a `Date32` type which is a 32-bit integer representing the number of days since the UNIX epoch. This type is not supported by most databases, so CloudQuery has to represent it as a string.

Here is list of all Arrow types and their string representation.

| Arrow Type                           | String Representation                                             | Example                                       |
|--------------------------------------|-------------------------------------------------------------------|-----------------------------------------------| 
| Binary, FixedSizeBinary, LargeBinary | base64 StdEncoding                                                | `YQ==`                                        |
| Boolean                              | `strconv.FormatBool` output                                       | "true"                                        |
| Decimal128, Decimal256               | GetOneForMarshal.(string)                                         | "12345"                                       |
| List, FixedSizeList                  | string(GetOneForMarshal.(json.RawMessage))                        | `[ 1, 2, 3 ]`                                  |
| MonthInterval                        | int32 as string                                                   | "123"                                         |
| DayTimeInterval                      | JSON {"days":int32, "milliseconds":int32}                         | `{"days":1, "milliseconds":234}`              |
| MonthDayNanoInterval                 | JSON {"months":int32, "days":int32, "nanoseconds": int64}         | `{"months":1, "days":2, "nanoseconds": 34567}` |
| Struct                               | JSON                                                              | `{"key": ["values", "value2"]}`               |
| String, LargeString                  | As is                                                             | "foo"                                         |
| Uint8, Uint16, Uint32, Uint64        | strconv.FormatUint                                                | "123"                                         |
| Int8, Int16, Int32, Int64            | strconv.FormatInt                                                 | "-123"                                        |
| Float16                              | GetOneForMarshal.(string)                                         | "123.45"                                      |
| Float32, Float64                     | strconv.FormatFloat                                               | "123.45"                                      |
| Timestamp                            | `YYYY-MM-DD HH:mm:ss.SSSSSSSSS`, no timezone info                 | "2006-01-02 15:04:05.999999999"               |
| Time32                               | `HH:mm:ss` or `HH:mm:ss.SSS`, depending on precision              | "15:04:05.000"                                |
| Time64                               | `HH:mm:ss.SSSSSS` or `HH:mm:ss.SSSSSSSSS`, depending on precision | "15:04:05.000000"                             |
| Date32, Date64                       | `YYYY-MM-DD`                                                      | "2006-01-02"                                  |
| Duration                             | Numeric amount and time unit, concatenated                        | "12345ms"                                     |
| SparseUnion, DenseUnion              | GetOneForMarshal.(string)                                         |                                               | 


 Null values are represented as `(null)`
