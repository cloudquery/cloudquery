let envJson: Record<string, any> = {};
try {
  // eslint-disable-next-line @typescript-eslint/no-require-imports, unicorn/prefer-module
  envJson = require('../.env.json');
} catch {
  envJson = {};
}

export { envJson };
