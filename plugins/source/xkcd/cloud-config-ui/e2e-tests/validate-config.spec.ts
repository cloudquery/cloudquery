import test, { expect } from '@playwright/test';
import fs from 'node:fs';
import YAML from 'yaml';

test('Submit the form', async ({ page }) => {
  await page.goto('/');
  await page.getByRole('button', { name: 'Submit' }).click();
  const valuesText = await page
    .locator('text=Values:')
    .locator('xpath=following-sibling::*[1]')
    .textContent();

  expect(valuesText).toBeTruthy();

  if (process.env.E2E_TESTS_GENERATE_CONFIG === 'true') {
    const spec = JSON.parse(valuesText as string);
    const sourceConfig = YAML.stringify({
      kind: 'source',
      spec: {
        name: 'xkcd',
        path: '../cq-source-xkcd',
        registry: 'local',
        destinations: ['postgresql'],
        spec: spec.spec,
        tables: ['*'],
      },
    });

    const destinationConfig = YAML.stringify({
      kind: 'destination',
      spec: {
        name: 'postgresql',
        path: 'cloudquery/postgresql',
        registry: 'cloudquery',
        version: 'v8.2.5',
        spec: {
          connection_string: 'test',
        },
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
