package netsrv

import (
	"bufio"
	"log"
	"net"

	"github.com/akekukh/gosearch/task-11/pkg/crawler"
	"github.com/akekukh/gosearch/task-11/pkg/index/hash"
)

func Handler(conn net.Conn, docs []crawler.Document) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	for {
		msg, _, err := r.ReadLine()
		if err != nil {
			return
		}
		if len(msg) == 0 {
			conn.Write([]byte("Nothing found"))
			return
		}
		conn.Write([]byte("Results:\n"))
		index := hash.New()
		index.Add(docs)
		arr := index.Search(string(msg))
		log.Println("Found:")
		for i := range arr {
			for j := range docs {
				if arr[i] == docs[j].ID {
					log.Print(docs[j].URL)
					conn.Write([]byte(docs[j].URL))
				}
			}
		}
	}
}
