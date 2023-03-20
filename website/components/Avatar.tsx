import Image from "next/image";
import type { Author } from "../content/authors";

export const Avatar = ({
  name,
  url,
  image_url,
  urlText,
  company,
}: Author) => {
  return (
    <div className="flex items-center flex-shrink-0 md:justify-start">
      <div className="w-[32px] h-[32px]">
        <Image
          src={image_url}
          height={32}
          width={32}
          layout="fixed"
          loading="lazy"
          title={name}
          className="w-full rounded-full"
          alt={name}
        />
      </div>
      <dl className="ml-2 text-sm font-medium leading-4 text-left whitespace-no-wrap">
        <dt className="sr-only">Name</dt>
        <dd className="text-gray-900 dark:text-white">{name}</dd>
        <dt className="sr-only">Twitter</dt>
        <dd>
          <span className="text-xs text-gray-500 dark:text-white">{company ? `${company} ` : ""}</span>
        </dd>
          {url ?
            <dd>
                <a
                  href={url}
                  className="text-xs text-blue-500 no-underline betterhover:hover:text-blue-600 betterhover:hover:underline"
                  target="_blank"
                  rel="noopener noreferrer"
                >
                  {urlText}
                </a>
            </dd>
              :
              null
          }
      </dl>
    </div>
  );
};
