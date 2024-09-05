import test, { expect } from '@playwright/test';

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

  await page.getByLabel('Batch size', { exact: true }).click();
  await page.getByLabel('Batch size', { exact: true }).fill('12');
  await page.getByLabel('Batch size (bytes)').click();
  await page.getByLabel('Batch size (bytes)').fill('2500');

  await page.getByRole('button', { name: 'Test connection' }).click();

  await expect(page.getByText('Testing the destination connection')).toBeVisible();
});
