package index

// Обратный индекс отсканированных документов.

import "github.com/akekukh/gosearch/task-5/pkg/crawler"

// Interface определяет контракт службы индексирования документов.
type Interface interface {
	Add([]crawler.Document)
	Search(string) []int
}
