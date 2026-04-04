import pyarrow as pa
from typing import Any, Generator

from cloudquery.sdk.schema import Column, Table
from cloudquery.sdk.scheduler import TableResolver
from plugin.client import Client
from plugin.oapi import OAPILoader
from cloudquery.sdk.transformers.openapi import oapi_definition_to_columns

disputes_columns = oapi_definition_to_columns(
    OAPILoader.get_definition("Dispute"),
    override_columns=[Column(name="id", type=pa.string(), primary_key=True)],
)


class Disputes(Table):
    def __init__(self) -> None:
        super().__init__(
            name="square_disputes",
            title="Square Disputes",
            columns=disputes_columns,
        )

    @property
    def resolver(self):
        return DisputesResolver(self)


class DisputesResolver(TableResolver):
    def __init__(self, table: Table) -> None:
        super().__init__(table=table)

    def resolve(self, client: Client, parent_resource) -> Generator[Any, None, None]:
        for dispute in client.client.disputes.list():
            yield dispute.model_dump(mode="json")
