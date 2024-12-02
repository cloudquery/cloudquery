from dataclasses import dataclass
import json
from typing import Any, Generator
from uuid import UUID

import pandas as pd

from plugin.sqlite.client import SQLClient
from cloudquery.sdk.schema import Table
import pyarrow as pa
from cloudquery.sdk import message
from cloudquery.sdk.schema.table import CQ_SYNC_TIME_COLUMN, CQ_SOURCE_NAME_COLUMN


@dataclass
class Spec:
    connection_string: str

    def validate(self):
        if self.connection_string is None:
            raise Exception("connection_string must be provided")

    @staticmethod
    def json_schema():
        return """{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://github.com/cloudquery/cloudquery/plugins/destination/sqlite/v2-python/spec",
  "$ref": "#/$defs/Spec",
  "$defs": {
    "Spec": {
      "properties": {
        "connection_string": {
          "type": "string",
          "description": "The connection string to the SQLite database."
        }
    }
  }
}
"""


class Client:
    def __init__(self, spec: Spec) -> None:
        self._spec = spec
        self._sqlite = SQLClient(spec.connection_string)

    def id(self):
        return "sqlite-python"

    @property
    def client(self) -> SQLClient:
        return self._sqlite

    def create_table(self, msg: message.WriteMigrateTableMessage):
        self._sqlite.migrate_client.migrate_tables([msg])

    def insert(self, record: pa.RecordBatch):
        table = Table.from_arrow_schema(record.schema)
        if table.name is None:
            raise ValueError("Missing table name in schema metadata")

        self._sqlite.insert(
            table_name=table.name,
            col_names=[c.name for c in table.columns],
            values=_record_to_sqlite(record),
            primary_keys=table.primary_keys,
        )

    def delete_stale(self, table_name: str, source_name: str, sync_time: Any):
        self._sqlite.delete_stale(
            table_name=table_name,
            source_name=source_name,
            sync_time=sync_time,
            cq_sync_time_column=CQ_SYNC_TIME_COLUMN,
            cq_source_name_column=CQ_SOURCE_NAME_COLUMN,
        )

    def read(self, table: Table) -> Generator[pa.RecordBatch, None, None]:
        schema = table.to_arrow_schema()
        for row in self._sqlite.read(
            table_name=table.name,
            col_names=[c.name for c in table.columns],
        ):
            yield pa.RecordBatch.from_pandas(
                pd.DataFrame([row], columns=[c.name for c in table.columns]),
                schema=schema,
            )

    def close(self):
        self._sqlite.close()


def get_value(field, arr, i):
    if not arr[i].is_valid:
        return None
    data_type = arr.type.id
    if data_type in [
        pa.bool_().id,
        pa.int8().id,
        pa.int16().id,
        pa.int32().id,
        pa.int64().id,
        pa.uint8().id,
        pa.uint16().id,
        pa.uint32().id,
        pa.uint64().id,
        pa.float32().id,
        pa.float64().id,
        pa.string().id,
        pa.binary().id,
        pa.large_binary().id,
    ] or isinstance(arr.type, pa.FixedSizeBinaryType):
        return arr[i].as_py()
    elif data_type == pa.uint64().id:
        return int(arr[i].as_py())
    elif str(field.type) == "uuid":
        return str(UUID(bytes=arr[i].as_py()))
    elif str(field.type) == "json":
        return json.dumps(json.loads(arr[i].as_py()))
    elif isinstance(arr[i], pa.TimestampScalar):
        return pd.Timestamp(
            arr[i].value, unit=arr[i].type.unit, tz=arr[i].type.tz
        ).isoformat()
    else:
        return str(arr[i])


def _record_to_sqlite(record: pa.RecordBatch):
    res = []
    for i in range(record.num_rows):
        row = []
        for j in range(record.num_columns):
            v = get_value(record.schema.field(j), record.column(j), i)
            row.append(v)
        res.append(row)
    return res
