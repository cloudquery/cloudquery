const fs = require("fs");
const path = require("path");

// Read the plugin data file
import {Plugin, SOURCE_PLUGINS, DESTINATION_PLUGINS} from "../components/pluginData";

// Define the directories to write the MDX files to
const outputDir = "./pages/integrations";
const mdxSourceComponentDir = "./components/mdx/plugins/source";
const mdxDestinationComponentDir = "./components/mdx/plugins/destination";

function recreateDirectory(dir: string) {
    if (fs.existsSync(dir)) {
        // Clear the directory if it exists
        fs.rmSync(dir, { recursive: true, force: true });
        fs.mkdirSync(dir);
    } else {
        // Create the directory if it doesn't exist
        fs.mkdirSync(dir, { recursive: true });
    }
}

// Copy the source authentication file if it exists
function copySourceAuthenticationFile(source: Plugin) : boolean {
    const sourceDir = `./pages/docs/plugins/sources/${source.id}`;

    // Copy the authentication and configuration files if they exist
    const authFilePath = path.join(sourceDir, "_authentication.mdx");

    if (fs.existsSync(authFilePath)) {
        const outputFilePath = path.join(mdxSourceComponentDir, `${source.id}/_authentication.mdx`);
        fs.copyFileSync(authFilePath, outputFilePath);
        return true;
    }
    return false;
}

// Copy the source configuration file if it exists and replace the destination name
function copySourceConfigurationFile(source: Plugin): boolean {
    const configFilePath = `./pages/docs/plugins/sources/${source.id}/_configuration.mdx`;
    if (fs.existsSync(configFilePath)) {
        DESTINATION_PLUGINS.forEach((destination) => {
            const sourceConfigDir = mdxSourceComponentDir + `/${source.id}/${destination.id}`;
            recreateDirectory(sourceConfigDir);
            let fileContents = fs.readFileSync(configFilePath, "utf8");
            fileContents = fileContents.replace(/DESTINATION_NAME/g, destination.id);
            const outputFilePath = path.join(sourceConfigDir, `_configuration.mdx`);
            fs.writeFileSync(outputFilePath, fileContents);
        })
        return true;
    }
    return false;
}

// Copy the destination authentication file if it exists
function copyDestinationAuthenticationFile(destination: Plugin) : boolean {
    const destinationPluginDir = `./pages/docs/plugins/destinations/${destination.id}`;

    // Copy the authentication and configuration files if they exist
    const authFilePath = path.join(destinationPluginDir, "_authentication.mdx");
    if (fs.existsSync(authFilePath)) {
        const outputFilePath = path.join(mdxDestinationComponentDir, `${destination.id}/_authentication.mdx`);
        fs.copyFileSync(authFilePath, outputFilePath);
        return true;
    }
    return false;
}

// Copy the destination configuration file if it exists
function copyDestinationConfigurationFile(destination: Plugin) : boolean {
    const destinationDir = `./pages/docs/plugins/destinations/${destination.id}`;
    const authFilePath = path.join(destinationDir, "_configuration.mdx");

    if (fs.existsSync(authFilePath)) {
        const outputFilePath = path.join(mdxDestinationComponentDir, `${destination.id}/_configuration.mdx`);
        fs.copyFileSync(authFilePath, outputFilePath);
        return true;
    }
    return false;
}


function createSourceIntegrationFile(source: Plugin) {
    // Define the file path for the new MDX file
    const filePath = path.join(outputDir, `${source.id}.mdx`);

    // Define the contents of the MDX file
    const fileContents = `---
title: ${source.name} Data Integration
---

import Integration from "../../components/pages/IntegrationSource";
import {getPlugin} from "../../components/pluginData";

<Integration
  source={getPlugin("source", "${source.id}")}
/>`;

    // Write the contents to the new file
    fs.writeFileSync(filePath, fileContents);
}

function createMetaFiles() {
    const metaFileContents = `{
  "*": {
    "theme": {
      "sidebar": false,
      "breadcrumb": true,
      "typesetting": "default",
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

    SOURCE_PLUGINS.forEach((source) => {
        const metaFileSourcePath = path.join(
            outputDir,
            `${source.id}/_meta.json`
        );
        fs.writeFileSync(metaFileSourcePath, metaFileContents);
    });
}

function createSourceDestinationIntegrationFile(source: Plugin, destination: Plugin, sourceHasAuth: boolean, destHasAuth: boolean) {
    // Define the file path for the new MDX file
    const filePath = path.join(
        outputDir,
        `${source.id}/${destination.id}.mdx`
    );

    const isOfficialSource = source.kind === "official";
    const isOfficialDestination = destination.kind === "official";

    // Define the contents of the MDX file
    const fileContents = `---
title: Export data from ${source.name} to ${destination.name}
---

import Integration from "../../../components/pages/IntegrationSourceDestination";
import {getPlugin} from "../../../components/pluginData";
${isOfficialSource ? `import SourceConfiguration from "../../../components/mdx/plugins/source/${source.id}/${destination.id}/_configuration.mdx";` : `` }
${sourceHasAuth ? `import SourceAuthentication from "../../../components/mdx/plugins/source/${source.id}/_authentication.mdx";` : ``} 
${isOfficialDestination ? `import DestinationConfiguration from "../../../components/mdx/plugins/destination/${destination.id}/_configuration.mdx";` : `` }
${destHasAuth ? `import DestinationAuthentication from "../../../components/mdx/plugins/destination/${destination.id}/_authentication.mdx";` : ``}
import SyncCommand from "../../../components/mdx/plugins/source/${source.id}/${destination.id}/_sync.mdx";

<Integration
  source={getPlugin("source", "${source.id}")}
  ${isOfficialSource ? `sourceConfiguration={<SourceConfiguration />}` : ``}
  ${sourceHasAuth ? `sourceAuthentication={<SourceAuthentication />}` : ``}
  destination={getPlugin("destination", "${destination.id}")}
  ${isOfficialDestination ? `destinationConfiguration={<DestinationConfiguration />}` : ``}
  ${destHasAuth ? `destinationAuthentication={<DestinationAuthentication />}` : ``}
  syncCommand={<SyncCommand />}
/>
`;

    // Write the contents to the new file
    fs.writeFileSync(filePath, fileContents);

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
}

function generateFiles() {
    let hasAuthFile = {};

    // Loop through each source plugin and generate or copy MDX files
    SOURCE_PLUGINS.forEach((source) => {
      recreateDirectory(outputDir + "/" + source.id);

      const hasConfiguration = copySourceConfigurationFile(source);
      const isOfficial = source.kind === "official";
      if (isOfficial && !hasConfiguration) {
          throw new Error("No _configuration.mdx file found for source: " + source.id);
      }
      const hasAuthentication = copySourceAuthenticationFile(source);
      hasAuthFile['source-' + source.id] = hasAuthentication;
      createSourceIntegrationFile(source);
    });

    // Loop through each destination plugin and generate or copy MDX files
    DESTINATION_PLUGINS.forEach((destination) => {
        recreateDirectory(mdxDestinationComponentDir + "/" + destination.id);

        const hasConfiguration = copyDestinationConfigurationFile(destination);
        const isOfficial = destination.kind === "official";
        if (isOfficial && !hasConfiguration && destination.id !== "more") {
            throw new Error("No _configuration.mdx file found for destination: " + destination.id);
        }
        const hasAuthentication = copyDestinationAuthenticationFile(destination);
        hasAuthFile['destination-' + destination.id] = hasAuthentication;
    });

    // Create the source -> destination integration files
    SOURCE_PLUGINS.forEach((source: Plugin) => {
       DESTINATION_PLUGINS.forEach((destination: Plugin) => {
           const sourceHasAuth = hasAuthFile['source-' + source.id];
           const destHasAuth = hasAuthFile['destination-' + destination.id];
           createSourceDestinationIntegrationFile(source, destination, sourceHasAuth, destHasAuth);
       });
    });

    createMetaFiles();
}


generateFiles()

console.log("MDX files generated successfully!");