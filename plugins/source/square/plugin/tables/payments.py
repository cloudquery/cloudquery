import pyarrow as pa
from typing import Any, Generator

from cloudquery.sdk.schema import Column, Table
from cloudquery.sdk.scheduler import TableResolver
from plugin.client import Client
from plugin.oapi import OAPILoader
from cloudquery.sdk.transformers.openapi import oapi_definition_to_columns
from square.api.payments_api import PaymentsApi
from square.http.api_response import ApiResponse


payments_columns = oapi_definition_to_columns(
    OAPILoader.get_definition("Payment"),
    override_columns=[Column(name="id", type=pa.string(), primary_key=True)],
)


class Payments(Table):
    def __init__(self) -> None:
        super().__init__(
            name="square_payments",
            title="Square Payments",
            columns=payments_columns,
        )

    @property
    def resolver(self):
        return PaymentsResolver(self)


class PaymentsResolver(TableResolver):
    def __init__(self, table: Table) -> None:
        super().__init__(table=table)

    def resolve(self, client: Client, parent_resource) -> Generator[Any, None, None]:
        payments: PaymentsApi = client.client.payments
        cursor = None
        while True:
            response: ApiResponse = payments.list_payments(cursor=cursor)
            if response.is_error():
                raise Exception(response)
            for payment in response.body.get("payments", []):
                yield payment
            if response.cursor is None:
                break
