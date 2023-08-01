import json
from typing import List, Generator

from cloudquery.sdk import message
from cloudquery.sdk import plugin
from cloudquery.sdk.scheduler import Scheduler

from plugin import tables
from plugin.client import Client, Spec

PLUGIN_NAME = "typeform"
PLUGIN_VERSION = "0.0.1"


class TypeformPlugin(plugin.Plugin):
    def __init__(self) -> None:
        super().__init__(PLUGIN_NAME, PLUGIN_VERSION)
        self._spec_json = None
        self._spec = None
        self._scheduler = None
        self._client = None

    def init(self, spec_bytes):
        self._spec_json = json.loads(spec_bytes)
        self._spec = Spec(**self._spec_json)
        self._spec.validate()
        self._scheduler = Scheduler(self._spec.concurrency, self._spec.queue_size)
        self._client = Client(self._spec)

    def get_tables(self, options: plugin.TableOptions) -> List[plugin.Table]:
        return [
            tables.Forms(),
        ]

    def sync(self, options: plugin.SyncOptions) -> Generator[message.SyncMessage, None, None]:
        return self._scheduler.sync(self._client, [tables.FormsResolver()])
