import os
from cloudquery.sdk import serve
from plugin import SquarePlugin


def test_plugin_package():
    p = SquarePlugin()
    cmd = serve.PluginCommand(p)
    cmd.run(["package", "-m", "test", "v1.0.0", "."])
    assert os.path.isfile("dist/tables.json")
    assert os.path.isfile("dist/package.json")
    assert os.path.isfile("dist/docs/overview.md")
    assert os.path.isfile("dist/plugin-square-v1.0.0-linux-amd64.tar")
    assert os.path.isfile("dist/plugin-square-v1.0.0-linux-arm64.tar")
