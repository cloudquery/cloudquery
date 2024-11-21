SOURCE_PLUGIN_DIRS=$(ls plugins/source)
DESTINATION_PLUGIN_DIRS=$(ls plugins/destination)
TRANSFORMER_PLUGIN_DIRS=$(ls plugins/transformer)

# append source / destination prefixes
SOURCE_PLUGIN_DIRS=$(echo $SOURCE_PLUGIN_DIRS | tr ' ' '\n' | sed 's/^/source\//g' | tr '\n' ' ')
DESTINATION_PLUGIN_DIRS=$(echo $DESTINATION_PLUGIN_DIRS | tr ' ' '\n' | sed 's/^/destination\//g' | tr '\n' ' ')
TRANSFORMER_PLUGIN_DIRS=$(echo $TRANSFORMER_PLUGIN_DIRS | tr ' ' '\n' | sed 's/^/transformer\//g' | tr '\n' ' ')

PLUGIN_DIRS="$SOURCE_PLUGIN_DIRS $DESTINATION_PLUGIN_DIRS $TRANSFORMER_PLUGIN_DIRS"
echo "Updating codegen for plugins: $PLUGIN_DIRS"

PARALLEL_EXISTS=$(which parallel)

generate() {
  plugin=$1
  if ! grep -q "gen:" "plugins/$plugin/Makefile"; then
    echo "no gen target found for $plugin"
    return
  fi

  echo "updating modules for $plugin"
  if [ -d "plugins/$plugin/vendor" ]; then
    (cd "plugins/$plugin" && go mod tidy && go mod vendor)
  else
    (cd "plugins/$plugin" && go mod tidy)
  fi

  echo "Generating code for $plugin"
  (cd "plugins/$plugin" && make gen)
}

if [ -z "$PARALLEL_EXISTS" ]; then
  for plugin in $PLUGIN_DIRS; do
    generate $plugin
  done
  exit 0
fi

export -f generate
echo $PLUGIN_DIRS | tr ' ' '\n' | parallel -j 8 generate