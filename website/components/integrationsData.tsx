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

export const INTEGRATIONS = [
  {
    name: "Amazon Web Services",
    logo: <AWSLogo />,
    id: "aws",
  },
  {
    name: "Microsoft Azure",
    logo: <AzureLogo />,
    id: "azure",
  },
  {
    name: "Cloudflare",
    logo: <CloudflareLogo />,
    id: "cloudflare",
  },
  {
    name: "Digital Ocean",
    logo: <DigitalOceanLogo />,
    id: "digitalocean",
  },
  {
    name: "Google Cloud Platform",
    logo: <GCPLogo />,
    id: "gcp",
  },
  {
    name: "Github",
    logo: <GithubLogo />,
    id: "github",
  },
  {
    name: "Kubernetes",
    logo: <KubernetesLogo />,
    id: "k8s",
  },
  {
    name: "Okta",
    logo: <OktaLogo />,
    id: "okta",
  },
  {
    name: "Terraform",
    logo: <TerraformLogo />,
    id: "terraform",
  },
  {
    name: "Yandex Cloud",
    logo: <YandexCloudLogo />,
    id: "yandexcloud",
    href: "https://github.com/yandex-cloud/cq-source-yandex",
  },
];
