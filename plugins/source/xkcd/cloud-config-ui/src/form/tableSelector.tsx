import { useCallback, useMemo } from 'react';

import tables from '../data/tables.json';
import { useFormContext, useWatch } from 'react-hook-form';
import React from 'react';
import FormHelperText from '@mui/material/FormHelperText';
import { FormValues } from '../utils/formSchema';
import {
  generatePluginTableList,
  generateTablesFromJson,
  TableSelector,
} from '@cloudquery/plugin-config-ui-lib';

function _PluginTableSelector() {
  const {
    control,
    formState: { errors, submitCount },
    setValue,
    trigger,
  } = useFormContext<FormValues>();
  const selectedTables: Record<string, boolean> = useWatch({
    exact: true,
    name: 'tables',
  });

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

  const tableList = useMemo(
    () => generatePluginTableList(generateTablesFromJson(tables as any)),
    [],
  );

  if (tableList.length === 0) {
    return null;
  }

  return (
    <>
      <TableSelector
        disabled={tableList.length < 2}
        errorMessage={errorMessage}
        onChange={handleChange}
        subscribeToTablesValueChange={subscribeToTablesValueChange}
        tableList={tableList}
        value={selectedTables}
      />
      <FormHelperText error={true}>{errors.tables?.message}</FormHelperText>
    </>
  );
}

export const PluginTableSelector = React.memo(_PluginTableSelector);
