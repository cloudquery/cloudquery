package manifest

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <path to manifest.json>", os.Args[0])
	}

}
