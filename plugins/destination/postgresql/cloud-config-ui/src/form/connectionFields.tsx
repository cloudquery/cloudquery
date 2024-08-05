import Box from '@mui/material/Box';
import Switch from '@mui/material/Switch';
import TextField from '@mui/material/TextField';
import FormControlLabel from '@mui/material/FormControlLabel';
import MenuItem from '@mui/material/MenuItem';
import Stack from '@mui/material/Stack';
import { Controller, useFormContext } from 'react-hook-form';
import { FormValues, sslModeValues } from '../utils/formSchema';
import { useEffect, useState } from 'react';
import { FormFieldReset } from './formFieldReset';
import { FormFieldGroup, ExclusiveToggle } from '@cloudquery/plugin-config-ui-lib';
import { getFieldHelperText } from '@cloudquery/cloud-ui';

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
  const [usernameResetted, setUsernameResetted] = useState(false);
  const [passwordResetted, setPasswordResetted] = useState(false);

  const {
    control,
    formState: { defaultValues, submitCount },
    setValue,
    watch,
    trigger,
  } = useFormContext<FormValues>();

  const values = watch();

  const defaultUsername = defaultValues?.username;
  const defaultPassword = defaultValues?.password;

  const handleReset = (field: 'username' | 'password') => {
    switch (field) {
      case 'username': {
        setUsernameResetted(true);
        setValue('username', '');

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

  const handelCancelReset = (field: 'username' | 'password') => {
    switch (field) {
      case 'username': {
        setUsernameResetted(false);
        setValue('username', defaultUsername || '');

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

  return (
    <FormFieldGroup
      title="Connect to your database"
      subheader="Set up a connection to your PostgreSQL instance."
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
              <TextField
                error={!!fieldState.error}
                fullWidth={true}
                helperText={getFieldHelperText(
                  fieldState.error?.message,
                  'Connection string to connect to the database. E.g. postgres://jack:secret@localhost:5432/mydb?sslmode=prefer',
                )}
                label="Connection string"
                autoComplete="off"
                required={true}
                {...field}
              />
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
                    'Host to connect to. E.g. 1.2.3.4 or mydb.host.com.',
                  )}
                  label="Host"
                  autoComplete="off"
                  required={true}
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
                    'Port to connect to. Optional, defaults to 5432.',
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
                    'Name of the PostgreSQL database you want to connect to.',
                  )}
                  label="Database"
                  required={true}
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
                    disabled={defaultUsername === '${username}' && !usernameResetted}
                    value={
                      defaultUsername === '${username}' && !usernameResetted
                        ? envPlaceholder
                        : field.value
                    }
                  />
                  {defaultUsername === '${username}' && (
                    <FormFieldReset
                      isResetted={usernameResetted}
                      inputSelectorToFocus='input[name="username"]'
                      onCancel={() => handelCancelReset('username')}
                      onReset={() => handleReset('username')}
                    />
                  )}
                </Stack>
              )}
            />
            <Controller
              control={control}
              name="password"
              render={({ field, fieldState }) => (
                <Stack direction="row" spacing={2}>
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
                    />
                  )}
                </Stack>
              )}
            />
            <Controller
              control={control}
              name="schemaName"
              render={({ field, fieldState }) => (
                <TextField
                  error={!!fieldState.error}
                  fullWidth={true}
                  helperText={getFieldHelperText(
                    fieldState.error?.message,
                    'Name of the PostgreSQL schema you want to connect to. Optional, defaults to public.',
                  )}
                  label="Schema"
                  autoComplete="off"
                  {...field}
                />
              )}
            />
            <Controller
              control={control}
              name="ssl"
              render={({ field }) => (
                <Box>
                  <FormControlLabel
                    control={<Switch checked={field.value} {...field} />}
                    label="SSL"
                  />
                </Box>
              )}
            />
            {values.ssl && (
              <Controller
                control={control}
                name="sslMode"
                render={({ field, fieldState }) => (
                  <TextField
                    error={!!fieldState.error}
                    fullWidth={true}
                    helperText={getFieldHelperText(
                      fieldState.error?.message,
                      'SSL connections to encrypt client/server communications using TLS protocols for increased security.',
                    )}
                    label="SSL Mode"
                    select={true}
                    SelectProps={{
                      MenuProps: {
                        autoFocus: false,
                        disableAutoFocus: true,
                      },
                    }}
                    {...field}
                  >
                    <MenuItem value={''} hidden={true} />
                    {sslModeValues.map((value) => (
                      <MenuItem key={value} value={value}>
                        {value}
                      </MenuItem>
                    ))}
                  </TextField>
                )}
              />
            )}
          </Stack>
        )}
      </Stack>
    </FormFieldGroup>
  );
}
