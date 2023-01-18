import { ComponentProps } from "react";
import {
  ArrowsExpandIcon,
  BeakerIcon,
  ChipIcon,
  DatabaseIcon,
  LightningBoltIcon,
  CodeIcon,
  ChartBarIcon,
  KeyIcon,
} from "@heroicons/react/outline";

export type Feature = {
  name: string;
  description: string;
  Icon: (props: ComponentProps<"svg">) => JSX.Element;
  page: "all" | "home" | "docs";
};

export type Features = Array<Feature>;

const FEATURES: Features = [
  {
    name: "Open source",
    description: `Extensible plugin architecture. Contribute to our official plugins or develop your own with CloudQuery SDK.`,
    Icon: CodeIcon,
    page: "all",
  },
  {
    name: "Blazing fast",
    description: `CloudQuery is optimized for performance, utilizing the excellent Go concurrency model with light-weight goroutines.`,
    Icon: LightningBoltIcon,
    page: "all",
  },
  {
    name: "Deploy anywhere",
    description: `CloudQuery plugins are single-binary executables and can be deployed and run anywhere.`,
    Icon: ChartBarIcon,
    page: "all",
  },
  {
    name: "Pre-built queries",
    description: `CloudQuery maintains a number of out-of-the-box security and compliance policies for cloud infrastructure.`,
    Icon: KeyIcon,
    page: "all",
  },
  {
    name: "Eliminate data silos",
    description: `Eliminate data silos across your organization, unifying data between security, infrastructure, marketing and finance teams.`,
    Icon: DatabaseIcon,
    page: "all",
  },
  {
    name: "Unlimited scale",
    description: `CloudQuery plugins are stateless and can scaled easily horizontally on any executor such as k8s, batch jobs.`,
    Icon: ArrowsExpandIcon,
    page: "all",
  },
];

export const DOCS_FEATURES = FEATURES.filter(
  (f) => f.page === "docs" || f.page === "all"
);

export const HOME_FEATURES = FEATURES.filter(
  (f) => f.page === "home" || f.page === "all"
);

export default FEATURES;
