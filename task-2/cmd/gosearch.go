package main

import (
	"flag"
	"fmt"
	"sort"

	"github.com/akekukh/gosearch/task-2/pkg/crawler"
	"github.com/akekukh/gosearch/task-2/pkg/crawler/spider"
	"github.com/akekukh/gosearch/task-2/pkg/index/hash"
)

func main() {
	word := flag.String("s", "", "world for search")
	flag.Parse()

	const depth = 3
	var urls = []string{"https://go.dev", "https://golang.org"}
	docs := scan(urls, depth)
	for i := range docs {
		docs[i].ID = i
	}
	sort.Slice(docs, func(i, j int) bool { return docs[i].ID <= docs[j].ID })
	indexer := index(docs)

	fmt.Println("Scan result: \n", docs)
	if *word == "" {
		return
	}

	fmt.Printf("Search word '%s' in the results:\n", *word)
	result := find(*word, docs, indexer)
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

func index(docs []crawler.Document) *hash.Index {
	hash := hash.New()
	hash.Add(docs)
	return hash
}

func find(word string, docs []crawler.Document, indexer *hash.Index) (responses []crawler.Document) {
	ids := indexer.Search(word)
	for _, id := range ids {
		docId := sort.Search(len(docs), func(i int) bool { return docs[i].ID >= id })
		if docId < len(docs) && docs[docId].ID == id {
			responses = append(responses, docs[docId])
		}
	}
	return responses
}
