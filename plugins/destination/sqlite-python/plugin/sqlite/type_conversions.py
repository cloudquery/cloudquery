from pyarrow import (
    DataType,
    binary,
    large_binary,
    string,
    large_string,
    int64,
    int8,
    int16,
    int32,
    uint8,
    uint16,
    uint32,
    uint64,
    float16,
    float32,
    float64,
    bool_,
    timestamp,
)


def arrow_type_to_sqlite_str(t: DataType) -> str:
    if t.equals(binary()) or t.equals(large_binary()):
        return "blob"
    elif t.equals(string()) or t.equals(large_string()):
        return "text"
    elif (
        t.equals(int8())
        or t.equals(int16())
        or t.equals(int32())
        or t.equals(int64())
        or t.equals(uint8())
        or t.equals(uint16())
        or t.equals(uint32())
        or t.equals(uint64())
    ):
        return "integer"
    elif t.equals(float16()) or t.equals(float32()) or t.equals(float64()):
        return "real"
    elif t.equals(bool_()):
        return "boolean"
    elif t.equals(timestamp("us")):
        return "timestamp"
    else:
        return "text"


def arrow_type_to_sqlite(t: DataType) -> DataType:
    if t.equals(binary()) or t.equals(large_binary()):
        return large_string()
    elif t.equals(string()) or t.equals(large_string()):
        return large_string()
    elif (
        t.equals(int8())
        or t.equals(int16())
        or t.equals(int32())
        or t.equals(int64())
        or t.equals(uint8())
        or t.equals(uint16())
        or t.equals(uint32())
        or t.equals(uint64())
    ):
        return int64()
    elif t.equals(float16()) or t.equals(float32()) or t.equals(float64()):
        return float64()
    elif t.equals(bool_()):
        return bool_()
    elif t.equals(timestamp("us")):
        return timestamp("us")
    else:
        return large_string()


def sqlite_type_to_arrow_type(t: str) -> DataType:
    if t == "integer":
        return int64()
    elif t == "real":
        return float64()
    elif t == "text":
        return large_string()
    elif t == "blob":
        return large_binary()
    elif t == "boolean":
        return bool_()
    elif t == "timestamp":
        return timestamp("us")
    else:
        raise ValueError(f"unknown type: {t}")
