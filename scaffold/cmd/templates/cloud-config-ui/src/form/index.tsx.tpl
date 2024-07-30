import TextField from '@mui/material/TextField';
import { Controller, useForm } from 'react-hook-form';
import { getYupValidationResolver } from '@cloudquery/cloud-ui';
import { FormValues, formValidationSchema } from '../utils/formSchema';
import { prepareSubmitValues } from '../utils/prepareSubmitValues';
import { pluginUiMessageHandler } from '../utils/messageHandler';
import { FormFieldGroup, useFormSubmit } from '@cloudquery/plugin-config-ui-lib';

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
    <FormFieldGroup title={initialValues ? 'Update a {{.Kind}}' : 'Create a {{.Kind}}'}>
      <Controller
        control={control}
        name="name"
        render={({ field, fieldState }) => (
          <TextField
            error={!!fieldState.error}
            fullWidth={true}
            helperText={fieldState.error?.message}
            label="{{.Kind}} name"
            disabled={!!initialValues}
            autoComplete="off"
            {...field}
          />
        )}
      />
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
  );
}
