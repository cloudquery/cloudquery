const fs = require("fs");
const path = require("path");

// Read the plugin data file
import {SOURCE_PLUGINS, DESTINATION_PLUGINS} from "../components/pluginData";

// Define the directories to write the MDX files to
const outputDir = "./pages/integrations";
const mdxSourceComponentDir = "./components/mdx/plugins/source";
const mdxDestinationComponentDir = "./components/mdx/plugins/destination";

// Loop through each source plugin and generate or copy MDX files
SOURCE_PLUGINS.forEach((plugin) => {
  const sourceDir = `./pages/docs/plugins/sources/${plugin.id}`;

  // Copy the authentication and configuration files if they exist
  const authFilePath = path.join(sourceDir, "_authentication.mdx");
  const configFilePath = path.join(sourceDir, "_configuration.mdx");

  if (!fs.existsSync(mdxSourceComponentDir + "/" + plugin.id)) {
    fs.mkdirSync(mdxSourceComponentDir+ "/" + plugin.id, { recursive: true });
  }

  if (fs.existsSync(authFilePath)) {
    const outputFilePath = path.join(mdxSourceComponentDir, `${plugin.id}/_authentication.mdx`);
    fs.copyFileSync(authFilePath, outputFilePath);
  }

  if (fs.existsSync(configFilePath)) {
    DESTINATION_PLUGINS.forEach((destination) => {
        const sourceConfigDir = mdxSourceComponentDir + `/${plugin.id}/${destination.id}`;
        if (!fs.existsSync(sourceConfigDir)) {
            fs.mkdirSync(sourceConfigDir, { recursive: true });
        }
        let fileContents = fs.readFileSync(configFilePath, "utf8");
        fileContents = fileContents.replace(/DESTINATION_NAME/g, destination.id);
        const outputFilePath = path.join(sourceConfigDir, `_configuration.mdx`);
        fs.writeFileSync(outputFilePath, fileContents);
    })
  }

  // Define the file path for the new MDX file
  const filePath = path.join(outputDir, `${plugin.id}.mdx`);

  // Define the contents of the MDX file
  const fileContents = `---
title: ${plugin.name} Data Integration
---

import Integration from "../../components/pages/IntegrationSource";
import {getPlugin} from "../../components/pluginData";

<Integration
  source={getPlugin("source", "${plugin.id}")}
/>`;

  // Write the contents to the new file
  fs.writeFileSync(filePath, fileContents);
});

// Loop through each destination plugin and generate or copy MDX files
DESTINATION_PLUGINS.forEach((plugin) => {
  const destinationDir = `./pages/docs/plugins/destinations/${plugin.id}`;

  // Copy the authentication and configuration files if they exist
  const authFilePath = path.join(destinationDir, "_authentication.mdx");
  const configFilePath = path.join(destinationDir, "_configuration.mdx");

  if (!fs.existsSync(mdxDestinationComponentDir + "/" + plugin.id)) {
    fs.mkdirSync(mdxDestinationComponentDir+ "/" + plugin.id, { recursive: true });
  }

  if (fs.existsSync(authFilePath)) {
    const outputFilePath = path.join(mdxDestinationComponentDir, `${plugin.id}/_authentication.mdx`);
    fs.copyFileSync(authFilePath, outputFilePath);
  }

  if (fs.existsSync(configFilePath)) {
    const outputFilePath = path.join(mdxDestinationComponentDir, `${plugin.id}/_configuration.mdx`);
    fs.copyFileSync(configFilePath, outputFilePath);
  }
});

const metaFileContents = `{
  "*": {
    "theme": {
      "sidebar": false,
      "breadcrumb": true,
      "typesetting": "page",
      "toc": false,
      "footer": true,
      "pagination": false
    }
  }
}
`;

// Write _meta.json for sources
const metaFilePath = path.join(
    outputDir,
    `_meta.json`
);
fs.writeFileSync(metaFilePath, metaFileContents);

// Loop through each source-destination combination and generate integration page MDX files
SOURCE_PLUGINS.forEach((source) => {
  // Create the output directory if it doesn't exist
  if (!fs.existsSync(outputDir + "/" + source.id)) {
    fs.mkdirSync(outputDir+ "/" + source.id, { recursive: true });
  }

  DESTINATION_PLUGINS.forEach((destination) => {
    // Define the file path for the new MDX file
    const filePath = path.join(
      outputDir,
      `${source.id}/${destination.id}.mdx`
    );

    // Define the contents of the MDX file
    const fileContents = `---
title: Export data from ${source.name} to ${destination.name}
---

import Integration from "../../../components/pages/IntegrationSourceDestination";
import {getPlugin} from "../../../components/pluginData";
import SourceConfiguration from "../../../components/mdx/plugins/source/${source.id}/${destination.id}/_configuration.mdx";
import SourceAuthentication from "../../../components/mdx/plugins/source/${source.id}/_authentication.mdx";
import DestinationConfiguration from "../../../components/mdx/plugins/destination/${destination.id}/_configuration.mdx";
import SyncCommand from "../../../components/mdx/plugins/source/${source.id}/${destination.id}/_sync.mdx";

<Integration
  source={getPlugin("source", "${source.id}")}
  sourceConfiguration={<SourceConfiguration />}
  sourceAuthentication={<SourceAuthentication />}
  destination={getPlugin("destination", "${destination.id}")}
  destinationConfiguration={<DestinationConfiguration />}
  syncCommand={<SyncCommand />}
>
</Integration>
`;

    // Write the contents to the new file
    fs.writeFileSync(filePath, fileContents);

    // Write _meta.json
    const metaFilePath = path.join(
      outputDir,
      `${source.id}/_meta.json`
    );
    fs.writeFileSync(metaFilePath, metaFileContents);

      // Prepare the sync command file directory
    const syncCommandDir = mdxSourceComponentDir + `/${source.id}/${destination.id}`;
    if (!fs.existsSync(syncCommandDir)) {
        fs.mkdirSync(syncCommandDir, { recursive: true });
    }
    // Write the sync command file
    const syncCommandFilePath = path.join(syncCommandDir, "_sync.mdx");
    const syncCommandFileContents = "```bash copy\n" +
        `cloudquery sync ${source.id}.yaml ${destination.id}.yaml\n` +
        "```"
    // Write the contents to the new file
    fs.writeFileSync(syncCommandFilePath, syncCommandFileContents);
})});

console.log("MDX files generated successfully!");