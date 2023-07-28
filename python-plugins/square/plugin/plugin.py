import json
from cloudquery.sdk import plugin
from cloudquery.sdk.scheduler import Scheduler
from cloudquery.sdk import message
from typing import List, Generator
from plugin.client import Spec
from plugin.client import Client
from plugin import tables


PLUGIN_NAME = "square"
PLUGIN_VERSION = "0.0.1"

class SquarePlugin(plugin.Plugin):
    def __init__(self) -> None:
        super().__init__(PLUGIN_NAME, PLUGIN_VERSION)
    
    def init(self, spec_bytes):
        self._spec_json = json.loads(spec_bytes)
        self._spec = Spec(**self._spec_json)
        self._spec.validate()
        self._scheduler = Scheduler(self._spec.concurrency, self._spec.queue_size)
        self._client = Client(self._spec)

    def get_tables(self, options: plugin.TableOptions) -> List[plugin.Table]:
        return [
            tables.Payments(),
        ]

    def sync(self, options: plugin.SyncOptions) -> Generator[message.SyncMessage, None, None]:
        return self._scheduler.sync(self._spec, [tables.PaymentsResolver()])

