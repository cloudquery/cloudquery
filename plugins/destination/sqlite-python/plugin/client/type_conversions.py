import pyarrow as pa


def arrow_type_to_sqlite_str(t: pa.DataType) -> str:
    type_id = t.id
    if type_id in (pa.binary().id, pa.large_binary().id):
        return "blob"
    elif type_id in (pa.string().id, pa.large_string().id):
        return "text"
    elif type_id in (
        pa.int8().id,
        pa.int16().id,
        pa.int32().id,
        pa.int64().id,
        pa.uint8().id,
        pa.uint16().id,
        pa.uint32().id,
        pa.uint64().id,
    ):
        return "integer"
    elif type_id in (pa.float16().id, pa.float32().id, pa.float64().id):
        return "real"
    elif type_id == pa.bool_().id:
        return "boolean"
    elif type_id == pa.timestamp("us").id:
        return "timestamp"
    else:
        return "text"
