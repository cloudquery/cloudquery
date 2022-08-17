import { DuplicateIcon } from "@heroicons/react/outline"

const SecurityQuery = () => (
    <>
        <div>
            <span style={{ color: "var(--shiki-token-keyword)" }}>SELECT</span>
            <span style={{ color: "var(--shiki-token-keyword)" }}> * </span>
            <span style={{ color: "var(--shiki-token-keyword)" }}>FROM </span>aws_elbv2_load_balancers
        </div>
        <div>
            <span style={{ color: "var(--shiki-token-keyword)" }}>WHERE </span>scheme
            <span style={{ color: "var(--shiki-token-keyword)" }}> = </span>
            <span style={{ color: "var(--shiki-token-string-expression)" }}>'internet-facing'</span>
        </div>
    </>
)

const ComplianceQuery = () => (
    <>
        <div>
            <span style={{ color: "var(--shiki-token-keyword)" }}>SELECT</span> account_id, require_uppercase_characters
        </div>
        <div>
            <span style={{ color: "var(--shiki-token-keyword)" }}>FROM</span> aws_iam_password_policies
        </div>
        <div>
            <span style={{ color: "var(--shiki-token-keyword)" }}>WHERE</span> require_uppercase_characters =
            <span style={{ color: "var(--shiki-token-keyword)" }}> FALSE</span>
        </div>
    </>
)

const QueryAcrossApps = () => (
    <>
        <div>
            <span style={{ color: "var(--shiki-token-keyword)" }}>SELECT</span> arn
            <span style={{ color: "var(--shiki-token-keyword)" }}> FROM</span> aws_iam_users
        </div>
        <div><span style={{ color: "var(--shiki-token-keyword)" }}>JOIN</span> aws_iam_user_tags</div>
        <div><span style={{ color: "var(--shiki-token-keyword)" }}>ON</span> aws_iam_users.id = aws_iam_user_tags.user_id</div>
        <div><span style={{ color: "var(--shiki-token-keyword)" }}>JOIN</span> okta_users</div>
        <div><span style={{ color: "var(--shiki-token-keyword)" }}>ON</span> aws_iam_users.tags.value = okta_users.profile_email</div>
        <div><span style={{ color: "var(--shiki-token-keyword)" }}>WHERE</span> aws_iam_users.tags_key = <span style={{ color: "var(--shiki-token-string-expression)" }}>&quot;email&quot;</span></div>
    </>
)

const QUERIES_EXAMPLES = [
    {
        code: 'SELECT * FROM aws_elbv2_load_balancers WHERE scheme = "internet-facing"',
        html: <SecurityQuery />,
        title: 'Security',
        description: 'Find all public facing AWS load balancers',
    },
    {
        code: 'SELECT account_id, require_uppercase_characters FROM aws_iam_password_policies WHERE require_uppercase_characters = FALSE',
        html: <ComplianceQuery />,
        title: 'Compliance',
        description: 'AWS CIS 1.5 Ensure IAM password policy requires at least one uppercase letter',
    },
    {
        code: 'SELECT arn FROM aws_iam_users JOIN aws_iam_user_tags ON aws_iam_users.id = aws_iam_user_tags.user_id JOIN okta_users ON aws_iam_users.tags.value = okta_users.profile_email WHERE aws_iam_users.tags_key = "email"',
        html: <QueryAcrossApps />,
        title: 'Query across clouds and SaaS apps',
        description: 'Find dormant access keys by joining your AWS IAM users and Okta directory ',
    },
]

export const QueriesExamples = ({ onClick }) => {
    return (
        <div className="grid grid-cols-2 gap-6 my-12 sm:grid-cols-3 ">
            {QUERIES_EXAMPLES.map(({ code, html, title, description }) => (
                <QueryItem
                    onClick={() => onClick(code)}
                    title={title}
                    description={description}
                    key={title}
                >
                    <div className="pb-10">{html}</div>
                </QueryItem>
            ))}
        </div>
    )
}

const QueryItem = ({ children, onClick, description, title }) => {
    const codeClasses = "w-fit h-full px-4 py-3 font-mono text-sm font-medium text-gray-600 bg-black bg-opacity-5 dark:bg-white dark:text-gray-300 dark:bg-opacity-5 betterhover:hover:bg-gray-50 betterhover:dark:hover:bg-gray-900 md:py-3 md:text-base md:leading-6 md:px-10"

    return (
        <div className='flex flex-col'>
            <pre className="h-[260px] bg-transparent dark:bg-transparent m-0 p-0 relative whitespace-pre-wrap border border-transparent border-gray-200 rounded-md dark:border-gray-700">
                <code className={codeClasses}>
                    {children}
                </code>
                <button
                    onClick={onClick}
                    className="absolute bottom-2 right-2 px-2 py-2 text-gray-600 bg-black rounded-md bg-opacity-5 dark:bg-white dark:text-gray-300 dark:border-gray-700 dark:bg-opacity-5 betterhover:hover:bg-gray-50 betterhover:dark:hover:bg-gray-900 md:py-3 md:text-base md:leading-6 md:px-10"
                >
                    <DuplicateIcon className="w-6 h-6 text-gray-400" />
                </button>
            </pre>
            <div className='mt-6'>
                <div className='text-lg font-medium dark:text-white text-center'>{title}</div>
                <div className='mt-2 text-center text-base font-medium text-gray-500 dark:text-gray-400'>{description}</div>
            </div>
        </div>
    )
}

