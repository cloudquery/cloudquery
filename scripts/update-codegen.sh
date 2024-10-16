PLUGIN_DIRS=$(ls plugins/source)
for plugin in $PLUGIN_DIRS; do
	if ! grep -q "gen:" "plugins/source/$plugin/Makefile"; then
	  continue;
	fi
	echo "Generating code for $plugin"

	(cd "plugins/source/$plugin" && make gen)
done
