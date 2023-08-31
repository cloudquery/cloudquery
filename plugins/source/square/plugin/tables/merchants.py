import pyarrow as pa
from typing import Any, Generator

from cloudquery.sdk.schema import Column, Table
from cloudquery.sdk.scheduler import TableResolver
from plugin.client import Client
from square.api.merchants_api import MerchantsApi
from square.http.api_response import ApiResponse
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
        merchants: MerchantsApi = client.client.merchants
        cursor = None
        while True:
            response: ApiResponse = merchants.list_merchants(cursor=cursor)
            if response.is_error():
                raise Exception(response)
            for merchant in response.body.get("merchant", []):
                yield merchant
            if response.cursor is None:
                break
