from cloudquery.sdk.schema.table import Table
import pyarrow as pa
from typing import Any, Generator, List

from cloudquery.sdk.schema import Table
from cloudquery.sdk.schema import Column
from cloudquery.sdk.schema import Table
from cloudquery.sdk.scheduler import TableResolver
from plugin.client import Client
from square.api.payments_api import PaymentsApi

class Payments(Table):
    def __init__(self) -> None:
        super().__init__("payments", [
            Column("id", pa.string(), primary_key=True),
            Column("status", pa.string()),
            Column("order_id", pa.string()),
        ])

class PaymentsResolver(TableResolver):
    def __init__(self) -> None:
        super().__init__(table=Payments())
    
    def resolve(self, client: Client, parent_resource) -> Generator[Any, None, None]:
        payments : PaymentsApi = client.client.payments
        cursor = None
        while True:
          response = payments.list_payments(cursor=cursor)
          for payment in response.payments:
            yield payment
          if response.cursor is None:
            break
        return