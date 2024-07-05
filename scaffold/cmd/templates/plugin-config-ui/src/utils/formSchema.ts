import { YupInferType, yup } from './validation';

export const existingSecretValue = Symbol('existing-secret-value');

export const sslModeValues = [
  'allow',
  'disable',
  'prefer',
  'require',
  'verify-ca',
  'verify-full',
] as const;

export const formValidationSchema = yup.object({
  migrateMode: yup.string(),
  writeMode: yup.string(),
  secrets: yup
    .array()
    .of(
      yup.object({
        name: yup.string().default('').required(),
        value: yup.string().default(''),
      }),
    )
    .default([]),
  tables: yup.array().of(yup.string().required()),
  spec: yup.object({
    originalProtocol: yup.string().required().default('postgresql'),
    username: yup.mixed().max(63).default(''),
    password: yup.mixed().max(63).default(''),
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

export type FormValues = YupInferType<typeof formValidationSchema>;
