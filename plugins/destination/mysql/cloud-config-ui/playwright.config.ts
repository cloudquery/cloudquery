import 'dotenv/config';

import { defineConfig } from '@playwright/test';
const { playwrightConfig } = require('@cloudquery/plugin-config-ui-lib/dist/e2e-utils');

/**
 * See https://playwright.dev/docs/test-configuration.
 */
export default defineConfig({
  ...playwrightConfig,
});
