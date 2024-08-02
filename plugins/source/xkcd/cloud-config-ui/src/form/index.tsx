import TextField from '@mui/material/TextField';
import Stack from '@mui/material/Stack';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import Typography from '@mui/material/Typography';
import { Logo } from '@cloudquery/plugin-config-ui-lib';
import Box from '@mui/material/Box';
import { Controller, FormProvider, useForm } from 'react-hook-form';
import { getYupValidationResolver } from '@cloudquery/cloud-ui';
import { FormValues, formValidationSchema } from '../utils/formSchema';
import { prepareSubmitValues } from '../utils/prepareSubmitValues';
import { pluginUiMessageHandler } from '../utils/messageHandler';

import { assetPrefix } from '../utils/constants';
import { useCallback } from 'react';
import { getFieldHelperText } from '@cloudquery/cloud-ui';
import {
  useFormSubmit,
  useFormCurrentValues,
  scrollToFirstFormFieldError,
} from '@cloudquery/plugin-config-ui-lib';

interface Props {
  initialValues: FormValues | undefined;
}

const formDefaultValues = formValidationSchema.getDefault();
const formValidationResolver = getYupValidationResolver(formValidationSchema);

export function Form({ initialValues }: Props) {
  const formContext = useForm<FormValues>({
    defaultValues: initialValues || formDefaultValues,
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
      <Card>
        <CardContent>
          <Stack gap={2}>
            <Box display="flex" justifyContent="space-between" alignItems="center">
              <Typography variant="h5">Configure source</Typography>
              <Box display="flex" justifyContent="space-between" alignItems="center" gap={1.5}>
                <Logo src={`${assetPrefix}/images/xkcd.png`} alt="XKCD" />
                <Typography variant="body1">XKCD</Typography>
              </Box>
            </Box>
            <Stack>
              <Controller
                control={control}
                name="name"
                render={({ field, fieldState }) => (
                  <TextField
                    error={!!fieldState.error}
                    fullWidth={true}
                    helperText={getFieldHelperText(
                      fieldState.error?.message,
                      'Pick a name to help you identify this source.',
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
        </CardContent>
      </Card>
    </FormProvider>
  );
}
