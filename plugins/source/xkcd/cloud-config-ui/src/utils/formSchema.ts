import { resetYupDefaultErrorMessages } from '@cloudquery/cloud-ui';
import { generateName } from '@cloudquery/plugin-config-ui-lib';
import * as yup from 'yup';

export const existingSecretValue = Symbol('existing-secret-value');

export const sslModeValues = [
  'allow',
  'disable',
  'prefer',
  'require',
  'verify-ca',
  'verify-full',
] as const;
export const pgxLogLevelValues = ['error', 'warn', 'info', 'debug', 'trace'] as const;
export const migrateModeValues = ['forced', 'safe'] as const;
export const writeModeValues = ['append', 'overwrite', 'overwrite-delete-stale'] as const;

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
