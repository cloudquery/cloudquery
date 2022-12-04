export const Badge = ({
                          text,
                      }: {
    text: string;
    color?: string;
}) => {
    const className = `bg-blue-100 text-blue-800 text-xs font-semibold mr-2 px-2.5 py-0.5 rounded dark:bg-blue-200 dark:text-blue-800`;
    return <span
        className={className}>{text}</span>;

}
