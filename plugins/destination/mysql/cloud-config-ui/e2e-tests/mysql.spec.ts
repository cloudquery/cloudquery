import { Frame, test } from '@playwright/test';

import {
  createPlugin,
  deletePlugin,
  editPlugin,
  getPersistentName,
  login,
} from './MOVE_TO_LIB/plugin-ui-helpers';

test.describe.configure({ mode: 'serial' });

test.describe('MySQL Destination', () => {
  const parameters = {
    pluginNewName: getPersistentName(),
    kind: 'destination' as 'destination',
    pluginName: 'mysql',
    pluginLabel: 'MySQL',
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
        await iframeElement.getByLabel('Host').click();
        await iframeElement
          .getByLabel('Host')
          .fill('mysql-2aafcd8a-cloudquery-c7ec.j.aivencloud.com');
        await iframeElement.getByLabel('Port').click();
        await iframeElement.getByLabel('Port').fill('20188');
        await iframeElement.getByLabel('Database').click();
        await iframeElement.getByLabel('Database').fill('defaultdb');
        await iframeElement.getByLabel('Username').click();
        await iframeElement.getByLabel('Username').fill('avnadmin');
        await iframeElement.getByLabel('Password *', { exact: true }).click();
        await iframeElement
          .getByLabel('Password *', { exact: true })
          .fill(process.env.CQ_CI_PLAYWRIGHT_MYSQL_PASSWORD!);

        await iframeElement.getByLabel('Migrate mode').click();
        await iframeElement.getByRole('option', { name: 'forced' }).click();
        await iframeElement.getByLabel('Write mode').click();
        await iframeElement.getByRole('option', { name: 'overwrite', exact: true }).click();

        await iframeElement.getByRole('button', { name: 'Advanced Sync Options' }).click();
        await iframeElement.getByLabel('Batch size', { exact: true }).click();
        await iframeElement.getByLabel('Batch size', { exact: true }).fill('12');
        await iframeElement.getByLabel('Batch size (bytes)').click();
        await iframeElement.getByLabel('Batch size (bytes)').fill('2500');
      },
    });
  });

  test('edit plugin', async ({ page }) => {
    await editPlugin({
      ...parameters,
      page,
      fillFieldsSteps: async (iframeElement: Frame) => {
        await iframeElement.getByRole('button', { name: 'Advanced Sync Options' }).click();
        await iframeElement.getByLabel('Batch size', { exact: true }).click();
        await iframeElement.getByLabel('Batch size', { exact: true }).fill('22');
        await iframeElement.getByLabel('Batch size (bytes)').click();
        await iframeElement.getByLabel('Batch size (bytes)').fill('2000');
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
