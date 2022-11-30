import { promises as fs } from "fs";
import algoliasearch from "algoliasearch";
import path from "path";
import { fileURLToPath } from "url";
import { execa } from "execa";

const __dirname = path.dirname(fileURLToPath(import.meta.url));
const tablesDir = `${__dirname}/algolia_tables`;
const indexFile = `${tablesDir}/index.json`;
const sourcesToSkip = ["test"];

const genTables = async () => {
  await fs.rm(tablesDir, { recursive: true, force: true });
  const sourcesDir = path.resolve(__dirname, "../../plugins/source");
  const sources = (await fs.readdir(sourcesDir, { withFileTypes: true }))
    .filter(
      (dirent) => dirent.isDirectory() && !sourcesToSkip.includes(dirent.name)
    )
    .map((dirent) => `${sourcesDir}/${dirent.name}`);

  await Promise.all(
    sources.map((source) =>
      execa(
        "go",
        [
          "run",
          "main.go",
          "doc",
          "--format",
          "json",
          `${tablesDir}/${path.basename(source)}`,
        ],
        {
          cwd: source,
        }
      )
    )
  );
};

const getTableRelations = (tableData) => {
  let { relations = [] } = tableData;
  for (const relation of relations) {
    relations = [...relations, ...getTableRelations(relation)];
  }
  return relations;
};

const createIndexFile = async () => {
  const tables = (await fs.readdir(tablesDir)).map((source) => ({
    source,
    tableFile: `${tablesDir}/${source}/__tables.json`,
  }));

  const tablesData = await Promise.all(
    tables.map(async ({ source, tableFile }) => {
      const content = await fs.readFile(tableFile, "utf8");
      const sourceTables = JSON.parse(content);
      return sourceTables.map(({ name, relations }) => ({
        source,
        name,
        relations,
      }));
    })
  );

  const tablesJson = tablesData.flat().reduce((acc, tableData) => {
    const { source, name } = tableData;
    const relations = getTableRelations(tableData).map(({ name }) => ({
      source,
      name,
    }));
    return [...acc, { source, name }, ...relations];
  }, []);

  await fs.writeFile(indexFile, JSON.stringify(tablesJson, null, 2));
};

const indexToAlgolia = async () => {
  try {
    const client = algoliasearch(
      "0OUG5EUZ6H",
      process.env.CQ_ALGOLIA_API_ADMIN_KEY // admin key is required to replace all objects
    );
    const tables = JSON.parse(await fs.readFile(indexFile, "utf8"));
    const withIds = tables.map((table) => ({
      objectID: table.name,
      ...table,
    }));
    const index = client.initIndex("all-sources-tables");
    await index.replaceAllObjects(withIds);
  } catch (err) {
    console.log(err);
    process.exitCode = 1;
  }
};

const main = async () => {
  await genTables();
  await createIndexFile();
  await indexToAlgolia();
};

main();
