import type { DataType } from '@cloudquery/plugin-sdk-javascript/arrow';
import { Utf8, Timestamp, TimeUnit, Float64, Bool, Int64, Uint64 } from '@cloudquery/plugin-sdk-javascript/arrow';
import type { ColumnResolver } from '@cloudquery/plugin-sdk-javascript/schema/column';
import { createColumn } from '@cloudquery/plugin-sdk-javascript/schema/column';
import { pathResolver } from '@cloudquery/plugin-sdk-javascript/schema/resolvers';
import type { Table, TableResolver } from '@cloudquery/plugin-sdk-javascript/schema/table';
import { createTable } from '@cloudquery/plugin-sdk-javascript/schema/table';
import { JSONType } from '@cloudquery/plugin-sdk-javascript/types/json';
import Airtable from 'airtable';
import { snakeCase } from 'change-case';
import dayjs from 'dayjs';
import customParseFormat from 'dayjs/plugin/customParseFormat.js';
import localizedFormat from 'dayjs/plugin/localizedFormat.js';
import timezone from 'dayjs/plugin/timezone.js';
import utc from 'dayjs/plugin/utc.js';
import { getProperty } from 'dot-prop';
import { got } from 'got';
import pMap from 'p-map';
import type { Logger } from 'winston';

import type { APIField, APITable, APIBase, APIFieldFormula } from './airtable.js';
import { APIFieldType } from './airtable.js';

/* eslint-disable import/no-named-as-default-member */
dayjs.extend(utc);
dayjs.extend(timezone);
dayjs.extend(customParseFormat);
dayjs.extend(localizedFormat);
/* eslint-enable import/no-named-as-default-member */

const timeout = {
  request: 10_000,
};

const options = (apiKey: string) => ({
  // eslint-disable-next-line @typescript-eslint/naming-convention
  headers: { Authorization: `Bearer ${apiKey}` },
  timeout,
});

const getBases = async (apiKey: string, endpointURL: string) => {
  const { bases } = (await got(`${endpointURL}/v0/meta/bases`, options(apiKey)).json()) as { bases: APIBase[] };
  return bases;
};

const getBaseTables = async (apiKey: string, endpointURL: string, baseId: string) => {
  const { tables } = (await got(`${endpointURL}/v0/meta/bases/${baseId}/tables`, options(apiKey)).json()) as {
    tables: APITable[];
  };
  return tables;
};

const airtableFieldToArrowField = (field: APIField): DataType => {
  switch (field.type) {
    case APIFieldType.checkbox: {
      return new Bool();
    }
    case APIFieldType.barcode:
    case APIFieldType.count:
    case APIFieldType.currency:
    case APIFieldType.externalSyncSource:
    case APIFieldType.multipleAttachments:
    case APIFieldType.multipleLookupValues:
    case APIFieldType.multipleRecordLinks:
    case APIFieldType.multipleSelects:
    case APIFieldType.rating:
    case APIFieldType.rollup:
    case APIFieldType.singleCollaborator: {
      return new JSONType();
    }
    case APIFieldType.createdTime:
    case APIFieldType.lastModifiedTime: {
      return airtableFieldToArrowField({ ...field, ...field.options.result });
    }
    // We don't use the Arrow Date type because most destinations don't support it
    case APIFieldType.date:
    case APIFieldType.dateTime: {
      return new Timestamp(TimeUnit.NANOSECOND);
    }
    // The duration in Airtable is saved in the format of `s.m` where `s` is the number of seconds and `m` is the number of milliseconds
    // For example `4980.123` means 4980 seconds and 123 milliseconds which is 1 hour, 23 minutes and 0.123 seconds
    case APIFieldType.duration: {
      return new Float64();
    }
    case APIFieldType.autoNumber: {
      return new Uint64();
    }
    case APIFieldType.number: {
      if (field.options.precision === 0) {
        return new Int64();
      }
      return new Float64();
    }
    case APIFieldType.percent: {
      return new Float64();
    }
    case APIFieldType.formula: {
      const formulaField = field as APIFieldFormula;
      if (formulaField.options.result !== null) {
        return airtableFieldToArrowField(formulaField.options.result);
      }
      return new Utf8();
    }
    default: {
      return new Utf8();
    }
  }
};

const normalizeDateFormat = (format: string) => {
  if (format === 'l' || format === 'LL') {
    return 'YYYY-MM-DD';
  }
  return format;
};

const normalizeTimeZone = (timeZone: string) => {
  if (timeZone === 'client') {
    // 'client' means Airtable uses the local timezone of the user to display the date, and stores it in UTC
    return 'utc';
  }
  return timeZone;
};

const getColumnResolver = (field: APIField): ColumnResolver => {
  switch (field.type) {
    case APIFieldType.createdTime:
    case APIFieldType.lastModifiedTime: {
      const resolver: ColumnResolver = (client, resource, column) => {
        const data = getProperty(resource.getItem(), field.name);
        if (!data) {
          return Promise.resolve(resource.setColumData(column.name, null));
        }

        const dateFormat = normalizeDateFormat(field.options.result.options.dateFormat.format);
        if (field.options.result.type === 'date') {
          const formatted = dayjs(data, dateFormat).toDate();
          return Promise.resolve(resource.setColumData(column.name, formatted));
        }

        const timeFormat = field.options.result.options.timeFormat.format;
        const format = `${dateFormat} ${timeFormat}`;
        const timezone = normalizeTimeZone(field.options.result.options.timeZone);
        const formatted = dayjs.tz(data, format, timezone).toDate();
        return Promise.resolve(resource.setColumData(column.name, formatted));
      };
      return resolver;
    }
    case APIFieldType.date: {
      const resolver: ColumnResolver = (client, resource, column) => {
        const data = getProperty(resource.getItem(), field.name);
        if (!data) {
          return Promise.resolve(resource.setColumData(column.name, null));
        }

        const dateFormat = normalizeDateFormat(field.options.dateFormat.format);
        const formatted = dayjs(data, dateFormat).toDate();
        return Promise.resolve(resource.setColumData(column.name, formatted));
      };
      return resolver;
    }
    case APIFieldType.dateTime: {
      const resolver: ColumnResolver = (client, resource, column) => {
        const data = getProperty(resource.getItem(), field.name);
        if (!data) {
          return Promise.resolve(resource.setColumData(column.name, null));
        }

        const dateFormat = normalizeDateFormat(field.options.dateFormat.format);
        const timeFormat = field.options.timeFormat.format;
        const format = `${dateFormat} ${timeFormat}`;
        const timeZone = normalizeTimeZone(field.options.timeZone);
        const formatted = dayjs.tz(data, format, timeZone).toDate();
        return Promise.resolve(resource.setColumData(column.name, formatted));
      };
      return resolver;
    }
    case APIFieldType.currency: {
      const resolver: ColumnResolver = (client, resource, column) => {
        const data = getProperty(resource.getItem(), field.name);
        if (!data) {
          return Promise.resolve(resource.setColumData(column.name, null));
        }

        const withCurrencySymbol = {
          symbol: field.options.symbol,
          value: data,
        };
        return Promise.resolve(resource.setColumData(column.name, withCurrencySymbol));
      };
      return resolver;
    }
    case APIFieldType.rating: {
      const resolver: ColumnResolver = (client, resource, column) => {
        const data = getProperty(resource.getItem(), field.name);
        if (!data) {
          return Promise.resolve(resource.setColumData(column.name, null));
        }

        const withMaxValue = {
          max: field.options.max,
          value: data,
        };
        return Promise.resolve(resource.setColumData(column.name, withMaxValue));
      };
      return resolver;
    }
    case APIFieldType.formula: {
      const formulaField = field as APIFieldFormula;
      if (formulaField.options.result !== null) {
        return getColumnResolver({ ...field, ...formulaField.options.result });
      }
      return pathResolver(field.name);
    }
    default: {
      return pathResolver(field.name);
    }
  }
};

const fieldToSchemaColumn = (field: APIField) => {
  const { name } = field;
  const normalizedName = snakeCase(name);
  return createColumn({
    name: normalizedName,
    type: airtableFieldToArrowField(field),
    resolver: getColumnResolver(field),
  });
};

const airtableToSchemaTable = (
  apiKey: string,
  endpointUrl: string,
  baseId: string,
  baseName: string,
  table: APITable,
) => {
  // Airtable base names are not unique, so we use the base id for uniqueness of table names
  // Airtable table names are unique within a base, so we use the Airtable table name as the table name
  const name = [baseId.toLowerCase(), snakeCase(baseName), snakeCase(table.name)].join('__');
  const columns = table.fields.map((field) => fieldToSchemaColumn(field));

  const resolver: TableResolver = async (clientMeta, parent, stream) => {
    const airtableClient = new Airtable({ apiKey, endpointUrl }).base(baseId);
    
    return new Promise((resolve, reject) => {
      airtableClient(table.name).select().eachPage(
        (records, fetchNextPage) => {
          for (const record of records) {
            const recordAsObject = Object.fromEntries(table.fields.map((field) => [field.name, record.get(field.name)]));
            stream.write(recordAsObject);
          }
          fetchNextPage();
        },
        (error) => {
          if (error) {
            reject(error);
          } else {
            resolve(undefined);
          }
        }
      );
    });
  };

  return createTable({ name, columns, description: table.description, resolver });
};

export const getTables = async (
  logger: Logger,
  apiKey: string,
  endpointUrl: string,
  concurrency: number,
): Promise<Table[]> => {
  logger.info('discovering Airtable bases');
  const bases = await getBases(apiKey, endpointUrl);
  logger.info(`done discovering Airtable bases. Found ${bases.length} bases`);

  const allTables = await pMap(
    bases,
    async ({ id: baseId, name: baseName }) => {
      logger.info(`discovering tables from Airtable base '(${baseId}) ${baseName}'`);
      const tables = await getBaseTables(apiKey, endpointUrl, baseId);
      logger.info(
        `done discovering tables from Airtable base '(${baseId}) ${baseName}'. Found ${tables.length} tables`,
      );
      return { baseId, baseName, tables };
    },
    {
      concurrency,
    },
  );

  const schemaTables = allTables.map(({ tables, baseId, baseName }) =>
    tables.map((table) => airtableToSchemaTable(apiKey, endpointUrl, baseId, baseName, table)),
  );

  return schemaTables.flat();
};
