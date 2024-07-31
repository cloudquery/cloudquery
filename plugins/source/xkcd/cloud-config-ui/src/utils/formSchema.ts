import { resetYupDefaultErrorMessages } from '@cloudquery/cloud-ui';
import { generateName } from '@cloudquery/plugin-config-ui-lib';
import * as yup from 'yup';
resetYupDefaultErrorMessages(yup);

export const formValidationSchema = yup.object({
  name: yup
    .string()
    .default(generateName('xkcd'))
    .matches(
      /^[a-z](-?[\da-z]+)+$/,
      'Name must consist of a lower case letter, followed by alphanumeric segments separated by single dashes',
    )
    .max(255)
    .required(),
  envs: yup
    .array()
    .of(
      yup.object({
        name: yup.string().default('').required(),
        value: yup.string().default(''),
      }),
    )
    .default([]),
  spec: yup.object({}),
});

export type FormValues = yup.InferType<typeof formValidationSchema>;
