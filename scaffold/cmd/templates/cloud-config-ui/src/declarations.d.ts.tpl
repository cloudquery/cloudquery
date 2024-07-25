declare module '@cloudquery-plugin/tables' {
  interface PluginTableColumn {
    name: string;
    type: string;
    is_primary_key?: boolean;
    is_primary_key_component?: boolean;
    is_incremental_key?: boolean;
  }

  interface PluginTable {
    name: string;
    title: string;
    description: string;
    columns: PluginTableColumn[];
    relations: PluginTable[];
  }

  const pluginTables: Array<PluginTable>;

  export default pluginTables;
}
