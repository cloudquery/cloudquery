import AWSLogo from "./logos/integrations/AWSLogo";
import AzureLogo from "./logos/integrations/AzureLogo";
import CloudflareLogo from "./logos/integrations/CloudflareLogo";
import DigitalOceanLogo from "./logos/integrations/DigitalOceanLogo";
import GCPLogo from "./logos/integrations/GCPLogo";
import GithubLogo from "./logos/integrations/GithubLogo";
import KubernetesLogo from "./logos/integrations/KubernetesLogo";
import OktaLogo from "./logos/integrations/OktaLogo";
import TerraformLogo from "./logos/integrations/TerraformLogo";
import YandexCloudLogo from "./logos/integrations/YandexCloudLogo";
import HerokuLogo from "./logos/integrations/HerokuLogo";

const getPluginURL = (id: string) =>
  `https://github.com/cloudquery/cloudquery/tree/${
    process.env.VERCEL_GIT_COMMIT_REF || `main`
  }/plugins/source/${id}/docs`;

export const INTEGRATIONS = [
  {
    name: "Amazon Web Services",
    logo: <AWSLogo />,
    id: "aws",
    href: getPluginURL("aws"),
  },
  {
    name: "Microsoft Azure",
    logo: <AzureLogo />,
    id: "azure",
    href: getPluginURL("azure"),
  },
  {
    name: "Cloudflare",
    logo: <CloudflareLogo />,
    id: "cloudflare",
    href: getPluginURL("cloudflare"),
  },
  {
    name: "Digital Ocean",
    logo: <DigitalOceanLogo />,
    id: "digitalocean",
    href: getPluginURL("digitalocean"),
  },
  {
    name: "Google Cloud Platform",
    logo: <GCPLogo />,
    id: "gcp",
    href: getPluginURL("gcp"),
  },
  {
    name: "GitHub",
    logo: <GithubLogo />,
    id: "github",
    href: getPluginURL("github"),
  },
  {
    name: "Heroku",
    logo: <HerokuLogo />,
    id: "heroku",
    href: getPluginURL("heroku"),
  },
  {
    name: "Kubernetes",
    logo: <KubernetesLogo />,
    id: "k8s",
    href: getPluginURL("k8s"),
  },
  {
    name: "Okta",
    logo: <OktaLogo />,
    id: "okta",
    href: getPluginURL("okta"),
  },
  {
    name: "Terraform",
    logo: <TerraformLogo />,
    id: "terraform",
    href: getPluginURL("terraform"),
  },
  {
    name: "Yandex Cloud",
    logo: <YandexCloudLogo />,
    id: "yandexcloud",
    href: `https://github.com/yandex-cloud/cq-provider-yandex`,
  },
];
