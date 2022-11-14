const osList = ["macOS", "Linux", "Windows"];

const getOSIndex = () => {
  if (typeof window === "undefined") {
    return 0;
  }
  const params = new URLSearchParams(window.location.search);
  const os = params.get("os");
  if (os) {
    const index = osList.indexOf(os);
    return index;
  }
  return 0;
};

export const OSes = {
  options: osList,
  index: () => getOSIndex(),
};
