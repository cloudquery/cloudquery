import test, { expect } from "@playwright/test";
import fs from "node:fs";
import YAML from "yaml";

test("Submit the form", async ({ page }) => {
  await page.goto("/");

  // fill the form

  await page.getByRole("button", { name: "Submit" }).click();
  const valuesText = await page
    .locator("text=Values:")
    .locator("xpath=following-sibling::*[1]")
    .textContent();

  expect(valuesText).toBeTruthy();

  if (process.env.E2E_TESTS_GENERATE_CONFIG === "true") {
    const spec = JSON.parse(valuesText as string);
    const localConfig = YAML.stringify({
      kind: "{plugin_kind}",
      spec: {
        name: "{plugin_name}",
        registry: "local",
        path: "../{plugin_name}",
        destinations: ["{plugin_name}"],
        spec: spec.spec,

        // use for destination
        // write_mode: spec.writeMode,
        // migrate_mode: spec.migrateMode,

        // use for source
        // destinations: ['postgresql'],
        // tables: spec.tables,
        // skip_tables: spec.skipTables,
      },
    });

    const anotherConfig = YAML.stringify({
      kind: "{source | destination}", // should be opposite to localConfig
      spec: {
        name: "postgresql",
        path: "cloudquery/postgresql",
        registry: "cloudquery",
        version: "{v6.2.5 | v8.2.7}", // use v6.2.5 for source or v8.2.7 for destination
        spec: {
          connection_string: "test",
        },

        // use for source
        // destinations: ['postgresql'],
        // tables: ['*'],
      },
    });

    if (!fs.existsSync("temp")) {
      fs.mkdirSync("temp");
    }

    fs.writeFileSync(
      "./temp/config.yml",
      `${localConfig}---\n${anotherConfig}`
    );

    fs.writeFileSync(
      "./temp/.env",
      `${spec.envs
        .map(
          (env: { name: string; value: string }) => `${env.name}=${env.value}`
        )
        .join("\n")}`
    );
  }
});
