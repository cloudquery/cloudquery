import test, { expect } from '@playwright/test';
import fs from 'node:fs';
import YAML from 'yaml';

test('Submit the form', async ({ page }) => {
  await page.goto('/');
  await page.getByLabel('Host *').click();
  await page.getByLabel('Host *').fill('database.example.com');
  await page.getByLabel('Port').click();
  await page.getByLabel('Port').fill('5432');
  await page.getByLabel('Database *').click();
  await page.getByLabel('Database *').fill('sample_db');
  await page.getByLabel('Username').click();
  await page.getByLabel('Username').fill('john_doe');
  await page.getByLabel('Password').click();
  await page.getByLabel('Password').fill('securePass123');
  await page.getByLabel('SSL').check();
  await page.getByLabel('SSL Mode').click();
  await page.getByRole('option', { name: 'verify-ca' }).click();
  await expect(
    page.getByText(
      `dbtype='postgresql' user='john_doe' password='\${password}' host='database.example.com' dbname='sample_db' port='5432' sslmode='verify-ca'`,
    ),
  ).toBeVisible();
  await page.getByLabel('Log level *').click();
  await page.getByRole('option', { name: 'warn' }).click();
  await page.getByLabel('Batch size *', { exact: true }).click();
  await page.getByLabel('Batch size *', { exact: true }).fill('12');
  await page.getByLabel('Batch size (bytes) *').click();
  await page.getByLabel('Batch size (bytes) *').fill('2500');
  await page.getByLabel('Batch timeout *').click();
  await page.getByLabel('Batch timeout *').fill('120s');
  await page.getByLabel('Migrate mode *').click();
  await page.getByRole('option', { name: 'forced' }).click();
  await page.getByLabel('Write mode *').click();
  await page.getByRole('option', { name: 'append' }).click();
  await page.getByRole('button', { name: 'Submit' }).click();
  const valuesText = await page
    .locator('text=Values:')
    .locator('xpath=following-sibling::*[1]')
    .textContent();

  expect(valuesText).toBeTruthy();

  if (process.env.E2E_TESTS_GENERATE_CONFIG === 'true') {
    const spec = JSON.parse(valuesText as string);
    const destinationConfig = YAML.stringify({
      kind: 'destination',
      spec: {
        name: 'postgresql',
        registry: 'local',
        path: '../postgresql',
        spec: spec.spec,
        write_mode: spec.writeMode,
        migrate_mode: spec.migrateMode,
      },
    });

    const sourceConfig = YAML.stringify({
      kind: 'source',
      spec: {
        name: 'postgresql',
        path: 'cloudquery/postgresql',
        registry: 'cloudquery',
        version: 'v6.2.5',
        destinations: ['postgresql'],
        spec: {
          connection_string: 'test',
        },
        tables: ['*'],
      },
    });

    if (!fs.existsSync('temp')) {
      fs.mkdirSync('temp');
    }

    fs.writeFileSync('./temp/config.yml', `${sourceConfig}---\n${destinationConfig}`);

    fs.writeFileSync(
      './temp/.env',
      `${spec.envs.map((env: { name: string; value: string }) => `${env.name}=${env.value}`).join('\n')}`,
    );
  }
});
