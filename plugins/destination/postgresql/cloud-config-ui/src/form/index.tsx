import TextField from '@mui/material/TextField';
import Stack from '@mui/material/Stack';
import MenuItem from '@mui/material/MenuItem';
import { Controller, FormProvider, useForm } from 'react-hook-form';
import { getYupValidationResolver } from '@cloudquery/cloud-ui';
import {
  FormValues,
  formValidationSchema,
  migrateModeValues,
  pgxLogLevelValues,
  writeModeValues,
} from '../utils/formSchema';
import { prepareSubmitValues } from '../utils/prepareSubmitValues';
import { pluginUiMessageHandler } from '../utils/messageHandler';
import { useCallback, useEffect, useState } from 'react';
import { FormConnectionFields } from './connectionFields';
import {
  FormFieldGroup,
  scrollToFirstFormFieldError,
  useFormCurrentValues,
  useFormSubmit,
} from '@cloudquery/plugin-config-ui-lib';

interface Props {
  initialValues: FormValues | undefined;
}

const formDefaultValues = formValidationSchema.getDefault();
const formValidationResolver = getYupValidationResolver(formValidationSchema);

export function Form({ initialValues }: Props) {
  const [specIsValid, setSpecIsValid] = useState(false);

  const formContext = useForm<FormValues>({
    defaultValues: initialValues || formDefaultValues,
    resolver: formValidationResolver,
  });
  const { control, handleSubmit: handleFormSubmit, watch, getValues } = formContext;

  const values = watch();

  const getCurrentValues = useCallback(() => prepareSubmitValues(getValues()), [getValues]);
  useFormCurrentValues(pluginUiMessageHandler, getCurrentValues);

  const handleValidate: Parameters<typeof useFormSubmit>[0] = async () => {
    try {
      const values: FormValues = await new Promise((resolve, reject) => {
        handleFormSubmit(resolve, reject)();
      });

      return {
        values: prepareSubmitValues(values),
      };
    } catch (error) {
      scrollToFirstFormFieldError(Object.keys(error as Record<string, any>));

      return { errors: error as Record<string, any> };
    }
  };

  useFormSubmit(handleValidate, pluginUiMessageHandler);

  useEffect(() => {
    try {
      formValidationSchema.validateSync(getValues(), {
        abortEarly: false,
      });

      setSpecIsValid(true);
    } catch (error: any) {
      const isSpecInvalid = error.inner.some((err: any) => err.path.startsWith('spec.'));
      setSpecIsValid(!isSpecInvalid);
    }
  }, [getValues, values]);

  return (
    <FormProvider {...formContext}>
      <Stack spacing={2}>
        <FormFieldGroup title={initialValues ? 'Update a destination' : 'Create a destination'}>
          <Stack>
            <Stack marginBottom={2}>
              <Controller
                control={control}
                name="name"
                render={({ field, fieldState }) => (
                  <TextField
                    error={!!fieldState.error}
                    fullWidth={true}
                    helperText={fieldState.error?.message}
                    label="Destination name"
                    disabled={!!initialValues}
                    autoComplete="off"
                    {...field}
                  />
                )}
              />
            </Stack>
          </Stack>
          <FormConnectionFields specIsValid={specIsValid} isUpdating={!!initialValues} />
        </FormFieldGroup>
        <FormFieldGroup title="Advanced">
          <Controller
            control={control}
            name="spec.pgxLogLevel"
            render={({ field, fieldState }) => (
              <TextField
                error={!!fieldState.error}
                fullWidth={true}
                helperText={
                  fieldState.error?.message || 'Defines what pgx call events should be logged.'
                }
                label="Log level"
                select={true}
                SelectProps={{
                  MenuProps: {
                    autoFocus: false,
                    disableAutoFocus: true,
                  },
                }}
                required={true}
                {...field}
              >
                {pgxLogLevelValues.map((value) => (
                  <MenuItem key={value} value={value}>
                    {value}
                  </MenuItem>
                ))}
              </TextField>
            )}
          />
          <Controller
            control={control}
            name="spec.batchSize"
            render={({ field, fieldState }) => (
              <TextField
                error={!!fieldState.error}
                fullWidth={true}
                required={true}
                helperText={
                  fieldState.error?.message ||
                  'Maximum number of items that may be grouped together to be written in a single write.'
                }
                label="Batch size"
                {...field}
              />
            )}
          />
          <Controller
            control={control}
            name="spec.batchSizeBytes"
            render={({ field, fieldState }) => (
              <TextField
                error={!!fieldState.error}
                fullWidth={true}
                required={true}
                helperText={
                  fieldState.error?.message ||
                  'Maximum size of items that may be grouped together to be written in a single write.'
                }
                label="Batch size (bytes)"
                {...field}
              />
            )}
          />
          <Controller
            control={control}
            name="spec.batchTimeout"
            render={({ field, fieldState }) => (
              <TextField
                error={!!fieldState.error}
                fullWidth={true}
                required={true}
                helperText={fieldState.error?.message || 'Maximum interval between batch writes.'}
                label="Batch timeout"
                {...field}
              />
            )}
          />
          <Controller
            control={control}
            name="migrateMode"
            render={({ field, fieldState }) => (
              <TextField
                error={!!fieldState.error}
                fullWidth={true}
                required={true}
                helperText={
                  fieldState.error?.message || (
                    <>
                      Specifies the migration mode to use when source tables are changed.{' '}
                      <a
                        href="https://docs.cloudquery.io/docs/reference/destination-spec#migrate_mode"
                        target="_blank"
                        rel="noreferrer"
                      >
                        Learn more
                      </a>
                    </>
                  )
                }
                label="Migrate mode"
                select={true}
                SelectProps={{
                  MenuProps: {
                    autoFocus: false,
                    disableAutoFocus: true,
                  },
                }}
                {...field}
              >
                {migrateModeValues.map((value) => (
                  <MenuItem key={value} value={value}>
                    {value}
                  </MenuItem>
                ))}
              </TextField>
            )}
          />
          <Controller
            control={control}
            name="writeMode"
            render={({ field, fieldState }) => (
              <TextField
                error={!!fieldState.error}
                fullWidth={true}
                required={true}
                helperText={
                  fieldState.error?.message || (
                    <>
                      Specifies the update method to use when inserting rows.{' '}
                      <a
                        href="https://docs.cloudquery.io/docs/reference/destination-spec#write_mode"
                        target="_blank"
                        rel="noreferrer"
                      >
                        Learn more
                      </a>
                    </>
                  )
                }
                label="Write mode"
                select={true}
                SelectProps={{
                  MenuProps: {
                    autoFocus: false,
                    disableAutoFocus: true,
                  },
                }}
                {...field}
              >
                {writeModeValues.map((value) => (
                  <MenuItem key={value} value={value}>
                    {value}
                  </MenuItem>
                ))}
              </TextField>
            )}
          />
        </FormFieldGroup>
      </Stack>
    </FormProvider>
  );
}
