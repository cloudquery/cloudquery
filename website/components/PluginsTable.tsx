import { Table, Tr, Td, Th } from "nextra/components";
import Link from "next/link";
import { getLatestVersion } from "../utils/versions";

const TableLink = ({ text, href, target = undefined }) => {
  return (
    <Link href={href}>
      <a
        className="nx-text-primary-500 nx-underline nx-decoration-from-font [text-underline-position:under]"
        target={target}
      >
        {text}
      </a>
    </Link>
  );
};

const pluralize = {
  source: "sources",
  destination: "destinations",
};

const TableRow = ({ type, plugin, stage, meta = () => null }) => {
  const pluginLowercase = plugin.toLowerCase();
  return (
    <Tr>
      <Td>
        <TableLink
          text={plugin}
          href={`/docs/plugins/${pluralize[type]}/${pluginLowercase}/overview`}
        />
        {meta()}
      </Td>
      <Td>{getLatestVersion(type, plugin.toLowerCase())}</Td>
      <Td>
        <TableLink
          text="Changelog"
          href={`https://github.com/cloudquery/cloudquery/blob/main/plugins/${type}/${pluginLowercase}/CHANGELOG.md`}
          target="_blank"
        />
      </Td>
      {type === "source" && (
        <Td>
          <TableLink
            text="Tables"
            href={`/docs/plugins/sources/${pluginLowercase}/tables`}
          />
        </Td>
      )}
      <Td>{stage}</Td>
    </Tr>
  );
};

export const PluginsTable = ({ plugins, type }) => {
  const sortedPlugins = plugins.sort((a, b) => a.name.localeCompare(b.name));
  return (
    <Table className="nx-mt-6">
      <Th>
        <strong>Name</strong>
      </Th>
      <Th>Latest Version</Th>
      <Th>Changelog</Th>
      {type === "source" && <Th>Tables</Th>}
      <Th>Stage</Th>
      {sortedPlugins.map(({ name, stage, meta }) => (
        <TableRow
          key={name}
          type={type}
          plugin={name}
          stage={stage}
          meta={meta}
        />
      ))}
    </Table>
  );
};
