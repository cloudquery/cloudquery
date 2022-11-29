import { promises as fs } from "fs";
import algoliasearch from "algoliasearch";
import { dirname } from "path";
import { fileURLToPath } from "url";

const __dirname = dirname(fileURLToPath(import.meta.url));

const index = async () => {
  try {
    const client = algoliasearch(
      "0OUG5EUZ6H",
      process.env.CQ_ALGOLIA_API_ADMIN_KEY // admin key is required to replace all objects
    );

    const tables = JSON.parse(
      await fs.readFile(`${__dirname}/tables.json`, "utf8")
    );
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

index();
