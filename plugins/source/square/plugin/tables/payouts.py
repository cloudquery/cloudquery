import pyarrow as pa
from typing import Any, Generator

from cloudquery.sdk.schema import Column, Table
from cloudquery.sdk.scheduler import TableResolver
from plugin.client import Client
from plugin.oapi import OAPILoader
from cloudquery.sdk.transformers.openapi import oapi_definition_to_columns
from square.api.payouts_api import PayoutsApi
from square.http.api_response import ApiResponse


payouts_columns = oapi_definition_to_columns(
    OAPILoader.get_definition("Payout"),
    override_columns=[Column(name="id", type=pa.string(), primary_key=True)],
)


class Payouts(Table):
    def __init__(self) -> None:
        super().__init__(
            name="square_payouts",
            title="Square Payouts",
            columns=payouts_columns,
        )

    @property
    def resolver(self):
        return PayoutsResolver(self)


class PayoutsResolver(TableResolver):
    def __init__(self, table: Table) -> None:
        super().__init__(table=table)

    def resolve(self, client: Client, parent_resource) -> Generator[Any, None, None]:
        payouts: PayoutsApi = client.client.payouts
        cursor = None
        while True:
            response: ApiResponse = payouts.list_payouts(cursor=cursor)
            if response.is_error():
                raise Exception(response)
            for payout in response.body.get("payouts", []):
                yield payout
            if response.cursor is None:
                break
