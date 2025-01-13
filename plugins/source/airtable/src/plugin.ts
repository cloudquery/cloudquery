import type {
  NewClientFunction,
  TableOptions,
  SyncOptions,
  Plugin,
} from '@cloudquery/plugin-sdk-javascript/plugin/plugin';
import { newPlugin, newUnimplementedDestination } from '@cloudquery/plugin-sdk-javascript/plugin/plugin';
import { sync } from '@cloudquery/plugin-sdk-javascript/scheduler';
import type { Table } from '@cloudquery/plugin-sdk-javascript/schema/table';
import { filterTables } from '@cloudquery/plugin-sdk-javascript/schema/table';
import { readPackageUp } from 'read-package-up';

import { parseSpec, JSON_SCHEMA } from './spec.js';
import type { Spec } from './spec.js';
import { getTables } from './tables.js';

const {
  packageJson: { version },
} = (await readPackageUp()) || { packageJson: { version: 'development' } };

type AirtableClient = {
  id: () => string;
};

export const newAirtablePlugin = () => {
  const pluginClient = {
    ...newUnimplementedDestination(),
    plugin: null as unknown as Plugin,
    spec: null as unknown as Spec,
    client: null as unknown as AirtableClient | null,
    allTables: null as unknown as Table[],
    close: () => Promise.resolve(),
    tables: ({ tables, skipTables, skipDependentTables }: TableOptions) => {
      const { allTables } = pluginClient;
      const filtered = filterTables(allTables, tables, skipTables, skipDependentTables);
      return Promise.resolve(filtered);
    },
    sync: (options: SyncOptions) => {
      const { client, allTables, plugin } = pluginClient;

      if (client === null) {
        return Promise.reject(new Error('Client not initialized'));
      }

      const logger = plugin.getLogger();
      const {
        spec: { concurrency },
      } = pluginClient;

      const { stream, tables, skipTables, skipDependentTables, deterministicCQId } = options;
      const filtered = filterTables(allTables, tables, skipTables, skipDependentTables);

      return sync({
        logger,
        client,
        stream,
        tables: filtered,
        deterministicCQId,
        concurrency,
      });
    },
  };

  const newClient: NewClientFunction = async (logger, spec, { noConnection }) => {
    pluginClient.client = { id: () => 'airtable' };
    if (noConnection) {
      pluginClient.allTables = [];
      return pluginClient;
    }
    pluginClient.spec = parseSpec(spec);
    pluginClient.allTables = await getTables(
      logger,
      pluginClient.spec.apiKey,
      pluginClient.spec.endpointUrl,
      pluginClient.spec.concurrency,
    );

    return pluginClient;
  };

  pluginClient.plugin = newPlugin('airtable', version, newClient, {
    kind: 'source',
    team: 'cloudquery',
    jsonSchema: JSON_SCHEMA,
  });
  return pluginClient.plugin;
};
