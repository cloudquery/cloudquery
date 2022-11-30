import { promises as fs } from "fs";
import path from "path";
import { fileURLToPath } from "url";
import { execa } from "execa";

const __dirname = path.dirname(fileURLToPath(import.meta.url));
const tablesJSONsDir = `${__dirname}/tables`;
const docsSourcesDir = path.resolve(__dirname, "../pages/docs/plugins/sources");
const sourcesToSkip = ["test"];

const genTablesJSONFiles = async () => {
  await fs.rm(tablesJSONsDir, { recursive: true, force: true });
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
          `${tablesJSONsDir}/${path.basename(source)}`,
        ],
        {
          cwd: source,
        }
      )
    )
  );
};

const getTablesListWithIndent = (tables, indent = 0) => {
  const indentString = " ".repeat(indent);
  let withIndent = [];

  for (const table of tables) {
    withIndent = [...withIndent, `${indentString}- ${table.name}`];
    withIndent = [
      ...withIndent,
      ...getTablesListWithIndent(table.relations, indent + 2),
    ];
  }

  return withIndent;
};

const createTablesMarkdowns = async () => {
  const tables = (await fs.readdir(tablesJSONsDir)).map((source) => ({
    source,
    tableFile: `${tablesJSONsDir}/${source}/__tables.json`,
  }));

  const tablesData = await Promise.all(
    tables.map(async ({ source, tableFile }) => {
      const content = await fs.readFile(tableFile, "utf8");
      const tables = JSON.parse(content);
      return { source, tables };
    })
  );

  await Promise.all(
    tablesData.map(async ({ source, tables }) => {
      const markdownFile = `${docsSourcesDir}/${source}/tables.md`;
      const tablesWithIndent = getTablesListWithIndent(tables).join("\n");
      const markdownContent = `# Plugin Tables\n\n${tablesWithIndent}\n`;
      await fs.mkdir(path.dirname(markdownFile), { recursive: true });
      await fs.writeFile(markdownFile, markdownContent);
    })
  );
};

const main = async () => {
  await genTablesJSONFiles();
  await createTablesMarkdowns();
};

main();
