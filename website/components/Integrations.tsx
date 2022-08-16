import { INTEGRATIONS } from "./integrationsData"

const LogoContainer: React.FC<{ title: string, href: string }> = ({ children, title, href = '/' }) => {
    return (
        <a href={href} title={title} className='w-9 h-9 flex items-center justify-center'>
            {children}
        </a>
    )
}

export function Integrations() {
    return (
        <div className="flex justify-center items-center flex-wrap gap-9 mt-8 sm:mt-4">
            {
                INTEGRATIONS.map(({ name, logo, link }) => (
                    <LogoContainer title={name} href={link}>
                        {logo}
                    </LogoContainer>
                ))
            }
        </div>
    )
}