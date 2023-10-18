import {locatePathSync} from 'locate-path';
import {ALL_PREMIUM_POLICIES, Policy} from "../components/policyData";
import fs from "fs";
import path from "path";

// Read the plugin data file
import {
    Plugin,
    ALL_SOURCE_PLUGINS, ALL_DESTINATION_PLUGINS, ALL_PLUGINS, PUBLISHED_SOURCE_PLUGINS, PUBLISHED_DESTINATION_PLUGINS
} from "../components/pluginData";

// Define the directories to write the MDX files to
const outputDir = "./integrations";
const mdxSourceComponentDir = "./components/mdx/plugins/source";
const mdxDestinationComponentDir = "./components/mdx/plugins/destination";

function getPluginRedirectContent(plugin: Plugin, licenseType: string, licenseName: string) {
   return `---
title: Buy ${plugin.name} (${licenseName})
---

import Head from "next/head";

<Head>
  <meta httpEquiv="refresh" content="5; url='${plugin.buyLinks[licenseType]}'" />
</Head>

## Purchase ${plugin.name} (${licenseName}${(plugin.availability === "unpublished") ? " - Pre-order" : ""})

You will be redirected to a Stripe checkout page to complete your purchase in 5 seconds…

If the page does not redirect automatically, please click this link: [${plugin.buyLinks[licenseType]}](${plugin.buyLinks[licenseType]})
`;
}

function createPluginBuyRedirects() {
    const buyDir = `./pages/buy`;
    ALL_PLUGINS.forEach((plugin) => {
        if (plugin.buyLinks && plugin.buyLinks['standard']) {
            const filePath = path.join(buyDir, `${plugin.id}-standard.mdx`);
            fs.writeFileSync(filePath, getPluginRedirectContent(plugin, 'standard', "Standard License"));
        }
        if (plugin.buyLinks && plugin.buyLinks['extended']) {
            const filePath = path.join(buyDir, `${plugin.id}-extended.mdx`);
            fs.writeFileSync(filePath, getPluginRedirectContent(plugin, 'extended', "Extended License"));
        }
    });
}


function getPolicyRedirectContent(policy: Policy, licenseType: string, licenseName: string) {
    return `---
title: Buy ${policy.name} (${licenseName})
---

import Head from "next/head";

<Head>
  <meta httpEquiv="refresh" content="5; url='${policy.buyLinks[licenseType]}'" />
</Head>

## Purchase ${policy.name} (${licenseName})

You will be redirected to a Stripe checkout page to complete your purchase in 5 seconds…

If the page does not redirect automatically, please click this link: [${policy.buyLinks[licenseType]}](${policy.buyLinks[licenseType]})
`;
}

function getPolicySignupContent(policy: Policy, licenseName: string) {
    return `---
title: Buy ${policy.name} (${licenseName})
---

## Purchase ${policy.name} (${licenseName})

This policy is currently only accessible through our early access program. [Get in touch](/contact-policies).
`;
}

function createPolicyBuyRedirects() {
    const buyDir = `./pages/buy`;
    ALL_PREMIUM_POLICIES.forEach((policy) => {
        if (!policy.availableForPurchase) {
            fs.writeFileSync(path.join(buyDir, `${policy.id}-standard.mdx`), getPolicySignupContent(policy, "Standard License"));
            fs.writeFileSync(path.join(buyDir, `${policy.id}-extended.mdx`), getPolicySignupContent(policy, "Extended License"));
            return;
        }
        if (policy.buyLinks && policy.buyLinks['standard']) {
            const filePath = path.join(buyDir, `${policy.id}-standard.mdx`);
            fs.writeFileSync(filePath, getPolicyRedirectContent(policy, 'standard', "Standard License"));
        }
        if (policy.buyLinks && policy.buyLinks['extended']) {
            const filePath = path.join(buyDir, `${policy.id}-extended.mdx`);
            fs.writeFileSync(filePath, getPolicyRedirectContent(policy, 'extended', "Extended License"));
        }
    });
}


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
    const authFilePath = locatePathSync([`${sourceDir}/_authentication.md`]);
    if (authFilePath) {
        const ext = path.extname(authFilePath);
        const outputFilePath = path.join(mdxSourceComponentDir, `${source.id}/_authentication${ext}`);
        fs.copyFileSync(authFilePath, outputFilePath);
        return true;
    }
    return false;
}

// Copy the source configuration file if it exists and replace the destination name
function copySourceConfigurationFile(source: Plugin): boolean {
    const sourceDir = `./pages/docs/plugins/sources/${source.id}`;
    const configFilePath = locatePathSync([`${sourceDir}/_configuration.md`]);
    if (configFilePath) {
        ALL_DESTINATION_PLUGINS.forEach((destination) => {
            const sourceConfigDir = mdxSourceComponentDir + `/${source.id}/${destination.id}`;
            recreateDirectory(sourceConfigDir);
            let fileContents = fs.readFileSync(configFilePath, "utf8");
            fileContents = fileContents.replace(/DESTINATION_NAME/g, destination.id);
            const ext = path.extname(configFilePath);
            const outputFilePath = path.join(sourceConfigDir, `_configuration${ext}`);
            fs.writeFileSync(outputFilePath, fileContents);
        })
        return true;
    }
    return false;
}

// Copy the destination authentication file if it exists
function copyDestinationAuthenticationFile(destination: Plugin) : boolean {
    const destinationDir = `./pages/docs/plugins/destinations/${destination.id}`;
    // Copy the authentication and configuration files if they exist
    const authFilePath = locatePathSync([`${destinationDir}/_authentication.md`]);
    if (authFilePath) {
        const ext = path.extname(authFilePath);
        const outputFilePath = path.join(mdxDestinationComponentDir, `${destination.id}/_authentication${ext}`);
        fs.copyFileSync(authFilePath, outputFilePath);
        return true;
    }
    return false;
}

// Copy the destination configuration file if it exists
function copyDestinationConfigurationFile(destination: Plugin) : boolean {
    const destinationDir = `./pages/docs/plugins/destinations/${destination.id}`;
    const configFilePath = locatePathSync([`${destinationDir}/_configuration.md`]);
    if (configFilePath) {
        const ext = path.extname(configFilePath);
        const outputFilePath = path.join(mdxDestinationComponentDir, `${destination.id}/_configuration${ext}`);
        fs.copyFileSync(configFilePath, outputFilePath);
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

<IntegrationSource source="${source.id}" />`;

    // Write the contents to the new file
    fs.writeFileSync(filePath, fileContents);
}

function createSourceDestinationIntegrationFile(source: Plugin, destination: Plugin, sourceHasAuth: boolean, destHasAuth: boolean) {
    // Define the file path for the new MDX file
    const filePath = path.join(
        outputDir,
        `${source.id}/${destination.id}.mdx`
    );

    const isOfficialSource = source.availability === "free";
    const isOfficialDestination = destination.availability === "free";

    // Define the contents of the MDX file
    const fileContents = `---
title: Export data from ${source.name} to ${destination.name}
---

<IntegrationDestination 
    source="${source.id}" 
    destination="${destination.id}" 
    isOfficialSource={${isOfficialSource}} 
    isOfficialDestination={${isOfficialDestination}}
    sourceHasAuth={${sourceHasAuth}}
    destHasAuth={${destHasAuth}}
/>`;

    // Write the contents to the new file
    fs.writeFileSync(filePath, fileContents);

    // Prepare the sync command file directory
    const syncCommandDir = mdxSourceComponentDir + `/${source.id}/${destination.id}`;
    if (!fs.existsSync(syncCommandDir)) {
        fs.mkdirSync(syncCommandDir, { recursive: true });
    }
    // Write the sync command file
    const syncCommandFilePath = path.join(syncCommandDir, "_sync.md");
    const sourceFilename = source.id === destination.id ? `source-${source.id}.yaml` : `${source.id}.yaml`;
    const destinationFilename = source.id === destination.id ? `destination-${destination.id}.yaml` : `${destination.id}.yaml`;
    const syncCommandFileContents = "```bash copy\n" +
        `cloudquery sync ${sourceFilename} ${destinationFilename}\n` +
        "```"
    // Write the contents to the new file
    fs.writeFileSync(syncCommandFilePath, syncCommandFileContents);
}

function generateFiles() {
    const sources = new Set<string>();
    const destinations = new Set<string>();

    let hasAuthFile = {};

    PUBLISHED_SOURCE_PLUGINS.forEach((source) => {
        if (source.availability === "premium") {
            return;
        }
        recreateDirectory(outputDir + "/" + source.id);
    });

    // Loop through each source plugin and generate or copy MDX files
    ALL_SOURCE_PLUGINS.forEach((source) => {
      if (sources.has(source.id)) {
        throw new Error("Duplicate source id: " + source.id + ". Did you forget to remove an unpublished plugin you implemented?");
      }
      sources.add(source.id);

      const hasConfiguration = copySourceConfigurationFile(source);
      const isOfficial = source.availability === "free";
      if (isOfficial && !hasConfiguration) {
          throw new Error("No _configuration.md file found for source: " + source.id);
      }
      const hasAuthentication = copySourceAuthenticationFile(source);
      hasAuthFile['source-' + source.id] = hasAuthentication;
      createSourceIntegrationFile(source);
    });

    // Loop through each destination plugin and generate or copy MDX files
    ALL_DESTINATION_PLUGINS.forEach((destination) => {
        if (destination.availability === "premium") {
            return;
        }
        if (destinations.has(destination.id)) {
            throw new Error("Duplicate destination id: " + destination.id);
        }
        destinations.add(destination.id);
        recreateDirectory(mdxDestinationComponentDir + "/" + destination.id);

        const hasConfiguration = copyDestinationConfigurationFile(destination);
        const isOfficial = destination.availability === "free";
        if (isOfficial && !hasConfiguration && destination.id !== "more") {
            throw new Error("No _configuration.md file found for destination: " + destination.id);
        }
        const hasAuthentication = copyDestinationAuthenticationFile(destination);
        hasAuthFile['destination-' + destination.id] = hasAuthentication;
    });

    // Create the source -> destination integration files
    PUBLISHED_SOURCE_PLUGINS.forEach((source: Plugin) => {
        if (source.availability === "premium") {
            return;
        }
        PUBLISHED_DESTINATION_PLUGINS.forEach((destination: Plugin) => {
           if (destination.availability === "premium") {
             return;
           }
           const sourceHasAuth = hasAuthFile['source-' + source.id];
           const destHasAuth = hasAuthFile['destination-' + destination.id];
           createSourceDestinationIntegrationFile(source, destination, sourceHasAuth, destHasAuth);
       });
    });
}


generateFiles()
createPluginBuyRedirects()
createPolicyBuyRedirects()

console.log("MDX files generated successfully!");