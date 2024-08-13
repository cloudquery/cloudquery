import { useCallback } from 'react';

import { getFieldHelperText, getYupValidationResolver } from '@cloudquery/cloud-ui';
import {
  Logo,
  scrollToFirstFormFieldError,
  useFormCurrentValues,
  useFormSubmit,
} from '@cloudquery/plugin-config-ui-lib';
import Box from '@mui/material/Box';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import Stack from '@mui/material/Stack';
import TextField from '@mui/material/TextField';
import Typography from '@mui/material/Typography';
import { Controller, FormProvider, useForm } from 'react-hook-form';

import { AdvancedFields } from './advancedFields';
import { FormConnectionFields } from './connectionFields';
import { FormSyncOptions } from './syncOptions';
import { FormValues, formValidationSchema } from '../utils/formSchema';
import { pluginUiMessageHandler } from '../utils/messageHandler';
import { prepareSubmitValues } from '../utils/prepareSubmitValues';

interface Props {
  initialValues: FormValues | undefined;
  isManagedDestination: boolean;
}

const formDefaultValues = formValidationSchema.getDefault();
const formValidationResolver = getYupValidationResolver(formValidationSchema);

export function Form({ initialValues, isManagedDestination }: Props) {
  const formContext = useForm<FormValues>({
    defaultValues: initialValues || {
      ...formDefaultValues,
      connectionType: isManagedDestination ? 'string' : 'fields',
    },
    resolver: formValidationResolver,
  });
  const { control, handleSubmit: handleFormSubmit, getValues } = formContext;

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
        <Card>
          <CardContent>
            <Box display="flex" marginBottom={3} justifyContent="space-between" alignItems="center">
              <Typography variant="h5">Configure destination</Typography>
              <Box display="flex" justifyContent="space-between" alignItems="center" gap={1.5}>
                <Logo src="images/mysql.webp" alt="MySQL" />
                <Typography variant="body1">MySQL</Typography>
              </Box>
            </Box>
            <Stack marginBottom={2}>
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
                    label="Destination name"
                    disabled={!!initialValues}
                    autoComplete="off"
                    {...field}
                  />
                )}
              />
            </Stack>
          </CardContent>
        </Card>
        <FormConnectionFields />
        <FormSyncOptions />
        <AdvancedFields />
      </Stack>
    </FormProvider>
  );
}
