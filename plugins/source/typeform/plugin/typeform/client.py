import requests

REQUEST_TIMEOUT_SECONDS = 60


class TypeformClient:
    def __init__(self, access_token, base_url="https://api.typeform.com"):
        self._access_token = access_token
        self._base_url = base_url

    def _get(self, path, params=None):
        url = self._base_url + path
        headers = {"Authorization": f"Bearer {self._access_token}"}
        return requests.get(
            url, headers=headers, params=params, timeout=REQUEST_TIMEOUT_SECONDS
        )

    def _paginate(self, path, params, what, page=1):
        while True:
            resp = self._get(path, params={**params, "page": page})
            if resp.status_code != 200:
                raise Exception(f"Failed to list {what}: {resp.text}")

            resp = resp.json()
            yield from resp["items"]

            if resp["page_count"] <= page:
                return
            page += 1

    def list_forms(self, page=1):
        yield from self._paginate("/forms", {"page_size": 200}, "forms", page)

    def list_form_responses(self, *, form_id, since, page=1):
        yield from self._paginate(
            f"/forms/{form_id}/responses",
            {"page_size": 1000, "since": since},
            "form responses",
            page,
        )
