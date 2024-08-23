from datetime import datetime
import json
from typing import List, Generator
import base64
from cloudquery.sdk.schema import Table

import pandas as pd
import structlog
from cloudquery.sdk import message
from cloudquery.sdk import plugin
from cloudquery.sdk import schema
from cloudquery.sdk.scheduler import Scheduler, TableResolver
from cloudquery.sdk.stateclient.stateclient import StateClientBuilder

from plugin.client import Client, Spec

PLUGIN_NAME = "sqlite-python"
PLUGIN_VERSION = "0.0.1"  # {x-release-please-version}


class SQLitePlugin(plugin.Plugin):
    def __init__(self) -> None:
        super().__init__(
            PLUGIN_NAME,
            PLUGIN_VERSION,
            opts=plugin.plugin.Options(
                team="cloudquery", kind="destination", json_schema=Spec.json_schema()
            ),
        )
        self._spec_json = None
        self._spec = None
        self._client = None
        self._logger = structlog.get_logger()

    def set_logger(self, logger) -> None:
        self._logger = logger

    def init(self, spec, no_connection: bool = False):
        if no_connection:
            return
        self._spec_json = json.loads(spec)
        self._spec = Spec(**self._spec_json)
        self._spec.validate()
        self._client = Client(self._spec)

    def write(self, writer: Generator[message.WriteMessage, None, None]) -> None:
        for msg in writer:
            if isinstance(msg, message.WriteMigrateTableMessage):
                self._client.create_table(msg)
            elif isinstance(msg, message.WriteInsertMessage):
                self._client.insert(msg.record)
            elif isinstance(msg, message.WriteDeleteStale):
                self._client.delete_stale(
                    msg.table_name, msg.source_name, msg.sync_time
                )

    def read(self, table: Table) -> Generator[message.ReadMessage, None, None]:
        for record in self._client.read(table):
            yield message.ReadMessage(record)

    def close(self) -> None:
        self._client.close()
