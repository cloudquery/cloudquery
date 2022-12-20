
export const getSlackAppLink = (type, name) => {
    const manifest = require(`../data/slack/manifest.json`);
    const encoded = encodeURIComponent(JSON.stringify(manifest, null, 2));
    return 'https://api.slack.com/apps?new_app=1&manifest_json=' + encoded;
};
