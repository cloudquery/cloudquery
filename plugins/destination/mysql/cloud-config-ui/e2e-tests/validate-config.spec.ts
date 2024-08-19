import test, { expect } from '@playwright/test';
import fs from 'node:fs';
import YAML from 'yaml';

test('Submit the form', async ({ page }) => {
  await page.goto('/');
  await page.getByLabel('Host').click();
  await page.getByLabel('Host').fill('localhost');
  await page.getByLabel('Port').click();
  await page.getByLabel('Port').fill('3306');
  await page.getByLabel('Database').click();
  await page.getByLabel('Database').fill('sample_db');
  await page.getByLabel('Username').click();
  await page.getByLabel('Username').fill('john_doe');
  await page.getByLabel('Password').click();
  await page.getByLabel('Password').fill('securePass123');

  await page.getByRole('button', { name: 'Advanced Connection Options' }).click();
  await page.getByLabel('TCP').click();
  await page.getByLabel('TLS').click();
  await page.getByLabel('TLS Mode').click();
  await page.getByRole('option', { name: 'preferred' }).click();
  await page.getByLabel('Parse Time').click();
  await page.getByLabel('Location').click();
  await page.getByLabel('Location').fill('Local');
  await page.getByLabel('Charset').click();
  await page.getByLabel('Charset').fill('utf8mb4');
  await page.getByLabel('Timeout', { exact: true }).click();
  await page.getByLabel('Timeout', { exact: true }).fill('6');
  await page.getByLabel('Read Timeout').click();
  await page.getByLabel('Read Timeout').fill('7');
  await page.getByLabel('Write Timeout').click();
  await page.getByLabel('Write Timeout').fill('8');

  await page.getByLabel('Migrate mode *').click();
  await page.getByRole('option', { name: 'forced' }).click();
  await page.getByLabel('Write mode *').click();
  await page.getByRole('option', { name: 'overwrite', exact: true }).click();

  await page.getByRole('button', { name: 'Advanced Sync Options' }).click();

  await page.getByLabel('Batch size *', { exact: true }).click();
  await page.getByLabel('Batch size *', { exact: true }).fill('12');
  await page.getByLabel('Batch size (bytes) *').click();
  await page.getByLabel('Batch size (bytes) *').fill('2500');

  await page.getByRole('button', { name: 'Submit' }).click();

  const valuesText = await page
    .locator('text=Values:')
    .locator('xpath=following-sibling::*[1]')
    .textContent();
  expect(valuesText).toBeTruthy();

  const spec = JSON.parse(valuesText as string);
  expect(spec.spec.connection_string).toBe(
    'john_doe:${password}@tcp(localhost:3306)/sample_db?tlsMode=preferred&parseTime=True&charset=utf8mb4&loc=Local&timeout=6s&readTimeout=7s&writeTimeout=8s',
  );

  if (process.env.E2E_TESTS_GENERATE_CONFIG === 'true') {
    const destinationConfig = YAML.stringify({
      kind: 'destination',
      spec: {
        name: 'mysql',
        registry: 'local',
        path: '../mysql',
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
