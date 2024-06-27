import { TextField } from '@mui/material';
import { Controller, useForm } from 'react-hook-form';
import { useFormSubmit } from '../utils/hooks/useFormSubmit';
import { FormFieldGroup } from '@cloudquery/cloud-ui';
import { formValidationSchema } from './utils';
import { getYupValidationResolver } from '../utils/validation';
import { FormValues } from '../types';

interface Props {
  initialValues: FormValues | undefined;
}

export function Form({ initialValues }: Props) {
  const { control, handleSubmit } = useForm({
    defaultValues: initialValues || formValidationSchema.getDefault(),
    resolver: getYupValidationResolver(formValidationSchema),
  });

  const handleValidate = async () => {
    try {
      const values: FormValues = await new Promise((resolve, reject) => {
        handleSubmit(resolve, reject)();
      });

      return { values };
    } catch (error) {
      return { errors: error as Record<string, any> };
    }
  };

  useFormSubmit(handleValidate);

  return (
    <FormFieldGroup>
      <Controller
        control={control}
        name="spec.email"
        render={({ field, fieldState }) => (
          <TextField
            error={!!fieldState.error}
            fullWidth={true}
            helperText={fieldState.error?.message}
            label="Email"
            {...field}
          />
        )}
      />
      <Controller
        control={control}
        name="spec.name"
        render={({ field, fieldState }) => (
          <TextField
            error={!!fieldState.error}
            fullWidth={true}
            helperText={fieldState.error?.message}
            label="Name"
            {...field}
          />
        )}
      />
    </FormFieldGroup>
  );
}
