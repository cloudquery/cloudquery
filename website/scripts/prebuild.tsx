const fs = require("fs");
const path = require("path");

// Read the plugin data file
import {SOURCE_PLUGINS, DESTINATION_PLUGINS} from "../components/pluginData";

// Define the directory to write the MDX files to
const outputDir = "./pages/integrations";

// Loop through each plugin and generate an MDX file
SOURCE_PLUGINS.forEach((plugin) => {
  // Define the file path for the new MDX file
  const filePath = path.join(outputDir, `${plugin.id}.mdx`);

  // Define the contents of the MDX file
  const fileContents = `---
title: ${plugin.name} Integration
---

import Integration from "../../components/pages/IntegrationSource";
import {getPlugin} from "../../components/pluginData";

<Integration
  source={getPlugin("source", "${plugin.id}")}
/>`;

  // Write the contents to the new file
  fs.writeFileSync(filePath, fileContents);
});

// Loop through each source-destination combination and generate an MDX file
SOURCE_PLUGINS.forEach((source) => {
  // Create the output directory if it doesn't exist
  if (!fs.existsSync(outputDir + "/" + source.id)) {
    fs.mkdirSync(outputDir+ "/" + source.id);
  }

  DESTINATION_PLUGINS.forEach((destination) => {
    // Define the file path for the new MDX file
    const filePath = path.join(
      outputDir,
      `${source.id}/${destination.id}.mdx`
    );

    // Define the contents of the MDX file
    const fileContents = `---
title: AWS to Postgres Integration
---

import Integration from "../../../components/pages/IntegrationSourceDestination";
import {getPlugin} from "../../../components/pluginData";
import SourceConfiguration from "../../../components/mdx/plugins/source/${source.id}/_configuration.mdx";
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
})});

console.log("MDX files generated successfully!");