import * as React from "react";

export const CustomErrorPage = () => {
  React.useEffect(() => {
    const data = { path: window.location.pathname };
    window.sa_event("404", data);
  }, []);
  return <></>;
};
