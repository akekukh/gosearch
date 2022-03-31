package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/akekukh/gosearch/task-2/pkg/crawler"
	"github.com/akekukh/gosearch/task-2/pkg/crawler/spider"
)

func main() {
	word := flag.String("s", "", "world for search")
	flag.Parse()

	const depth = 3
	var urls = []string{"https://go.dev", "https://golang.org"}
	docs := scan(urls, depth)

	if *word == "" {
		fmt.Println("Scan result: \n", docs)
		return
	}

	fmt.Printf("Search word '%s' in the results:\n", *word)
	result := find(*word, docs)
	fmt.Println("Find result: \n", result)
}

func scan(urls []string, depth int) (data []crawler.Document) {
	spider := spider.New()

	for _, url := range urls {
		fmt.Println("Scan: ", url)
		docs, err := spider.Scan(url, depth)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
		data = append(data, docs...)
	}

	return data
}

func find(word string, docs []crawler.Document) (responses []crawler.Document) {
	for _, doc := range docs {
		if strings.Contains(strings.ToLower(doc.Title), strings.ToLower(word)) || strings.Contains(strings.ToLower(doc.URL), strings.ToLower(word)) {
			responses = append(responses, doc)
		}
	}
	return responses
}
