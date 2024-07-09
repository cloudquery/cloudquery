from typing import Any, Generator

import pyarrow as pa
from cloudquery.sdk.scheduler import TableResolver
from cloudquery.sdk.schema import Column
from cloudquery.sdk.schema import Table
from cloudquery.sdk.types import JSONType

from plugin.client import Client
from plugin.tables.form_responses import FormResponses


class Forms(Table):
    def __init__(self) -> None:
        super().__init__(
            name="typeform_forms",
            title="Typeform Forms",
            columns=[
                Column("id", pa.string(), primary_key=True),
                Column("created_at", pa.timestamp(unit="s")),
                Column("last_updated_at", pa.timestamp(unit="s")),
                Column("self", JSONType()),
                Column("type", pa.string()),
                Column("settings", JSONType()),
                Column("theme", JSONType()),
                Column("title", pa.string()),
                Column("_links", JSONType()),
            ],
            relations=[FormResponses()],
        )
        self._resolver = FormsResolver(table=self)

    @property
    def resolver(self):
        return self._resolver


class FormsResolver(TableResolver):
    def __init__(self, table=None) -> None:
        super().__init__(table=table)

    def resolve(self, client: Client, parent_resource) -> Generator[Any, None, None]:
        for form in client.client.list_forms():
            yield form

    @property
    def child_resolvers(self):
        return [table.resolver for table in self._table.relations]
