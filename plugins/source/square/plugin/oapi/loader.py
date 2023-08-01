import json
import os


class OAPILoader:
    DATA = None

    @classmethod
    def load_spec(cls):
        if cls.DATA is None:
            with open(os.path.join(os.path.dirname(__file__), "api.json"), "r") as f:
                cls.DATA = json.load(f)
        return cls.DATA

    @classmethod
    def get_definition(self, definition):
        data = self.load_spec()
        return data["definitions"][definition]
