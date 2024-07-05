import { FormControlLabel, MenuItem, Stack, Switch, TextField } from '@mui/material';
import { Controller, useForm } from 'react-hook-form';
import { useFormSubmit } from '../utils/hooks/useFormSubmit';
import { FormFieldGroup } from '@cloudquery/cloud-ui';
import { FormValues, formValidationSchema, sslModeValues } from '../utils/formSchema';
import { getYupValidationResolver } from '../utils/validation';
import { FormFieldReset } from './formFieldReset';
import get from 'lodash.get';
import { prepareSubmitValues } from '../utils/prepareSubmitValues';

interface Props {
  initialValues: FormValues | undefined;
}

const secretPlaceholder = '************';

export function Form({ initialValues }: Props) {
  const {
    control,
    handleSubmit,
    formState: { defaultValues },
    setValue,
    watch,
  } = useForm({
    defaultValues: initialValues || formValidationSchema.getDefault(),
    resolver: getYupValidationResolver(formValidationSchema),
  });
  const sslValue = watch('spec.ssl');

  const handleValidate: Parameters<typeof useFormSubmit>[0] = async () => {
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

  useFormSubmit(handleValidate);

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
              disabled={typeof field.value === 'symbol'}
              value={typeof field.value === 'symbol' ? secretPlaceholder : field.value}
            />
            <FormFieldReset
              initialValue={get(defaultValues?.spec, 'username')}
              isReadonly={typeof field.value === 'symbol'}
              path="spec.username"
              setValue={setValue}
            />
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
              disabled={typeof field.value === 'symbol'}
              value={typeof field.value === 'symbol' ? secretPlaceholder : field.value}
            />
            <FormFieldReset
              initialValue={get(defaultValues?.spec, 'password')}
              isReadonly={typeof field.value === 'symbol'}
              path="spec.password"
              setValue={setValue}
            />
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
