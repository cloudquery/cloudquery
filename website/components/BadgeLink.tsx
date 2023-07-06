export const BadgeLink = ({
                          text,
                          href,
                      }: {
    text: string;
    href: string;
    color?: string;
}) => {
    const className = `bg-green-100 text-blue-800 text-xs font-semibold mr-2 px-2.5 py-0.5 rounded dark:bg-blue-200 dark:text-blue-800`;
    return  <a href={href} className={className} target="_blank" rel="noreferrer">
            <span className={className}>{text}</span></a>;

}
