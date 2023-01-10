export const getLatestVersion = (type, name) => {
  try {
    const { latest: version } = require(`../versions/${type}-${name}.json`);
    return version.split("-")[3];
  } catch (e) {
    return "Unpublished";
  }
};
