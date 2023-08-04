import pyarrow as pa
from typing import Any, Generator

from cloudquery.sdk.schema import Column, Table, Resource
from cloudquery.sdk.scheduler import TableResolver
from plugin.client import Client
from plugin.oapi import OAPILoader
from cloudquery.sdk.transformers.openapi import oapi_definition_to_columns
from square.api.invoices_api import InvoicesApi
from square.http.api_response import ApiResponse


invoices_columns = oapi_definition_to_columns(
    OAPILoader.get_definition("Invoice"),
    override_columns=[Column(name="id", type=pa.string(), primary_key=True)],
)


class Invoices(Table):
    def __init__(self) -> None:
        super().__init__(
            name="square_invoices",
            title="Square Invoices",
            columns=invoices_columns,
        )

    @property
    def resolver(self):
        return InvoicesResolver(self)


class InvoicesResolver(TableResolver):
    def __init__(self, table: Table) -> None:
        super().__init__(table=table)

    def resolve(
        self, client: Client, parent_resource: Resource
    ) -> Generator[Any, None, None]:
        loc_id = parent_resource.item["id"]
        invoices: InvoicesApi = client.client.invoices
        cursor = None
        while True:
            response: ApiResponse = invoices.list_invoices(
                location_id=loc_id, cursor=cursor
            )
            if response.is_error():
                raise Exception(response)
            for invoice in response.body.get("invoices", []):
                yield invoice
            if response.cursor is None:
                break
