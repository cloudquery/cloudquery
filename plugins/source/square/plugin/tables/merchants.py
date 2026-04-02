import pyarrow as pa
from typing import Any, Generator

from cloudquery.sdk.schema import Column, Table
from cloudquery.sdk.scheduler import TableResolver
from plugin.client import Client
from plugin.oapi import OAPILoader
from cloudquery.sdk.transformers.openapi import oapi_definition_to_columns

merchants_columns = oapi_definition_to_columns(
    OAPILoader.get_definition("Merchant"),
    override_columns=[Column(name="id", type=pa.string(), primary_key=True)],
)


class Merchants(Table):
    def __init__(self) -> None:
        super().__init__(
            name="square_merchants",
            title="Square Merchants",
            columns=merchants_columns,
        )

    @property
    def resolver(self):
        return MerchantsResolver(self)


class MerchantsResolver(TableResolver):
    def __init__(self, table: Table) -> None:
        super().__init__(table=table)

    def resolve(self, client: Client, parent_resource) -> Generator[Any, None, None]:
        for merchant in client.client.merchants.list():
            yield merchant.model_dump(mode="json")
