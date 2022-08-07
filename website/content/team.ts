const CLOUDQUERY_TEAM = {
  yevgenypats: {
    name: "Yevgeny Pats",
    twitterUsername: "yevgenypats",
    picture: "/images/people/yevgenypats.jpeg",
  },
};

export type Author = keyof typeof CLOUDQUERY_TEAM;
export type AuthorDetails = typeof CLOUDQUERY_TEAM[Author];

export default CLOUDQUERY_TEAM;
