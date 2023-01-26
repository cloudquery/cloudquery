import { Tab as NextraTab, Tabs as NextraTabs } from "nextra-theme-docs";
import React from "react";

export const Tab = NextraTab;

interface Props {
  options: Array<string>;
  children: React.ReactNode;
}

export function Tabs({ options, children }: Props) {
  const items = options.map((value) => ({ label: value })) as any;
  return (
    <NextraTabs items={items}>
      <div className="mb-4" />
      {children}
    </NextraTabs>
  );
}
