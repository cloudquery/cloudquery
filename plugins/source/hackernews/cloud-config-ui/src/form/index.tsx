import { useCallback, useMemo } from 'react';

import { getYupValidationResolver, getFieldHelperText } from '@cloudquery/cloud-ui';
import { FormMessagePayload } from '@cloudquery/plugin-config-ui-connector';
import {
  Logo,
  FormFieldGroup,
  scrollToFirstFormFieldError,
  useFormCurrentValues,
  useFormSubmit,
  FormWrapper,
} from '@cloudquery/plugin-config-ui-lib';
import Box from '@mui/material/Box';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import InputAdornment from '@mui/material/InputAdornment';
import Stack from '@mui/material/Stack';
import Switch from '@mui/material/Switch';
import TextField from '@mui/material/TextField';
import Typography from '@mui/material/Typography';
import { LocalizationProvider } from '@mui/x-date-pickers';
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import { DateTimeField } from '@mui/x-date-pickers/DateTimeField';
import { Controller, FormProvider, useForm, useWatch } from 'react-hook-form';

import { PluginTableSelector } from './tableSelector';
import { FormValues, getFormValidationSchema, getDefaultStartTime } from '../utils/formSchema';

import { pluginUiMessageHandler } from '../utils/messageHandler';
import { prepareSubmitValues } from '../utils/prepareSubmitValues';

interface Props {
  initialValues?: FormMessagePayload['init']['initialValues'];
}

export function Form({ initialValues }: Props) {
  const formSchema = useMemo(() => getFormValidationSchema(initialValues), [initialValues]);

  const defaultValues = { ...formSchema.getDefault() };
  // Needed because yup strips the prototype off of the dayjs object
  defaultValues.spec.startTime = getDefaultStartTime(initialValues?.spec?.start_time);

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

  const { formDisabled } = useFormSubmit(handleValidate, pluginUiMessageHandler);

  return (
    <FormWrapper formDisabled={formDisabled}>
      <FormProvider {...formContext}>
        <Stack spacing={2}>
          <Card>
            <CardContent>
              <Stack gap={2}>
                <Box display="flex" justifyContent="space-between" alignItems="center">
                  <Typography variant="h5">Configure source</Typography>
                  <Box display="flex" justifyContent="space-between" alignItems="center" gap={1.5}>
                    <Logo src={`/images/hackernews.webp`} alt="Hacker News" />
                    <Typography variant="body1">Hacker News</Typography>
                  </Box>
                </Box>
                <Stack>
                  <Stack spacing={2}>
                    <Controller
                      control={control}
                      name="name"
                      render={({ field, fieldState }) => (
                        <TextField
                          error={!!fieldState.error}
                          fullWidth={true}
                          helperText={getFieldHelperText(
                            fieldState.error?.message,
                            'Unique destination name that helps identify the destination within your workspace.',
                          )}
                          label="Source name"
                          disabled={!!initialValues}
                          autoComplete="off"
                          {...field}
                        />
                      )}
                    />
                  </Stack>
                </Stack>
              </Stack>
            </CardContent>
          </Card>
          <FormFieldGroup title="Tables">
            <PluginTableSelector />
          </FormFieldGroup>
          <FormFieldGroup title={'Options'}>
            <Stack>
              <Stack spacing={2}>
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
                            helperText: getFieldHelperText(
                              fieldState.error?.message,
                              'The earliest news date that the source should fetch.',
                            ),
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
                    </LocalizationProvider>
                  )}
                />
                <Controller
                  control={control}
                  name="spec.itemConcurrency"
                  render={({ field, fieldState }) => (
                    <TextField
                      error={!!fieldState.error}
                      fullWidth={true}
                      required={true}
                      helperText={getFieldHelperText(
                        fieldState.error?.message,
                        'Maximum number of news items to fetch concurrently. Recommended value is 100.',
                      )}
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
    </FormWrapper>
  );
}
