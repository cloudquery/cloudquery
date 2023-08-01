from typing import Any, Generator

import pyarrow as pa
from cloudquery.sdk.scheduler import TableResolver
from cloudquery.sdk.schema import Column
from cloudquery.sdk.schema import Table
from cloudquery.sdk.types import JSONType

from plugin.client import Client
from plugin.tables.form_responses import FormResponsesResolver, FormResponses


class Forms(Table):
    def __init__(self) -> None:
        super().__init__("typeform_forms",
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


class FormsResolver(TableResolver):
    def __init__(self) -> None:
        super().__init__(table=Forms())

    def resolve(self, client: Client, parent_resource) -> Generator[Any, None, None]:
        for form in client.client.list_forms():
            yield form

    @property
    def child_resolvers(self):
        return [
            FormResponsesResolver(),
        ]
