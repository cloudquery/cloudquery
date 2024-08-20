import { FormMessagePayload } from '@cloudquery/plugin-config-ui-connector';
import { PluginTable } from '@cloudquery/plugin-config-ui-lib';

import { FormValues, formValidationSchema } from './formSchema';

const formDefaultValues = formValidationSchema.getDefault();

export function prepareInitialValues(
  initialValues: Exclude<FormMessagePayload['init']['initialValues'], undefined> | undefined,
  tablesList: PluginTable[],
): FormValues {
  return {
    name: initialValues?.displayName || '',
    token: initialValues?.spec?.token || '',
    tables: initialValues?.tables
      ? getEnabledTablesObject(initialValues.tables, tablesList)
      : formDefaultValues.tables,
  };
}

const getEnabledTablesObject = (
  tables: string[],
  tablesList: PluginTable[],
): Record<string, boolean> => {
  const enabledTablesObject: Record<string, boolean> = {};

  if (tables.length === 1 && tables[0] === '*') {
    for (const table of tablesList) {
      enabledTablesObject[table.name] = true;
    }

    return enabledTablesObject;
  } else {
    for (const table of tables) {
      if (table !== '*') {
        enabledTablesObject[table] = true;
      }
    }
  }

  return enabledTablesObject;
};
