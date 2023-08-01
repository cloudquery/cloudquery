import json
from cloudquery.sdk import plugin
from cloudquery.sdk.scheduler import Scheduler, TableResolver
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
            tables.Merchants(),
        ]

    def sync(
        self, options: plugin.SyncOptions
    ) -> Generator[message.SyncMessage, None, None]:
        resolvers: list[TableResolver] = []
        for table in self.get_tables(
            plugin.TableOptions(
                tables=options.tables,
                skip_tables=options.skip_tables,
                skip_dependent_tables=options.skip_dependent_tables,
            )
        ):
            resolvers.append(table.resolverClass)

        return self._scheduler.sync(self._client, resolvers)
