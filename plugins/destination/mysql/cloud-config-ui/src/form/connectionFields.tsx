import { useEffect, useState } from 'react';

import { getFieldHelperText } from '@cloudquery/cloud-ui';
import { FormFieldGroup, ExclusiveToggle, FormFieldReset } from '@cloudquery/plugin-config-ui-lib';

import Stack from '@mui/material/Stack';
import TextField from '@mui/material/TextField';
import { Controller, useFormContext } from 'react-hook-form';

import { FormValues } from '../utils/formSchema';

const envPlaceholder = '************';

const connectionTypeOptions = [
  {
    label: 'Regular setup',
    value: 'fields',
  },
  {
    label: 'Connection string',
    value: 'string',
  },
];

export function FormConnectionFields() {
  const [connectionStringResetted, setConnectionStringResetted] = useState(false);
  const [passwordResetted, setPasswordResetted] = useState(false);

  const {
    control,
    formState: { defaultValues, submitCount },
    setValue,
    watch,
    trigger,
  } = useFormContext<FormValues>();

  const values = watch();

  const defaultConnectionString = defaultValues?.connectionString;
  const defaultPassword = defaultValues?.password;

  const handleReset = (field: 'connectionString' | 'password') => {
    switch (field) {
      case 'connectionString': {
        setConnectionStringResetted(true);
        setValue('connectionString', '');

        break;
      }
      case 'password': {
        setPasswordResetted(true);
        setValue('password', '');

        break;
      }
      // No default
    }
  };

  const handelCancelReset = (field: 'connectionString' | 'password') => {
    switch (field) {
      case 'connectionString': {
        setConnectionStringResetted(false);
        setValue('connectionString', defaultConnectionString || '');

        break;
      }
      case 'password': {
        setPasswordResetted(false);
        setValue('password', defaultPassword || '');

        break;
      }
      // No default
    }
  };

  useEffect(() => {
    if (submitCount > 0) {
      trigger();
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [values]);

  const defaultConnectionStringIsSecret = defaultConnectionString?.includes('${password}');

  return (
    <FormFieldGroup
      title="Connect to your database"
      subheader="Set up a connection to your MySQL instance."
    >
      <Stack spacing={3}>
        <ExclusiveToggle
          value={values.connectionType}
          onChange={(value: string | number | boolean) =>
            setValue('connectionType', value as 'string' | 'fields')
          }
          options={connectionTypeOptions}
        />
        {values.connectionType === 'string' ? (
          <Controller
            control={control}
            name="connectionString"
            render={({ field, fieldState }) => (
              <Stack direction="row" alignItems="flex-start" spacing={2}>
                <TextField
                  error={!!fieldState.error}
                  fullWidth={true}
                  helperText={getFieldHelperText(
                    fieldState.error?.message,
                    'Connection string to connect to the database. E.g. user:password@localhost:3306/dbname?tls=preferred\u0026readTimeout=1s\u0026writeTimeout=1s',
                  )}
                  label="Connection string"
                  autoComplete="off"
                  required={true}
                  {...field}
                  disabled={defaultConnectionStringIsSecret && !connectionStringResetted}
                  value={
                    defaultConnectionStringIsSecret && !connectionStringResetted
                      ? defaultConnectionString?.replace('${password}', envPlaceholder)
                      : field.value
                  }
                />
                {defaultConnectionStringIsSecret && (
                  <FormFieldReset
                    isResetted={connectionStringResetted}
                    inputSelectorToFocus='input[name="connectionString"]'
                    onCancel={() => handelCancelReset('connectionString')}
                    onReset={() => handleReset('connectionString')}
                    sx={{ minHeight: 55 }}
                  />
                )}
              </Stack>
            )}
          />
        ) : (
          <Stack spacing={2}>
            <Controller
              control={control}
              name="host"
              render={({ field, fieldState }) => (
                <TextField
                  error={!!fieldState.error}
                  fullWidth={true}
                  helperText={getFieldHelperText(
                    fieldState.error?.message,
                    'Host to connect to. E.g. 1.2.3.4 or mydb.host.com. Optional, defaults to empty.',
                  )}
                  label="Host"
                  autoComplete="off"
                  {...field}
                />
              )}
            />
            <Controller
              control={control}
              name="port"
              render={({ field, fieldState }) => (
                <TextField
                  error={!!fieldState.error}
                  fullWidth={true}
                  helperText={getFieldHelperText(
                    fieldState.error?.message,
                    'Port to connect to. Optional, defaults to empty.',
                  )}
                  label="Port"
                  autoComplete="off"
                  {...field}
                />
              )}
            />
            <Controller
              control={control}
              name="database"
              render={({ field, fieldState }) => (
                <TextField
                  error={!!fieldState.error}
                  fullWidth={true}
                  helperText={getFieldHelperText(
                    fieldState.error?.message,
                    'Name of the MySQL database you want to connect to .Optional, defaults to empty.',
                  )}
                  label="Database"
                  autoComplete="off"
                  {...field}
                />
              )}
            />
            <Controller
              control={control}
              name="username"
              render={({ field, fieldState }) => (
                <Stack direction="row" spacing={2}>
                  <TextField
                    error={!!fieldState.error}
                    fullWidth={true}
                    helperText={getFieldHelperText(
                      fieldState.error?.message,
                      'Username to use when authenticating. Optional, defaults to empty.',
                    )}
                    label="Username"
                    autoComplete="off"
                    {...field}
                  />
                </Stack>
              )}
            />
            <Controller
              control={control}
              name="password"
              render={({ field, fieldState }) => (
                <Stack direction="row" alignItems="flex-start" spacing={2}>
                  <TextField
                    error={!!fieldState.error}
                    fullWidth={true}
                    helperText={getFieldHelperText(
                      fieldState.error?.message,
                      'Password to use when authenticating. Optional, defaults to empty.',
                    )}
                    label="Password"
                    autoComplete="off"
                    {...field}
                    type="password"
                    disabled={defaultPassword === '${password}' && !passwordResetted}
                    value={
                      defaultPassword === '${password}' && !passwordResetted
                        ? envPlaceholder
                        : field.value
                    }
                  />
                  {defaultPassword === '${password}' && (
                    <FormFieldReset
                      isResetted={passwordResetted}
                      inputSelectorToFocus='input[name="password"]'
                      onCancel={() => handelCancelReset('password')}
                      onReset={() => handleReset('password')}
                      sx={{ minHeight: 55 }}
                    />
                  )}
                </Stack>
              )}
            />
          </Stack>
        )}
      </Stack>
    </FormFieldGroup>
  );
}
