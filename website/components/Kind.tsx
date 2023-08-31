export type Kind = "source" | "destination";

export const KINDS: Kind[] = ["source", "destination"];

export function KindTitle(k: Kind) {
    switch (k) {
        case "source":
            return "Source";
        case "destination":
            return "Destination";
        default:
            return "Unknown";
    }
}