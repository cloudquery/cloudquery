export const getLatestVersion = (type, name) => {
  const {
    latest: version,
  } = require(`../../sites/versions/v2/${type}-${name}.json`);
  return version.split("-")[3];
};
