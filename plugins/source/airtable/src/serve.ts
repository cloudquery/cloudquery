import { createServeCommand } from '@cloudquery/plugin-sdk-javascript/plugin/serve';

import { newAirtablePlugin } from './plugin.js';

export const createAirtableServeCommand = () => createServeCommand(newAirtablePlugin());
