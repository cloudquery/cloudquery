import { FormMessagePayload } from '@cloudquery/plugin-config-ui-connector';

type InitialValues = Exclude<FormMessagePayload['init']['initialValues'], undefined>;

export interface FormValues extends InitialValues {
  spec: {
    email: string;
    name: string;
  };
}
