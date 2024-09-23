const config = require('@cloudquery/plugin-config-ui-lib/.eslintrc.js');
module.exports = {
  ...config,
  parserOptions: {
    ...config.parserOptions,
    project: './tsconfig.json',
    tsconfigRootDir: __dirname,
  },
};
