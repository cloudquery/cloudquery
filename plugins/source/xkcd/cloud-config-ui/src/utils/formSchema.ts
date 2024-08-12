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
  tables: yup
    .object()
    .test('valid tables', function (value: Record<string, true>) {
      if (Object.keys(value || {}).filter((key) => value[key]).length === 0) {
        return this.createError({
          message: 'At least one table must be selected',
          path: 'tables',
        });
      }

      return true;
    })
    .default({ xkcd_comics: true }),
});

export type FormValues = yup.InferType<typeof formValidationSchema>;
