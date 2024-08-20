import { PluginUiMessagePayload } from '@cloudquery/plugin-config-ui-connector';
import { PluginTable } from '@cloudquery/plugin-config-ui-lib';

import { FormValues } from './formSchema';

export function prepareSubmitValues(
  values: FormValues,
  tablesList: PluginTable[],
): PluginUiMessagePayload['validation_passed']['values'] {
  return {
    name: values.name,
    displayName: values.name,
    spec: {
      token: values.token,
    },
    envs: [],

    // required for source plugin
    tables: getEnabledTablesArray(values.tables, tablesList),
    // skipTables: [],

    // required for destination plugin
    // migrateMode: 'forced',
    // writeMode: 'append',
  };
}

const getEnabledTablesArray = (
  tables: Record<string, boolean>,
  tablesList: PluginTable[],
): string[] => {
  const enabledTables = Object.entries(tables)
    .filter(([, isEnabled]) => !!isEnabled)
    .map(([tableName]) => tableName);

  return enabledTables.length === tablesList.length ? ['*'] : enabledTables;
};
