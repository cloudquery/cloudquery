import { useMemo } from 'react';

import { FormMessagePayload } from '@cloudquery/plugin-config-ui-connector';
import { useCoreFormSchema } from '@cloudquery/plugin-config-ui-lib';

import * as yup from 'yup';

export function useFormSchema({
  initialValues,
}: {
  initialValues?: FormMessagePayload['init']['initialValues'];
}) {
  const formFields = useMemo(
    () => ({
      fields: {
        item_concurrency: yup
          .number()
          .default(initialValues?.spec?.item_concurrency ?? 100)
          .required(),
        start_time: yup.mixed().default(initialValues?.spec?.start_time),
      },
      secretFields: {
        _startTimeEnabled: yup
          .boolean()
          .default(!initialValues || !!initialValues?.spec?.start_time)
          .required(),
      },
    }),
    [initialValues],
  );

  return useCoreFormSchema({
    initialValues,
    ...formFields,
  });
}
