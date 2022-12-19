PLUGIN_DIRS=$(ls plugins/source)
for plugin in $PLUGIN_DIRS; do
	if [ ! -d "plugins/source/$plugin/docs/tables" ]; then
	  continue;
	fi
	echo "Updating docs for $plugin"

	(cd "plugins/source/$plugin" && make gen-docs)
done