import AWSLogo from "./logos/integrations/AWSLogo";
import AzureLogo from "./logos/integrations/AzureLogo";
import CloudflareLogo from "./logos/integrations/CloudflareLogo";
import DigitalOceanLogo from "./logos/integrations/DigitalOceanLogo";
import GCPLogo from "./logos/integrations/GCPLogo";
import KubernetesLogo from "./logos/integrations/KubernetesLogo";
import OktaLogo from "./logos/integrations/OktaLogo";
import TerraformLogo from "./logos/integrations/TerraformLogo";
import YandexCloudLogo from "./logos/integrations/YandexCloudLogo";

const logoClasses =
  "text-gray-600 hover:text-gray-900 dark:text-gray-300 dark:hover:text-gray-50 transition-colors ease-in-out transition hover:scale-105";

export const INTEGRATIONS = [
  {
    name: "Amazon Web Services",
    logo: <AWSLogo className={logoClasses} />,
    link: "https://hub.cloudquery.io/providers/cloudquery/aws/latest",
  },
  {
    name: "Microsoft Azure",
    logo: <AzureLogo className={logoClasses} />,
    link: "https://hub.cloudquery.io/providers/cloudquery/azure/latest",
  },
  {
    name: "Cloudflare",
    logo: <CloudflareLogo className={logoClasses} />,
    link: "https://hub.cloudquery.io/providers/cloudquery/cloudflare/latest",
  },
  {
    name: "Digital Ocean",
    logo: <DigitalOceanLogo className={logoClasses} />,
    link: "https://hub.cloudquery.io/providers/cloudquery/digitalocean/latest",
  },
  {
    name: "Google Cloud Platform",
    logo: <GCPLogo className={logoClasses} />,
    link: "https://hub.cloudquery.io/providers/cloudquery/gcp/latest",
  },
  {
    name: "Kubernetes",
    logo: <KubernetesLogo className={logoClasses} />,
    link: "https://hub.cloudquery.io/providers/cloudquery/k8s/latest",
  },
  {
    name: "Okta",
    logo: <OktaLogo className={logoClasses} />,
    link: "https://hub.cloudquery.io/providers/cloudquery/okta/latest",
  },
  {
    name: "Terraform",
    logo: <TerraformLogo className={logoClasses} />,
    link: "https://hub.cloudquery.io/providers/cloudquery/terraform/latest",
  },
  {
    name: "Yandex Cloud",
    logo: <YandexCloudLogo className={logoClasses} />,
    link: "https://hub.cloudquery.io/providers/yandex-cloud/yandex/latest",
  },
];
