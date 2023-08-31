from dataclasses import dataclass, field
import square
import square.client

DEFAULT_CONCURRENCY = 100
DEFAULT_QUEUE_SIZE = 10000


@dataclass
class Spec:
    access_token: str
    environment: str
    concurrency: int = field(default=DEFAULT_CONCURRENCY)
    queue_size: int = field(default=DEFAULT_QUEUE_SIZE)

    def validate(self):
        if self.environment not in ["sandbox", "production"]:
            raise Exception("environment must be one of 'sandbox' or 'production'")
        if self.access_token is None:
            raise Exception("access_token must be provided")


class Client:
    def __init__(self, spec: Spec) -> None:
        self._spec = spec
        self._client = square.client.Client(
            access_token=self._spec.access_token, environment=self._spec.environment
        )

    def id(self):
        return "square"

    @property
    def client(self) -> square.client.Client:
        return self._client
