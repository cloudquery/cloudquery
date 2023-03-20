export interface Author {
    name: string;
    company?: string;
    url?: string;
    urlText?: string;
    image_url: string;
}

export interface Map<T> {
    [index: string]: T;
}

export const EXTERNAL_AUTHORS: Map<Author> = {
    cloudmatt: {
        name: "cloudmatt",
        url: "https://github.com/cloudmatt",
        image_url: "/images/people/cloudmatt.png",
    },
    peterfigueiredo: {
        name: "Peter Figueiredo",
        company: "Hexagon",
        image_url: "/images/people/peterfigueiredo.jpg",
    },
};

export const CLOUDQUERY_TEAM: Map<Author> = {
    benjamin: {
        name: "Ben Bernays",
        url: "https://twitter.com/bbernays",
        urlText: "@bbernays",
        image_url: "/images/people/benjamin.jpg",
    },
    jsonkao: {
        name: "Jason Kao",
        url: "https://www.linkedin.com/in/kaojason/",
        urlText: "kaojason",
        image_url: "/images/people/jason.png",
    },
    hermanschaaf: {
        name: "Herman Schaaf",
        url: "https://github.com/hermanschaaf",
        urlText: "hermanschaaf",
        image_url: "/images/people/hermanschaaf.jpg",
    },
    shimon: {
        name: "Shimon Pats",
        url: "https://www.linkedin.com/in/shimon-pats-592046177/",
        urlText: "shimon-pats",
        image_url: "/images/people/shimon.jpg",
    },
    roneliahu: {
        name: "Ron Eliahu",
        url: "https://twitter.com/p0werhead",
        urlText: "@p0werhead",
        image_url: "/images/people/roneliahu.jpg",
    },
    michelvocks: {
        name: "Michel Vocks",
        url: "https://twitter.com/michelvocks",
        urlText: "@michelvocks",
        image_url: "/images/people/michelvocks.jpg",
    },
    yevgenypats: {
        name: "Yevgeny Pats",
        url: "https://twitter.com/yevgenypats",
        urlText: "@yevgenypats",
        image_url: "/images/people/yevgenypats.jpg",
    },
    mikeelsmore: {
        name: "Mike Elsmore",
        url: "https://twitter.com/ukmadlz",
        urlText: "@ukmadlz",
        image_url: "/images/people/mikeelsmore.png",
    },
    itay: {
        name: "Itay Zagron",
        url: "https://www.linkedin.com/in/zagron/",
        urlText: "zagron",
        image_url: "/images/people/itay.jpg",
    },
    giselatorres: {
        name: "Gisela Torres",
        url: "https://twitter.com/0gis0",
        urlText: "@0gis0",
        image_url: "/images/people/giselatorres.jpeg",
    },
    danielspangenberg: {
        name: "Daniel Spangenberg",
        url: "https://twitter.com/spangenberg_d",
        urlText: "@spangenberg_d",
        image_url: "/images/people/danielspangenberg.jpg",
    },
    SCKelemen: {
        name: "Samuel Kelemen",
        url: "https://www.linkedin.com/in/skelemen/",
        urlText: "skelemen",
        image_url: "/images/people/samuel.png",
    },
    alex: {
        name: "Aleksandr Shcherbakov",
        url: "https://twitter.com/candiduslynx",
        urlText: "@candiduslynx",
        image_url: "/images/people/alex.png",
    },
    kemal: {
        name: "Kemal Hadimli",
        url: "https://twitter.com/disq",
        urlText: "disq",
        image_url: "/images/people/kemal.jpg",
    },
};

export function AuthorByName(name: string) {
    return EXTERNAL_AUTHORS[name] || CLOUDQUERY_TEAM[name];
};