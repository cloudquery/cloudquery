import TextField from '@mui/material/TextField';
import Stack from '@mui/material/Stack';
import { Controller, FormProvider, useForm } from 'react-hook-form';
import { getYupValidationResolver } from '@cloudquery/cloud-ui';
import { FormValues, formValidationSchema } from '../utils/formSchema';
import { prepareSubmitValues } from '../utils/prepareSubmitValues';
import { pluginUiMessageHandler } from '../utils/messageHandler';
import { useCallback } from 'react';
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
      <Stack spacing={2}>
        <FormFieldGroup title={initialValues ? 'Update a source' : 'Create a source'}>
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
                    label="Source name"
                    disabled={!!initialValues}
                    autoComplete="off"
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
