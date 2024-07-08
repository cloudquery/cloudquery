import json
import socket
from http.server import BaseHTTPRequestHandler, HTTPServer
from threading import Thread
from urllib.parse import parse_qs

import requests

from plugin.typeform.client import TypeformClient


# Modified version of example in https://realpython.com/testing-third-party-apis-with-mock-servers/
class MockServerRequestHandler(BaseHTTPRequestHandler):
    def do_GET(self):
        parts = self.path.split("?")
        if parts[0] == "/forms":
            query = parse_qs(parts[1])
            self._handle_forms(q=query)
        elif parts[0] == "/forms/form1/responses":
            query = parse_qs(parts[1])
            self._handle_forms_responses(q=query)
        else:
            self.send_response(requests.codes.not_found)
            self.end_headers()
            self.wfile.write(b"Not found")
        return

    def _handle_forms(self, q):
        self.send_response(requests.codes.ok)
        self.end_headers()
        if q["page"] == ["1"]:
            resp = {
                "items": [
                    {
                        "id": "form1",
                        "title": "Form 1",
                    },
                ],
                "page_count": 2,
            }
            self.wfile.write(json.dumps(resp).encode("utf-8"))
        else:
            resp = {
                "items": [
                    {
                        "id": "form2",
                        "title": "Form 2",
                    },
                ],
                "page_count": 2,
            }
            self.wfile.write(json.dumps(resp).encode("utf-8"))

    def _handle_forms_responses(self, q):
        self.send_response(requests.codes.ok)
        self.end_headers()
        if q["page"] == ["1"]:
            resp = {
                "items": [
                    {
                        "answers": [],
                        "submitted_at": "2017-09-14T22:38:22Z",
                    },
                ],
                "page_count": 2,
            }
            self.wfile.write(json.dumps(resp).encode("utf-8"))
        else:
            resp = {
                "items": [
                    {
                        "answers": [],
                        "submitted_at": "2017-09-14T22:33:56Z",
                    },
                ],
                "page_count": 2,
            }
            self.wfile.write(json.dumps(resp).encode("utf-8"))


def get_free_port():
    s = socket.socket(socket.AF_INET, type=socket.SOCK_STREAM)
    s.bind(("localhost", 0))
    address, port = s.getsockname()
    s.close()
    return port


class TestMockServer(object):
    @classmethod
    def setup_class(cls):
        # Configure mock server.
        cls.mock_server_port = get_free_port()
        cls.mock_server = HTTPServer(
            ("localhost", cls.mock_server_port), MockServerRequestHandler
        )

        # Start running mock server in a separate thread.
        # Daemon threads automatically shut down when the main process exits.
        cls.mock_server_thread = Thread(target=cls.mock_server.serve_forever)
        cls.mock_server_thread.daemon = True
        cls.mock_server_thread.start()

    def test_list_forms(self):
        client = TypeformClient(
            base_url="http://localhost:{}".format(self.mock_server_port),
            access_token="fake",
        )
        forms = list(client.list_forms())
        assert len(forms) == 2
        assert forms[0]["id"] == "form1"
        assert forms[1]["id"] == "form2"

    def test_list_forms_responses(self):
        client = TypeformClient(
            base_url="http://localhost:{}".format(self.mock_server_port),
            access_token="fake",
        )
        forms = list(client.list_form_responses("form1"))
        assert len(forms) == 2
        assert forms[0]["submitted_at"] == "2017-09-14T22:38:22Z"
        assert forms[1]["submitted_at"] == "2017-09-14T22:33:56Z"
