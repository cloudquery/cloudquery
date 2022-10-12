package main

import (
	"flag"
	"fmt"
	"go/format"
	"os"
	"regexp"
	"strings"
)

var (
	file        string
	urlTemplate = "https://docs.aws.amazon.com/%s/latest/APIReference/API_%s.html"
)

func getDescriptionsForFile(f string) string {
	b, err := os.ReadFile(f)
	if err != nil {
		panic(err)
	}
	content := string(b)
	name := getPackageName(content)
	r := regexp.MustCompile(`Struct:\s+&types\.(\w+)\{\},`)
	matches := r.FindAllStringSubmatchIndex(content, -1)
	newContent := content
	for _, m := range matches {
		fullMatch := content[m[0]:m[1]]
		start, end := m[2], m[3]
		url := fmt.Sprintf(urlTemplate, name, content[start:end])
		fmt.Println(fullMatch)
		fmt.Println(url)
		newContent = strings.ReplaceAll(newContent, fullMatch, fullMatch+"\n"+fmt.Sprintf(`Description: "%s",`, url))
	}
	formattedContent, err := format.Source([]byte(newContent))
	if err != nil {
		panic(err)
	} else {
		newContent = string(formattedContent)
	}
	return newContent
}

func getPackageName(content string) string {
	mr := regexp.MustCompile(`client.ServiceAccountRegionMultiplexer\("([\w\-]+)"\)`)
	rMatches := mr.FindAllStringSubmatch(content, -1)
	if len(rMatches) > 0 {
		return rMatches[0][1]
	}

	ir := regexp.MustCompile(`github.com/aws/aws-sdk-go-v2/service/(\w+)/types`)
	iMatches := ir.FindAllStringSubmatch(content, -1)
	return iMatches[0][1]
}

func main() {
	flag.StringVar(&file, "f", "", "file to add descriptions for")
	flag.Parse()

	newContent := getDescriptionsForFile(file)
	//fmt.Println(newContent)
	if err := os.WriteFile(file, []byte(newContent), 0644); err != nil {
		panic(err)
	}
}
