import { Grid, Stack } from '@mui/material';
import { useForm, FormProvider } from 'react-hook-form';
import { FormFieldGroup, FormFieldReset, getYupValidationResolver } from '@cloudquery/cloud-ui';
import { FormValues, formValidationSchema, sslModeValues } from '../utils/formSchema';
import { prepareSubmitValues } from '../utils/prepareSubmitValues';
import { pluginUiMessageHandler } from '../utils/messageHandler';
import { useState } from 'react';
import { useFormSubmit } from '@cloudquery/plugin-config-ui-lib';

import { Connect } from './connect';
import { AWSFormStepper } from './stepper';
import { Guides } from '../components/guides';
import { SelectServices } from './selectServices';

interface Props {
  initialValues: FormValues | undefined;
}

const formDefaultValues = formValidationSchema.getDefault();
const formValidationResolver = getYupValidationResolver(formValidationSchema);

export function Form({ initialValues }: Props) {
  const [activeIndex, setActiveIndex] = useState(0);
  const form = useForm<FormValues>({
    defaultValues: initialValues || formDefaultValues,
    resolver: formValidationResolver,
  });
  const {
    control,
    handleSubmit,
    formState: { defaultValues },
    setValue,
    watch,
  } = form;

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

  // const handleReset = (field: 'username' | 'password') => {
  //   if (field === 'username') {
  //     setUsernameResetted(true);
  //     setValue('spec.username', '');
  //   } else {
  //     setPasswordResetted(true);
  //     setValue('spec.password', '');
  //   }
  // };

  // const handelCancelReset = (field: 'username' | 'password') => {
  //   if (field === 'username') {
  //     setUsernameResetted(false);
  //     setValue('spec.username', defaultUsername || '');
  //   } else {
  //     setPasswordResetted(false);
  //     setValue('spec.password', defaultPassword || '');
  //   }
  // };

  return (
    <Stack gap={5}>
      <FormProvider {...form}>
        <AWSFormStepper activeIndex={activeIndex} />
        <Grid container spacing={2}>
          <Grid item xs={8} md={6}>
            {activeIndex === 0 && <Connect />}
            {activeIndex === 1 && <SelectServices />}
          </Grid>
          <Grid item xs={4} md={6}>
            <Guides />
          </Grid>
        </Grid>
      </FormProvider>
    </Stack>
  );
}
