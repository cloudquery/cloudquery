import { useMemo } from 'react';

import { FormMessagePayload } from '@cloudquery/plugin-config-ui-connector';
import { useCoreFormSchema } from '@cloudquery/plugin-config-ui-lib';

export function useFormSchema({
  initialValues,
}: {
  initialValues?: FormMessagePayload['init']['initialValues'];
}) {
  const formFields = useMemo(() => ({ fields: {} }), [initialValues]);

  return useCoreFormSchema({
    initialValues,
    ...formFields,
  });
}
