import { resetYupDefaultErrorMessages } from '@cloudquery/cloud-ui';
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

export enum SetupType {
  Manual = 'manual',
  Console = 'console',
}
export const setupTypes = Object.values(SetupType);

resetYupDefaultErrorMessages(yup);

// TODO: the shape of this, per the API
export const formValidationSchema = yup.object({
  name: yup.string().default('').required(),
  arn: yup.string().default('').required(),
  regions: yup.array().of(yup.string().required()).default([]),
  services: yup.array().of(yup.string().required()).default([]),
  _setupType: yup.string().oneOf(setupTypes).default(SetupType.Console),

  migrateMode: yup
    .string()
    .oneOf(['forced', 'safe'] as const)
    .default('safe'),
  writeMode: yup
    .string()
    .oneOf(['append', 'overwrite', 'overwrite-delete-stale'] as const)
    .default('append'),
  envs: yup
    .array()
    .of(
      yup.object({
        name: yup.string().default('').required(),
        value: yup.string().default(''),
      }),
    )
    .default([]),
  tables: yup.array().of(yup.string().required()).default(['*']),
  spec: yup.object({
    originalProtocol: yup.string().required().default('postgresql'),
    username: yup.string().max(63).default(''),
    password: yup.string().max(63).default(''),
    host: yup.string().max(253).required().default(''),
    port: yup
      .string()
      .max(5)
      .matches(/^($)|(\d+)$/, 'Port must be a number')
      .default(''),
    database: yup.string().max(63).required().default(''),
    clientEncoding: yup.string().max(255).default(''),
    ssl: yup.bool().default(false),
    sslMode: yup.string().oneOf(sslModeValues).default('require'),
  }),
});

export type FormValues = yup.InferType<typeof formValidationSchema>;
