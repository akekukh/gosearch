package storage

import (
	"encoding/json"
	"io"

	"github.com/akekukh/gosearch/task-5/pkg/crawler"
)

func StoreDocs(docs []crawler.Document, w io.Writer) (bool, error) {
	data, err := json.Marshal(docs)

	if err != nil {
		return false, err
	}
	_, err = w.Write(data)

	if err != nil {
		return false, err
	}

	return true, err
}

func GetDocs(r io.Reader) ([]crawler.Document, error) {
	var buf = make([]byte, 10)
	var data []byte

	for {
		n, err := r.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if n > 0 {
			data = append(data, buf[:n]...)
		}
	}
	var res []crawler.Document
	json.Unmarshal([]byte(data), &res)

	return res, nil
}
