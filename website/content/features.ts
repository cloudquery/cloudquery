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
    name: "Cloud asset inventory",
    description: `Build your own multi-cloud asset inventory with standard SQL and BI tools.`,
    Icon: ChartBarIcon,
    page: "all",
  },
  {
    name: "CSPM",
    description: `Customize pre-built open source SQL policies and visualize them with your any of your favorite BI tools.`,
    Icon: KeyIcon,
    page: "all",
  },
  {
    name: "Open source",
    description: `Easily extensible plugin architecture. Contribute to our official plugins or develop your own with CloudQuery SDK.`,
    Icon: CodeIcon,
    page: "all",
  },
  {
    name: "Blazing fast",
    description: `CloudQuery concurrency system utilizes the excellent Go concurrency model with light-weight goroutines.`,
    Icon: LightningBoltIcon,
    page: "all",
  },
  {
    name: "Database agnostic",
    description: `CloudQuery can store your configuration in any supported destination such as database, data lake or streaming platform for further analysis.`,
    Icon: DatabaseIcon,
    page: "all",
  },
  {
    name: "Raw access to data",
    description: `Decouple data ingestion and get raw access to your data in structured and unstructured formats.`,
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
