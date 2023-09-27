(cd cli && go get -u github.com/cloudquery/plugin-sdk && go mod tidy)

PLUGIN_DIRS=$(ls plugins/source)
for plugin in $PLUGIN_DIRS; do
	echo "Updating plugin-sdk for $plugin"

	(cd "plugins/source/$plugin" && go get -u github.com/cloudquery/plugin-sdk/v4 && go mod tidy)
done
