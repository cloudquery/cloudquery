from cloudquery.sdk.schema.table import Table
import pyarrow as pa
from typing import Any, Generator, List

from cloudquery.sdk.schema import Table
from cloudquery.sdk.schema import Column
from cloudquery.sdk.schema import Table
from cloudquery.sdk.scheduler import TableResolver
from plugin.client import Client
from square.api.merchants_api import MerchantsApi


class Merchants(Table):
    def __init__(self) -> None:
        super().__init__(
            "merchants",
            [
                Column("id", pa.string(), primary_key=True),
                Column("business_name", pa.string()),
                Column("country", pa.string()),
                Column("language_code", pa.string()),
                Column("currency", pa.string()),
                Column("status", pa.string()),
                Column("main_location_id", pa.string()),
                Column("created_at", pa.timestamp("ms")),
            ],
        )

    @property
    def resolverClass(self):
        return MerchantsResolver(self)


class MerchantsResolver(TableResolver):
    def __init__(self, table) -> None:
        super().__init__(table=table)

    def resolve(self, client: Client, parent_resource) -> Generator[Any, None, None]:
        merchants: MerchantsApi = client.client.merchants
        cursor = None
        while True:
            response = merchants.list_merchants(cursor=cursor)
            for merchant in response.merchants:
                yield merchant
            if response.cursor is None:
                break
        return
