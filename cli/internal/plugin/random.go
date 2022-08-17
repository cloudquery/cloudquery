package plugin

import (
	"math/rand"
	"os"
	"path"
)

var unixSocketDir = os.TempDir()

// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go/22892986#22892986

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func generateRandomUnixSocketName() string {
	return path.Join(unixSocketDir, "cq-"+randSeq(16)+".sock")
}
