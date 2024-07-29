import Box from '@mui/material/Box';
import Switch from '@mui/material/Switch';
import TextField from '@mui/material/TextField';
import Typography from '@mui/material/Typography';
import FormControlLabel from '@mui/material/FormControlLabel';
import useTheme from '@mui/material/styles/useTheme';
import MenuItem from '@mui/material/MenuItem';
import Stack from '@mui/material/Stack';
import { Controller, useFormContext } from 'react-hook-form';
import { FormValues, sslModeValues } from '../utils/formSchema';
import { useEffect, useState } from 'react';
import { generateConnectionUrl } from '../utils/generateConnectionUrl';
import { FormFieldReset } from './formFieldReset';

const envPlaceholder = '************';

interface Props {
  specIsValid: boolean;
  isUpdating: boolean;
}

export function FormConnectionFields({ specIsValid, isUpdating }: Props) {
  const { palette } = useTheme();
  const [usernameResetted, setUsernameResetted] = useState(false);
  const [passwordResetted, setPasswordResetted] = useState(false);

  const {
    control,
    formState: { defaultValues, submitCount },
    setValue,
    watch,
    getValues,
    trigger,
  } = useFormContext<FormValues>();

  const values = watch();

  const defaultUsername = defaultValues?.spec?.username;
  const defaultPassword = defaultValues?.spec?.password;

  const handleReset = (field: 'username' | 'password') => {
    switch (field) {
      case 'username': {
        setUsernameResetted(true);
        setValue('spec.username', '');

        break;
      }
      case 'password': {
        setPasswordResetted(true);
        setValue('spec.password', '');

        break;
      }
      // No default
    }
  };

  const handelCancelReset = (field: 'username' | 'password') => {
    switch (field) {
      case 'username': {
        setUsernameResetted(false);
        setValue('spec.username', defaultUsername || '');

        break;
      }
      case 'password': {
        setPasswordResetted(false);
        setValue('spec.password', defaultPassword || '');

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
    <>
      <Controller
        control={control}
        name="spec.host"
        render={({ field, fieldState }) => (
          <TextField
            error={!!fieldState.error}
            fullWidth={true}
            helperText={fieldState.error?.message}
            label="Host"
            autoComplete="off"
            required={true}
            {...field}
          />
        )}
      />
      <Controller
        control={control}
        name="spec.port"
        render={({ field, fieldState }) => (
          <TextField
            error={!!fieldState.error}
            fullWidth={true}
            helperText={fieldState.error?.message}
            label="Port"
            autoComplete="off"
            {...field}
          />
        )}
      />
      <Controller
        control={control}
        name="spec.database"
        render={({ field, fieldState }) => (
          <TextField
            error={!!fieldState.error}
            fullWidth={true}
            helperText={fieldState.error?.message}
            label="Database"
            autoComplete="off"
            required={true}
            {...field}
          />
        )}
      />
      <Controller
        control={control}
        name="spec.username"
        render={({ field, fieldState }) => (
          <Stack direction="row" spacing={2}>
            <TextField
              error={!!fieldState.error}
              fullWidth={true}
              helperText={fieldState.error?.message}
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
                inputSelectorToFocus='input[name="spec.username"]'
                onCancel={() => handelCancelReset('username')}
                onReset={() => handleReset('username')}
              />
            )}
          </Stack>
        )}
      />
      <Controller
        control={control}
        name="spec.password"
        render={({ field, fieldState }) => (
          <Stack direction="row" spacing={2}>
            <TextField
              error={!!fieldState.error}
              fullWidth={true}
              helperText={fieldState.error?.message}
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
                inputSelectorToFocus='input[name="spec.password"]'
                onCancel={() => handelCancelReset('password')}
                onReset={() => handleReset('password')}
              />
            )}
          </Stack>
        )}
      />
      <Controller
        control={control}
        name="spec.clientEncoding"
        render={({ field, fieldState }) => (
          <TextField
            error={!!fieldState.error}
            fullWidth={true}
            helperText={fieldState.error?.message}
            label="Client Encoding"
            {...field}
          />
        )}
      />
      <Controller
        control={control}
        name="spec.ssl"
        render={({ field }) => (
          <Box>
            <FormControlLabel control={<Switch checked={field.value} {...field} />} label="SSL" />
          </Box>
        )}
      />
      {values.spec.ssl && (
        <Controller
          control={control}
          name="spec.sslMode"
          render={({ field, fieldState }) => (
            <TextField
              error={!!fieldState.error}
              fullWidth={true}
              helperText={fieldState.error?.message}
              label="SSL Mode"
              select={true}
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
      <Typography variant="h6">Generated connection string</Typography>
      <Box bgcolor={palette.background.paperTertiary} padding={2} borderRadius={1.25}>
        {specIsValid
          ? generateConnectionUrl(
              getValues(),
              isUpdating && defaultUsername === '${username}' && !usernameResetted,
            )
          : 'Fill inputs to see the generated connection string'}
      </Box>
    </>
  );
}
