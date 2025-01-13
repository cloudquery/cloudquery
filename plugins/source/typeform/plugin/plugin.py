import json
from typing import List, Generator

import structlog
from cloudquery.sdk import message
from cloudquery.sdk import plugin
from cloudquery.sdk import schema
from cloudquery.sdk.scheduler import Scheduler, TableResolver
from cloudquery.sdk.stateclient.stateclient import StateClientBuilder

from plugin import tables
from plugin.client import Client, Spec

PLUGIN_NAME = "typeform"
PLUGIN_VERSION = "1.5.10"  # {x-release-please-version}


class TypeformPlugin(plugin.Plugin):
    def __init__(self) -> None:
        super().__init__(
            PLUGIN_NAME,
            PLUGIN_VERSION,
            opts=plugin.plugin.Options(
                team="cloudquery", kind="source", json_schema=Spec.json_schema()
            ),
        )
        self._spec_json = None
        self._spec = None
        self._scheduler = None
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
        self._scheduler = Scheduler(
            self._spec.concurrency, self._spec.queue_size, logger=self._logger
        )
        self._client = Client(self._spec)

    def get_tables(self, options: plugin.TableOptions) -> List[plugin.Table]:
        all_tables: List[plugin.Table] = [
            tables.Forms(),
        ]

        # set parent table relationships
        for table in all_tables:
            for relation in table.relations:
                relation.parent = table

        # set initial values
        if options.tables is None:
            options.tables = []
        if options.skip_tables is None:
            options.skip_tables = []

        return schema.filter_dfs(all_tables, options.tables, options.skip_tables)

    def sync(
        self, options: plugin.SyncOptions
    ) -> Generator[message.SyncMessage, None, None]:
        state_client = StateClientBuilder.build(backend_options=options.backend_options)
        self._scheduler.set_post_sync_hook(state_client.flush)

        resolvers: list[TableResolver] = []
        for table in self.get_tables(
            plugin.TableOptions(
                tables=options.tables,
                skip_tables=options.skip_tables,
                skip_dependent_tables=options.skip_dependent_tables,
            )
        ):
            resolvers.append(table.resolver)

        for resolver in resolvers:
            resolver.set_state_client(state_client)
            for r in resolver.child_resolvers:
                r.set_state_client(state_client)

        return self._scheduler.sync(
            self._client, resolvers, options.deterministic_cq_id
        )
