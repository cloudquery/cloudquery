import { expect, Frame, test } from '@playwright/test';

import {
  click,
  createPlugin,
  deletePlugin,
  editPlugin,
  getPersistentName,
  fillInput,
  login,
} from '@cloudquery/plugin-config-ui-lib/e2e-utils';

test.describe.configure({ mode: 'serial' });

test.describe('BigQuery Destination', () => {
  const parameters = {
    pluginNewName: getPersistentName(),
    kind: 'destination' as 'destination',
    pluginName: 'bigquery',
    pluginLabel: 'BigQuery',
    pluginUrl: '',
  };

  test.beforeEach('login', async ({ page }) => {
    await login(page);
  });

  test('create plugin', async ({ page }) => {
    parameters.pluginUrl = await createPlugin({
      ...parameters,
      page,
      fillFieldsSteps: async (iframeElement: Frame) => {
        await click(
          iframeElement,
          iframeElement.getByRole('button', { name: 'Create credentials' }),
        );

        await expect(iframeElement.getByText('Successfully created GCP credentials')).toBeVisible();
        await fillInput(
          iframeElement,
          iframeElement.getByLabel('Fastly API Key *'),
          'gosurf-338418',
        );
        await fillInput(
          iframeElement,
          iframeElement.getByLabel('Fastly API Key *'),
          'my_default_dataset',
        );
      },
    });
  });

  test('edit plugin', async ({ page }) => {
    await editPlugin({
      ...parameters,
      page,
      fillFieldsSteps: async (iframeElement: Frame) => {
        await fillInput(
          iframeElement,
          iframeElement.getByLabel('Fastly API Key *'),
          'my_second_data_set',
        );
      },
    });
  });

  test('delete plugin', async ({ page }) => {
    await deletePlugin({
      ...parameters,
      page,
    });
  });
});
