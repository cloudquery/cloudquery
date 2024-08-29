import { useCallback, useMemo } from 'react';

import { getYupValidationResolver } from '@cloudquery/cloud-ui';
import { FormMessagePayload } from '@cloudquery/plugin-config-ui-connector';
import {
  FormProvider,
  useForm,
  GuideComponent,
  ConfigUIForm,
  usePluginContext,
} from '@cloudquery/plugin-config-ui-lib';

import { Stack, Box } from '@mui/system';

import { useFormSchema } from '../utils/formSchema';

import { pluginUiMessageHandler } from '../utils/messageHandler';

import { prepareSubmitValues } from '../utils/prepareSubmitValues';

interface Props {
  initialValues?: FormMessagePayload['init']['initialValues'] | undefined;
}

export function Form({ initialValues }: Props) {
  const { tablesList } = usePluginContext();

  const formValidationSchema = useFormSchema({ initialValues });
  const formValidationResolver = getYupValidationResolver(formValidationSchema);

  const values = useMemo(() => {
    return formValidationSchema.getDefault();
  }, [formValidationSchema]);

  const form = useForm({
    defaultValues: values,
    values,
    resolver: formValidationResolver,
  });

  const getCurrentValues = useCallback(
    () => prepareSubmitValues(form.getValues(), tablesList),
    [form, tablesList],
  );

  return (
    <FormProvider {...form}>
      <Stack direction="row" gap={3} flexWrap="wrap">
        <Box flex="1 1 0" minWidth={480}>
          <ConfigUIForm
            getCurrentValues={getCurrentValues}
            pluginUiMessageHandler={pluginUiMessageHandler}
          />
        </Box>
        <Box sx={{ width: 360, minWidth: 360 }}>
          <GuideComponent pluginUiMessageHandler={pluginUiMessageHandler} />
        </Box>
      </Stack>
    </FormProvider>
  );
}
