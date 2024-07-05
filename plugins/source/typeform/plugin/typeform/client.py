import requests


class TypeformClient:
    def __init__(self, access_token, base_url="https://api.typeform.com"):
        self._access_token = access_token
        self._base_url = base_url

    def _get(self, path, params=None):
        url = self._base_url + path
        headers = {"Authorization": f"Bearer {self._access_token}"}
        return requests.get(url, headers=headers, params=params)

    def list_forms(self, page=1):
        params = {"page": page, "page_size": 200}
        resp = self._get("/forms", params=params)
        if resp.status_code != 200:
            raise Exception(f"Failed to list forms: {resp.text}")

        resp = resp.json()
        for form in resp["items"]:
            yield form

        if resp["page_count"] > page:
            yield from self.list_forms(page + 1)

    def list_form_responses(self, *, form_id, since, page=1):
        params = {"page": page, "page_size": 1000, "since": since}
        resp = self._get(f"/forms/{form_id}/responses", params=params)
        if resp.status_code != 200:
            raise Exception(f"Failed to list form responses: {resp.text}")

        resp = resp.json()
        for form in resp["items"]:
            yield form

        if resp["page_count"] > page:
            yield from self.list_form_responses(form_id=form_id, since=since, page=page + 1)
