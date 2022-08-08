import React from "react";
import { AutodeskLogo } from "../logos/Autodesk";
import { FastlyLogo } from "../logos/Fastly";
import { TempusLogo } from "../logos/Tempus";
import { ZendeskLogo } from "../logos/Zendesk";

export function LogosBlock() {
  return (
    < div className="flex justify-center items-center flex-wrap gap-4" >
      <FastlyLogo />
      <AutodeskLogo />
      <TempusLogo />
      <ZendeskLogo />
    </div >
  );
}
