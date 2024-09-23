import 'dotenv/config';

import { defineConfig, devices } from '@playwright/test';
const config = require('@cloudquery/plugin-config-ui-lib/dist/e2e-utils/e2e-utils/config');
/**
 * See https://playwright.dev/docs/test-configuration.
 */
export default defineConfig({
  ...config,
});
