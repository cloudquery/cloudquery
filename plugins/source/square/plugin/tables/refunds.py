import pyarrow as pa
from typing import Any, Generator

from cloudquery.sdk.schema import Column, Table
from cloudquery.sdk.scheduler import TableResolver
from plugin.client import Client
from plugin.oapi import OAPILoader
from cloudquery.sdk.transformers.openapi import oapi_definition_to_columns
from square.api.refunds_api import RefundsApi
from square.http.api_response import ApiResponse


refunds_columns = oapi_definition_to_columns(
    OAPILoader.get_definition("Refund"),
    override_columns=[Column(name="id", type=pa.string(), primary_key=True)],
)


class Refunds(Table):
    def __init__(self) -> None:
        super().__init__(
            name="square_refunds",
            title="Square Refunds",
            columns=refunds_columns,
        )

    @property
    def resolver(self):
        return RefundsResolver(self)


class RefundsResolver(TableResolver):
    def __init__(self, table: Table) -> None:
        super().__init__(table=table)

    def resolve(self, client: Client, parent_resource) -> Generator[Any, None, None]:
        refunds: RefundsApi = client.client.refunds
        cursor = None
        while True:
            response: ApiResponse = refunds.list_payment_refunds(cursor=cursor)
            if response.is_error():
                raise Exception(response)
            for refund in response.body.get("refunds", []):
                yield refund
            if response.cursor is None:
                break
