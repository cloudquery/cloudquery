export const getLatestVersion = (type, name) => {
  const {
    latest: version,
  } = require(`../versions/${type}-${name}.json`);
  return version.split("-")[3];
};
