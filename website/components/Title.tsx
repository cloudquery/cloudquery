type PostMeta = {
  title: string;
};


export function Title({ data }: { data: PostMeta }) {
  const title = data.title;

  if (!title) {
    return null;
  }

  return <h1 className="my-10 text-center font-bold leading-tight lg:text-3xl">{title}</h1>;
}