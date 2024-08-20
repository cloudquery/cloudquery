import sys
from cloudquery.sdk import serve

from plugin import SQLitePlugin


def main():
    p = SQLitePlugin()
    serve.PluginCommand(p).run(sys.argv[1:])


if __name__ == "__main__":
    main()
