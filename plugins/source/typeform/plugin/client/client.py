from dataclasses import dataclass, field

from plugin.typeform.client import TypeformClient

DEFAULT_CONCURRENCY = 100
DEFAULT_QUEUE_SIZE = 10000


@dataclass
class Spec:
    access_token: str
    base_url: str = field(default="https://api.typeform.com")
    concurrency: int = field(default=DEFAULT_CONCURRENCY)
    queue_size: int = field(default=DEFAULT_QUEUE_SIZE)

    def validate(self):
        if self.access_token is None:
            raise Exception("access_token must be provided")


class Client:
    def __init__(self, spec: Spec) -> None:
        self._spec = spec
        self._client = TypeformClient(spec.access_token, spec.base_url)

    def id(self):
        return "typeform"

    @property
    def client(self) -> TypeformClient:
        return self._client
