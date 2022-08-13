import AWSLogo from "./logos/integrations/AWSLogo";
import AzureLogo from "./logos/integrations/AzureLogo";
import CloudflareLogo from "./logos/integrations/CloudflareLogo";
import DigitalOceanLogo from "./logos/integrations/DigitalOceanLogo";
import GCPLogo from "./logos/integrations/GCPLogo";
import KubernetesLogo from "./logos/integrations/KubernetesLogo";
import OktaLogo from "./logos/integrations/OktaLogo";
import TerraformLogo from "./logos/integrations/TerraformLogo";
import YandexCloudLogo from "./logos/integrations/YandexCloudLogo";

const LogoContainer: React.FC<{title: string}> = ({children, title}) => {
    return (<div title={title}>{children}</div>)
}

export function Integrations() {
    return (
        <div className="flex justify-center items-center flex-wrap gap-9 mt-8 sm:mt-4">
            <LogoContainer title="Amazon Web Services">
                <AWSLogo/>
            </LogoContainer>
            <LogoContainer title="Microsoft Azure">
                <AzureLogo />
            </LogoContainer>
            <LogoContainer title="Cloudflare">
                <CloudflareLogo />
            </LogoContainer>
            <LogoContainer title="Digital Ocean">
                <DigitalOceanLogo />
            </LogoContainer>
            <LogoContainer title="Google Cloud Platform">
                <GCPLogo />
            </LogoContainer>
            <LogoContainer title="Kubernetes">
                <KubernetesLogo />
            </LogoContainer>
            <LogoContainer title="Okta">
                <OktaLogo />
            </LogoContainer>
            <LogoContainer title="Terraform">
                <TerraformLogo />
            </LogoContainer>
            <LogoContainer title="Yandex Cloud">
                <YandexCloudLogo />
            </LogoContainer>
        </div>
    )
}