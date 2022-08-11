import React from "react";
import { AuroralabsLogo } from "../logos/Auroralabs";
import { AutodeskLogo } from "../logos/Autodesk";
import { BloombergLogo } from "../logos/Bloomberg";
import { FastlyLogo } from "../logos/Fastly";
import { InfosysLogo } from "../logos/Infosys";
import { InfrastructureLogo } from "../logos/Infrastructure";
import { PaloAltoNetworksLogo } from "../logos/PaloAltoNetworks";
import { TempusLogo } from "../logos/Tempus";
import { ZendeskLogo } from "../logos/Zendesk";

export function LogosBlock() {
  return (
    <div className="flex justify-center items-center flex-wrap gap-8">
      <FastlyLogo />
      <AutodeskLogo />
      <PaloAltoNetworksLogo />
      <BloombergLogo />
      <InfrastructureLogo />
      <TempusLogo />
      <ZendeskLogo />
      <InfosysLogo />
      <AuroralabsLogo />
    </div>
  );
}
