import { resetYupDefaultErrorMessages } from '@cloudquery/cloud-ui';
import { generateName } from '@cloudquery/plugin-config-ui-lib';
import * as yup from 'yup';

resetYupDefaultErrorMessages(yup);

export const formValidationSchema = yup.object({
  /** This is the name of plugin source/destination */
  name: yup
    .string()
    .default(generateName('{{.Name}}'))
    .matches(
      /^[a-z](-?[\da-z]+)+$/,
      'Name must consist of a lower case letter, followed by alphanumeric segments separated by single dashes',
    )
    .max(255)
    .required(),
  token: yup.string().default('').required(),
});

export type FormValues = yup.InferType<typeof formValidationSchema>;
