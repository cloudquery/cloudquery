import { createServeCommand } from '@cloudquery/plugin-sdk-javascript/plugin/serve';

import { newAirtablePlugin } from './plugin.js';

const main = () => {
  createServeCommand(newAirtablePlugin()).parse();
};

main();
