import sys
from cloudquery.sdk import serve
from plugin import SquarePlugin


def main():
    p = SquarePlugin()
    serve.PluginCommand(p).run(sys.argv[1:])


if __name__ == "__main__":
    main()
