import json
from cloudquery.sdk import plugin
from cloudquery.sdk.scheduler import Scheduler, TableResolver
from cloudquery.sdk import message, schema
from typing import List, Generator
import structlog
from plugin.client import Client, Spec
from plugin import tables


PLUGIN_NAME = "square"
PLUGIN_VERSION = "0.0.1"


class SquarePlugin(plugin.Plugin):
    def __init__(self) -> None:
        super().__init__(PLUGIN_NAME, PLUGIN_VERSION)
        self._logger = structlog.get_logger()

    def set_logger(self, logger) -> None:
        self._logger = logger

    def init(self, spec_bytes):
        self._spec_json = json.loads(spec_bytes)
        self._spec = Spec(**self._spec_json)
        self._spec.validate()
        self._scheduler = Scheduler(self._spec.concurrency, self._spec.queue_size, logger=self._logger)
        self._client = Client(self._spec)

    def get_tables(self, options: plugin.TableOptions) -> List[plugin.Table]:
        t: List[plugin.Table] = [
            tables.Bookings(),
            tables.Disputes(),
            tables.Locations(),
            tables.Merchants(),
            tables.Payments(),
            tables.Payouts(),
            tables.Refunds(),
        ]
        return schema.filter_dfs(t, options.tables, options.skip_tables)

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
            resolvers.append(table.resolver)

        return self._scheduler.sync(
            self._client, resolvers, options.deterministic_cq_id
        )
