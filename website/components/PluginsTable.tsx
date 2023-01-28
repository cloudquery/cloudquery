import { Table, Tr, Td, Th } from "nextra/components";
import Link from "next/link";
import { getLatestVersion } from "../utils/versions";

const TableLink = ({ text, href, target = undefined }) => {
  return (
    <Link
      href={href}
      className="nx-text-primary-500 nx-underline nx-decoration-from-font [text-underline-position:under]"
      target={target}
      >
        {text}
      </Link>
  );
};

const pluralize = {
  source: "sources",
  destination: "destinations",
};

const TableRow = ({ type, name, id, stage, meta = () => null }) => {
  return (
    <Tr>
      <Td>
        <TableLink
          text={name}
          href={`/docs/plugins/${pluralize[type]}/${id}/overview`}
        />
        {meta()}
      </Td>
      <Td>{getLatestVersion(type, id)}</Td>
      <Td>
        <TableLink
          text="Changelog"
          href={`https://github.com/cloudquery/cloudquery/blob/main/plugins/${type}/${id}/CHANGELOG.md`}
          target="_blank"
        />
      </Td>
      {type === "source" && (
        <Td>
          <TableLink
            text="Tables"
            href={`/docs/plugins/sources/${id}/tables`}
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
      {sortedPlugins.map(({ name, stage, meta, id = name }) => (
        <TableRow
          key={id}
          type={type}
          name={name}
          id={id.toLowerCase()}
          stage={stage}
          meta={meta}
        />
      ))}
    </Table>
  );
};
