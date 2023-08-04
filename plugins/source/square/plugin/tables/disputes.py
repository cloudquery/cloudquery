import pyarrow as pa
from typing import Any, Generator

from cloudquery.sdk.schema import Column, Table
from cloudquery.sdk.scheduler import TableResolver
from plugin.client import Client
from plugin.oapi import OAPILoader
from cloudquery.sdk.transformers.openapi import oapi_definition_to_columns
from square.api.disputes_api import DisputesApi
from square.http.api_response import ApiResponse


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
        disputes: DisputesApi = client.client.disputes
        cursor = None
        while True:
            response: ApiResponse = disputes.list_disputes(cursor=cursor)
            if response.is_error():
                raise Exception(response)
            for dispute in response.body.get("disputes", []):
                yield dispute
            if response.cursor is None:
                break
