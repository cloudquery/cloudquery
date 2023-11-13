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

const TableRow = ({ type, name, id, stage, meta = () => null, openInHub = false }) => {
  const pluginLinkProps = openInHub ? 
  { target: "_blank", href: `https://hub.cloudquery.io/plugins/${type}/cloudquery/${id}` } : { href:`/docs/plugins/${pluralize[type]}/${id}/overview` };
  const tablesLinkProps =  openInHub ? 
  { target: "_blank", href: `https://hub.cloudquery.io/plugins/${type}/cloudquery/${id}/tables` } : { href: `/docs/plugins/sources/${id}/tables` };
  const changeLogProps = openInHub ?
  { target: "_blank", href: `https://hub.cloudquery.io/plugins/${type}/cloudquery/${id}` } : { target: "_blank", href: `https://github.com/cloudquery/cloudquery/blob/main/plugins/${type}/${id}/CHANGELOG.md` };
  return (
    <Tr>
      <Td>
        <TableLink
          text={name}
          {...pluginLinkProps}
        />
        {meta()}
      </Td>
      <Td>{stage.endsWith("(Premium)") ? "PREMIUM" : getLatestVersion(type, id)}</Td>
      <Td>
        <TableLink
          text="Changelog"
          {...changeLogProps}
        />
      </Td>
      {type === "source" && (
        <Td>
          <TableLink
            text="Tables"
            {...tablesLinkProps}
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
      <thead>
        <Th>
          <strong>Name</strong>
        </Th>
        <Th>Latest Version</Th>
        <Th>Changelog</Th>
        {type === "source" && <Th>Tables</Th>}
        <Th>Stage</Th>
      </thead>
      <tbody>
        {sortedPlugins.map(({ name, stage, meta, id = name, openInHub }) => (
          <TableRow
            key={id}
            type={type}
            name={name}
            id={id.toLowerCase()}
            openInHub={openInHub}
            stage={stage}
            meta={meta}
          />
        ))}
      </tbody>
    </Table>
  );
};
