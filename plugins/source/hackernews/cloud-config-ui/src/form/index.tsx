import TextField from '@mui/material/TextField';
import Switch from '@mui/material/Switch';
import Stack from '@mui/material/Stack';
import FormHelperText from '@mui/material/FormHelperText';
import InputAdornment from '@mui/material/InputAdornment';
import { Controller, FormProvider, useForm, useWatch } from 'react-hook-form';
import { getYupValidationResolver } from '@cloudquery/cloud-ui';
import { FormValues, getFormValidationSchema, getDefaultStartTime } from '../utils/formSchema';
import { FormMessagePayload } from '@cloudquery/plugin-config-ui-connector';
import { prepareSubmitValues } from '../utils/prepareSubmitValues';
import { pluginUiMessageHandler } from '../utils/messageHandler';
import { useCallback, useMemo } from 'react';

import { LocalizationProvider } from '@mui/x-date-pickers';
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import { DateTimeField } from '@mui/x-date-pickers/DateTimeField';

import {
  FormFieldGroup,
  scrollToFirstFormFieldError,
  useFormCurrentValues,
  useFormSubmit,
} from '@cloudquery/plugin-config-ui-lib';

interface Props {
  initialValues?: FormMessagePayload['init']['initialValues'];
}

export function Form({ initialValues }: Props) {
  const formSchema = useMemo(() => getFormValidationSchema(initialValues), [initialValues]);

  const defaultValues = { ...formSchema.getDefault() };
  // Needed because yup strips the prototype off of the dayjs object
  defaultValues.spec.startTime = getDefaultStartTime(initialValues?.spec?.startTime);

  const formContext = useForm<FormValues>({
    defaultValues,
    resolver: getYupValidationResolver(formSchema),
  });

  const { control, handleSubmit: handleFormSubmit, getValues } = formContext;

  const startTimeEnabled = useWatch({ control, exact: true, name: 'spec.startTimeEnabled' });

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

  return (
    <FormProvider {...formContext}>
      <Stack spacing={2}>
        <FormFieldGroup title={initialValues ? 'Update a source' : 'Create a source'}>
          <Stack>
            <Stack marginBottom={2} spacing={2}>
              <Controller
                control={control}
                name="name"
                render={({ field, fieldState }) => (
                  <TextField
                    error={!!fieldState.error}
                    fullWidth={true}
                    helperText={fieldState.error?.message}
                    label="Source name"
                    disabled={!!initialValues}
                    autoComplete="off"
                    {...field}
                  />
                )}
              />
              <Stack>
                <Controller
                  control={control}
                  name="spec.startTime"
                  render={({ field, fieldState }) => (
                    <LocalizationProvider dateAdapter={AdapterDayjs}>
                      <DateTimeField
                        disableFuture={true}
                        disabled={!startTimeEnabled}
                        label="Start time"
                        slotProps={{
                          textField: {
                            error: !!fieldState.error,
                            name: field.name,
                            InputProps: {
                              endAdornment: (
                                <InputAdornment position="end">
                                  <Controller
                                    control={control}
                                    name="spec.startTimeEnabled"
                                    render={({ field }) => (
                                      <Switch {...field} checked={field.value} />
                                    )}
                                  />
                                </InputAdornment>
                              ),
                            },
                          },
                        }}
                        {...field}
                      />
                      <FormHelperText sx={{ pl: '1em' }}>
                        {fieldState.error?.message ||
                          'The earliest news date that the source should fetch.'}
                      </FormHelperText>
                    </LocalizationProvider>
                  )}
                />
              </Stack>
              <Controller
                control={control}
                name="spec.itemConcurrency"
                render={({ field, fieldState }) => (
                  <TextField
                    error={!!fieldState.error}
                    fullWidth={true}
                    required={true}
                    helperText={
                      fieldState.error?.message ||
                      'Maximum number of news items to fetch concurrently.'
                    }
                    label="Item concurrency"
                    {...field}
                  />
                )}
              />
            </Stack>
          </Stack>
        </FormFieldGroup>
      </Stack>
    </FormProvider>
  );
}
