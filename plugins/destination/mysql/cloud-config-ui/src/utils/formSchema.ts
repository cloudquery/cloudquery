import { useMemo } from 'react';

import { FormMessagePayload } from '@cloudquery/plugin-config-ui-connector';
import { secretFieldValue, useCoreFormSchema } from '@cloudquery/plugin-config-ui-lib';

import * as yup from 'yup';

import { convertConnectionStringToFields } from './convertConnectionStringToFields';

export const connectionTypeValues = ['string', 'fields'] as const;
export const tlsModeValues = ['true', 'false', 'skip-verify', 'preferred'] as const;
export const migrateModeValues = ['forced', 'safe'] as const;
export const writeModeValues = ['append', 'overwrite', 'overwrite-delete-stale'] as const;

export function useFormSchema({
  initialValues,
}: {
  initialValues?: FormMessagePayload['init']['initialValues'];
}) {
  const formFields = useMemo(() => {
    const url = initialValues?.spec?.connection_string || '';
    const connectionObj: Record<string, any> = convertConnectionStringToFields(url);
    const partialSecretRegex = /\${[^}]+}/g;

    return {
      secretFields: {
        password: yup
          .string()
          .max(63)
          .default(connectionObj.password ? secretFieldValue : ''),
      },
      fields: {
        connection_string: yup
          .string()
          .default(url.replaceAll(partialSecretRegex, secretFieldValue))
          .when('_connectionType', {
            is: 'string',
            // eslint-disable-next-line unicorn/no-thenable
            then: (schema) => schema.required(),
          }),
        username: yup
          .string()
          .max(63)
          .default(connectionObj.username ?? ''),
        host: yup
          .string()
          .max(253)
          .default(connectionObj.host ?? ''),
        port: yup
          .string()
          .max(5)
          .matches(/^($)|(\d+)$/, 'Port must be a number')
          .default(connectionObj.port ?? ''),
        database: yup
          .string()
          .max(63)
          .default(connectionObj.database ?? ''),
        tcp: yup.bool().default(connectionObj.tcp ?? true),
        connectionParams: yup.object({
          tls: yup.bool().default(connectionObj.connectionParams?.tls ?? false),
          tlsMode: yup
            .string()
            .oneOf(tlsModeValues)
            .default(connectionObj.connectionParams?.tlsMode ?? 'preferred')
            .when('tls', {
              is: (tls: boolean) => !tls,
              // eslint-disable-next-line unicorn/no-thenable
              then: (schema: any) => schema.strip(),
            }),
          parseTime: yup.bool().default(connectionObj.connectionParams?.parseTime ?? false),
          charset: yup.string().default(''),
          loc: yup
            .string()
            .default(connectionObj.connectionParams?.loc ?? '')
            .when('parseTime', {
              is: (parseTime: boolean) => !parseTime,
              // eslint-disable-next-line unicorn/no-thenable
              then: (schema: any) => schema.strip(),
            }),
          timeout: yup
            .number()
            .integer()
            .default(connectionObj.connectionParams?.timeout ?? 0),
          readTimeout: yup
            .number()
            .integer()
            .default(connectionObj.connectionParams?.readTimeout ?? 0),
          writeTimeout: yup
            .number()
            .integer()
            .default(connectionObj.connectionParams?.writeTimeout ?? 0),
          allowNativePasswords: yup
            .bool()
            .default(connectionObj.connectionParams?.allowNativePasswords ?? false),
        }),
        // spec
        batch_size: yup
          .number()
          .integer()
          .default(initialValues?.spec?.batch_size ?? 10_000)
          .required(),
        batch_size_bytes: yup
          .number()
          .integer()
          .default(initialValues?.spec?.batch_size_bytes ?? 100_000_000)
          .required(),

        // destination settings
        migrateMode: yup
          .string()
          .oneOf(migrateModeValues)
          .default(initialValues?.migrateMode ?? 'safe')
          .required(),
        writeMode: yup
          .string()
          .oneOf(writeModeValues)
          .default(initialValues?.writeMode ?? 'overwrite-delete-stale')
          .required(),
      },
      stateFields: {
        _connectionType: yup
          .string()
          .oneOf(connectionTypeValues)
          .default(connectionTypeValues[1])
          .required(),
      },
    };
  }, [initialValues]);

  return useCoreFormSchema({
    initialValues,
    ...formFields,
  });
}
