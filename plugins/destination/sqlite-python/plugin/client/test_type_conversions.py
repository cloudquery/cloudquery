import pyarrow as pa
from plugin.client.type_conversions import arrow_type_to_sqlite_str


def test_arrow_type_to_sqlite_str():
    # Test binary types
    assert arrow_type_to_sqlite_str(pa.binary()) == "blob"
    assert arrow_type_to_sqlite_str(pa.large_binary()) == "blob"

    # Test string types
    assert arrow_type_to_sqlite_str(pa.string()) == "text"
    assert arrow_type_to_sqlite_str(pa.large_string()) == "text"

    # Test integer types
    assert arrow_type_to_sqlite_str(pa.int8()) == "integer"
    assert arrow_type_to_sqlite_str(pa.int16()) == "integer"
    assert arrow_type_to_sqlite_str(pa.int32()) == "integer"
    assert arrow_type_to_sqlite_str(pa.int64()) == "integer"
    assert arrow_type_to_sqlite_str(pa.uint8()) == "integer"
    assert arrow_type_to_sqlite_str(pa.uint16()) == "integer"
    assert arrow_type_to_sqlite_str(pa.uint32()) == "integer"
    assert arrow_type_to_sqlite_str(pa.uint64()) == "integer"

    # Test float types
    assert arrow_type_to_sqlite_str(pa.float16()) == "real"
    assert arrow_type_to_sqlite_str(pa.float32()) == "real"
    assert arrow_type_to_sqlite_str(pa.float64()) == "real"

    # Test boolean type
    assert arrow_type_to_sqlite_str(pa.bool_()) == "boolean"

    # Test timestamp type
    assert arrow_type_to_sqlite_str(pa.timestamp("us")) == "timestamp"

    # Test unknown type
    assert arrow_type_to_sqlite_str(pa.null()) == "text"
