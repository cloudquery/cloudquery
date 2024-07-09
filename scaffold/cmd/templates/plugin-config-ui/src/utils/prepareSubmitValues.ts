import { PluginUiMessagePayload } from '@cloudquery/plugin-config-ui-connector';
import { FormValues } from './formSchema';

export function prepareSubmitValues(
  values: FormValues,
): PluginUiMessagePayload['validation_passed']['values'] {
  const url = new URL('https://cloudquery.io');
  url.protocol = values.spec.originalProtocol;
  url.hostname = values.spec.host;
  url.port = values.spec.port;
  url.pathname = encodeURIComponent(values.spec.database || '');

  if (values.spec.ssl && values.spec.sslMode) {
    url.searchParams.set('sslmode', values.spec.sslMode);
  }
  if (values.spec.clientEncoding) {
    url.searchParams.set('client_encoding', values.spec.clientEncoding);
  }

  const envs = [] as Array<{ name: string; value: string }>;

  const userInfo = [
    values.spec.username ? '${username}' : '',
    values.spec.password ? '${password}' : '',
  ]
    .filter(Boolean)
    .join(':');

  if (values.spec.username) {
    envs.push({
      name: 'username',
      value: values.spec.username === '${username}' ? '' : values.spec.username,
    });
  }
  if (values.spec.password) {
    envs.push({
      name: 'password',
      value: values.spec.password === '${password}' ? '' : values.spec.password,
    });
  }

  let finalUrl = url.toString().replace(/^[a-z]*:\/\//, `${values.spec.originalProtocol}://`);

  if (userInfo) {
    finalUrl = finalUrl.replace(
      `${values.spec.originalProtocol}://`,
      `${values.spec.originalProtocol}://${userInfo}@`,
    );
  }

  return {
    migrateMode: values.migrateMode,
    writeMode: values.writeMode,
    envs,
    tables: ['*'],
    spec: {
      connection_string: finalUrl,
    },
  };
}
