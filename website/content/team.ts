const CLOUDQUERY_TEAM = {
  benjamin: {
    name: "Ben Bernays",
    url: "https://twitter.com/bbernays",
    twitterUsername: "bbernays",
    image_url: "/images/people/benjamin.jpg",
  },
  jsonkao: {
    name: "Jason Kao",
    url: "https://www.linkedin.com/in/kaojason/",
    twitterUsername: "",
    image_url: "/images/people/jason.png",
  },
  hermanschaaf: {
    name: "Herman Schaaf",
    url: "https://github.com/hermanschaaf",
    twitterUsername: "",
    image_url: "/images/people/hermanschaaf.jpg",
  },
  roneliahu: {
    name: "Ron Eliahu",
    url: "https://twitter.com/p0werhead",
    twitterUsername: "p0werhead",
    image_url: "/images/people/roneliahu.jpg",
  },
  michelvocks: {
    name: "Michel Vocks",
    url: "https://twitter.com/michelvocks",
    twitterUsername: "michelvocks",
    image_url: "/images/people/michelvocks.jpg",
  },
  yevgenypats: {
    name: "Yevgeny Pats",
    url: "https://twitter.com/yevgenypats",
    twitterUsername: "yevgenypats",
    image_url: "/images/people/yevgenypats.jpg",
  },
  mikeelsmore: {
    name: "Mike Elsmore",
    url: "https://twitter.com/ukmadlz",
    twitterUsername: "ukmadlz",
    image_url: "/images/people/mikeelsmore.png",
  },
  itay: {
    name: "Itay Zagron",
    url: "https://www.linkedin.com/in/zagron/",
    twitterUsername: "",
    image_url: "/images/people/itay.jpg",
  },
  giselatorres: {
    name: "Gisela Torres",
    url: "https://twitter.com/0gis0",
    twitterUsername: "0gis0",
    image_url: "/images/people/giselatorres.jpeg",
  },
  danielspangenberg: {
    name: "Daniel Spangenberg",
    url: "https://twitter.com/spangenberg_d",
    twitterUsername: "spangenberg_d",
    image_url: "/images/people/danielspangenberg.jpg",
  },
  SCKelemen: {
    name: "Samuel Kelemen",
    url: "https://www.linkedin.com/in/skelemen/",
    twitterUsername: "",
    image_url: "/images/people/samuel.png",
  },
  alex: {
    name: "Aleksandr Shcherbakov",
    url: "https://www.linkedin.com/in/alex-shcherbakov/",
    twitterUsername: "candiduslynx",
    image_url: "/images/people/alex.png",
  },
};

export type Author = keyof typeof CLOUDQUERY_TEAM;
export type AuthorDetails = typeof CLOUDQUERY_TEAM[Author];

export default CLOUDQUERY_TEAM;
