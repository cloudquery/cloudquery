import { FormControlLabel, MenuItem, Stack, Switch, TextField } from '@mui/material';
import { Controller, useForm } from 'react-hook-form';
import {
  FormFieldGroup,
  FormFieldReset,
  getYupValidationResolver,
  usePluginUiFormSubmit,
} from '@cloudquery/cloud-ui';
import { FormValues, formValidationSchema, sslModeValues } from '../utils/formSchema';
import { prepareSubmitValues } from '../utils/prepareSubmitValues';
import { pluginUiMessageHandler } from '../utils/messageHandler';
import { useState } from 'react';

interface Props {
  initialValues: FormValues | undefined;
}

const envPlaceholder = '************';

const formDefaultValues = formValidationSchema.getDefault();
const formValidationResolver = getYupValidationResolver(formValidationSchema);

export function Form({ initialValues }: Props) {
  const [usernameResetted, setUsernameResetted] = useState(false);
  const [passwordResetted, setPasswordResetted] = useState(false);

  const {
    control,
    handleSubmit,
    formState: { defaultValues },
    setValue,
    watch,
  } = useForm<FormValues>({
    defaultValues: initialValues || formDefaultValues,
    resolver: formValidationResolver,
  });

  const sslValue = watch('spec.ssl');

  const handleValidate: Parameters<typeof usePluginUiFormSubmit>[0] = async () => {
    try {
      const values: FormValues = await new Promise((resolve, reject) => {
        handleSubmit(resolve, reject)();
      });

      return {
        values: prepareSubmitValues(values),
      };
    } catch (error) {
      return { errors: error as Record<string, any> };
    }
  };

  usePluginUiFormSubmit(handleValidate, pluginUiMessageHandler);

  const defaultUsername = defaultValues?.spec?.username;
  const defaultPassword = defaultValues?.spec?.password;

  const handleReset = (field: 'username' | 'password') => {
    if (field === 'username') {
      setUsernameResetted(true);
      setValue('spec.username', '');
    } else {
      setPasswordResetted(true);
      setValue('spec.password', '');
    }
  };

  const handelCancelReset = (field: 'username' | 'password') => {
    if (field === 'username') {
      setUsernameResetted(false);
      setValue('spec.username', defaultUsername || '');
    } else {
      setPasswordResetted(false);
      setValue('spec.password', defaultPassword || '');
    }
  };

  return (
    <FormFieldGroup title="PostgreSQL Connection">
      <Controller
        control={control}
        name="spec.host"
        render={({ field, fieldState }) => (
          <TextField
            error={!!fieldState.error}
            fullWidth={true}
            helperText={fieldState.error?.message}
            label="Host"
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
              {...field}
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
        render={({ field }) => <FormControlLabel control={<Switch {...field} />} label="SSL" />}
      />
      {sslValue && (
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
    </FormFieldGroup>
  );
}
