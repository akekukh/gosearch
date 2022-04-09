package filestore

import (
	"os"

	"github.com/akekukh/gosearch/task-5/pkg/crawler"
	"github.com/akekukh/gosearch/task-5/pkg/storage"
)

const storePath = "./result.json"

func GetFromStore() ([]crawler.Document, error) {
	f, err := os.Open(storePath)

	if err != nil {
		return nil, err
	}
	data, err := storage.GetDocs(f)

	return data, err
}

func SaveToStore(docs []crawler.Document) error {
	f, err := os.Create(storePath)

	if err != nil {
		return err
	}
	_, err = storage.StoreDocs(docs, f)

	return err
}
