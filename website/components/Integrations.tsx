import AWSLogo from "./logos/integrations/AWSLogo";
import AzureLogo from "./logos/integrations/AzureLogo";
import CloudflareLogo from "./logos/integrations/CloudflareLogo";
import DigitalOceanLogo from "./logos/integrations/DigitalOceanLogo";
import GCPLogo from "./logos/integrations/GCPLogo";
import KubernetesLogo from "./logos/integrations/KubernetesLogo";
import OktaLogo from "./logos/integrations/OktaLogo";
import TerraformLogo from "./logos/integrations/TerraformLogo";
import YandexCloudLogo from "./logos/integrations/YandexCloudLogo";

export function Integrations() {
    return (
        <div className="flex justify-center items-center flex-wrap gap-8">
            <AWSLogo />
            <AzureLogo />
            <CloudflareLogo />
            <DigitalOceanLogo />
            <GCPLogo />
            <KubernetesLogo />
            <OktaLogo />
            <TerraformLogo />
            <YandexCloudLogo />
        </div>
    )
}