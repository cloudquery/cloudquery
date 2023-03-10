import React from "react";
import { AuroralabsLogo } from "../logos/Auroralabs";
import { AutodeskLogo } from "../logos/Autodesk";
import { FastlyLogo } from "../logos/Fastly";
import { InfosysLogo } from "../logos/Infosys";
import { InfrastructureLogo } from "../logos/Infrastructure";
import { PaloAltoNetworksLogo } from "../logos/PaloAltoNetworks";
import { ZendeskLogo } from "../logos/Zendesk";

const LogosList = () => {
  const wrapperClassnames = "flex justify-center items-center gap-8 wrapper"
  return (
    <div className={wrapperClassnames}>
      <FastlyLogo />
      <AutodeskLogo />
      <PaloAltoNetworksLogo />
      <InfrastructureLogo />
      <ZendeskLogo />
      <InfosysLogo />
      <AuroralabsLogo />
    </div>
  )
}

export function LogosBlock() {
  return (
    <div className="flex mx-auto 2xl:max-w-7xl overflow-hidden">
      <LogosList />
      <LogosList />
    </div>
  );
}
