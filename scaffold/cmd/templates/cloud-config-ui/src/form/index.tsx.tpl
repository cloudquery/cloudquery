import { getFieldHelperText, getYupValidationResolver } from '@cloudquery/cloud-ui';
import { FormFieldGroup, Logo, useFormSubmit } from '@cloudquery/plugin-config-ui-lib';
import Box from '@mui/material/Box';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import Stack from '@mui/material/Stack';
import TextField from '@mui/material/TextField';
import Typography from '@mui/material/Typography';

import { Controller, useForm } from 'react-hook-form';

import { FormValues, formValidationSchema } from '../utils/formSchema';
import { pluginUiMessageHandler } from '../utils/messageHandler';
import { prepareSubmitValues } from '../utils/prepareSubmitValues';

interface Props {
  initialValues: FormValues | undefined;
}

const formDefaultValues = formValidationSchema.getDefault();
const formValidationResolver = getYupValidationResolver(formValidationSchema);

export function Form({ initialValues }: Props) {
  const { control, handleSubmit } = useForm<FormValues>({
    defaultValues: initialValues || formDefaultValues,
    resolver: formValidationResolver,
  });

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

  useFormSubmit(handleValidate, pluginUiMessageHandler);

  return (
    <Stack spacing={2}>
      <Card>
        <CardContent>
          <Stack gap={2}>
            <Box display="flex" justifyContent="space-between" alignItems="center">
              <Typography variant="h5">Configure {`{pluginKind}`}</Typography>
              <Box display="flex" justifyContent="space-between" alignItems="center" gap={1.5}>
                <Logo src={`images/xkcd.webp`} alt="{pluginTitle}" />
                <Typography variant="body1">{`{pluginTitle}`}</Typography>
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
      <FormFieldGroup title="Configration">
        <Controller
          control={control}
          name="token"
          render={({ field, fieldState }) => (
            <TextField
              error={!!fieldState.error}
              fullWidth={true}
              helperText={fieldState.error?.message}
              label="Token"
              {...field}
            />
          )}
        />
      </FormFieldGroup>
    </Stack>
  );
}
