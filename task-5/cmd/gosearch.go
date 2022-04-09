package main

import (
	"flag"
	"fmt"
	"os"
	"sort"

	"github.com/akekukh/gosearch/task-5/pkg/crawler"
	"github.com/akekukh/gosearch/task-5/pkg/crawler/spider"
	"github.com/akekukh/gosearch/task-5/pkg/index/hash"
	"github.com/akekukh/gosearch/task-5/pkg/storage/filestore"
)

const depth = 3

func main() {
	word := flag.String("s", "", "world for search")
	flag.Parse()
	var urls = []string{"https://go.dev", "https://golang.org"}
	docs, err := filestore.GetFromStore()

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File not found, scaning urls..")
		} else {
			fmt.Println("Error: ", err)
			return
		}
	}

	if docs == nil {
		docs = scan(urls, depth)
		err = filestore.SaveToStore(docs)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
	}

	for i := range docs {
		docs[i].ID = i
	}

	sort.Slice(docs, func(i, j int) bool { return docs[i].ID <= docs[j].ID })

	hash := hash.New()
	hash.Add(docs)

	fmt.Println("Scan result: \n", docs)

	if *word == "" {
		fmt.Println("Error: ", "set word for scan (f.e. -s Google )")
		return
	}

	fmt.Printf("Search word '%s' in the results:\n", *word)
	result := find(*word, docs, hash)
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
