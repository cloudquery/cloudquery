import React, { useCallback, useMemo } from 'react';

import { TableSelector, generatePluginTableList } from '@cloudquery/plugin-config-ui-lib';
import FormHelperText from '@mui/material/FormHelperText';
import Stack from '@mui/system/Stack';
import { useFormContext, useWatch } from 'react-hook-form';

interface PluginTable {
  /** Description of the table */
  description: string;
  /** Whether the table is incremental */
  is_incremental: boolean;
  /** Whether the table is paid */
  is_paid?: boolean;
  name: string;
  /** Name of the parent table, if any */
  parent?: string;
  /** Names of the tables that depend on this table */
  relations: string[];
  /** Title of the table */
  title: string;
}

interface Props {
  pluginTables: PluginTable[];
}

function _PluginTableSelector({ pluginTables }: Props) {
  const {
    control,
    formState: { errors, submitCount },
    setValue,
    trigger,
  } = useFormContext();
  const selectedTables: Record<string, boolean> = useWatch({
    exact: true,
    name: 'tables',
  });

  const tableList = useMemo(() => generatePluginTableList(pluginTables), [pluginTables]);

  const handleChange = useCallback(
    (value: Record<string, boolean>) => {
      setValue('tables', value);
      trigger('tables');
    },
    [setValue, trigger],
  );
  const errorMessage = submitCount > 0 ? (errors?.tables?.message as any) : null;

  const subscribeToTablesValueChange = useCallback(
    (callback: (value: Record<string, boolean>) => void) => {
      const { unsubscribe } = control._subjects.values.subscribe({
        next(payload) {
          callback(payload.values.tables);
        },
      });

      return unsubscribe;
    },
    [control],
  );

  if (tableList.length === 0) {
    return null;
  }

  return (
    <Stack gap={1}>
      <TableSelector
        errorMessage={errorMessage}
        onChange={handleChange}
        subscribeToTablesValueChange={subscribeToTablesValueChange}
        tableList={tableList}
        value={selectedTables}
      />
      <FormHelperText error={!!errorMessage}>{errorMessage}</FormHelperText>
    </Stack>
  );
}

export const PluginTableSelector = React.memo(_PluginTableSelector);
