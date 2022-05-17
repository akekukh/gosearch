package main

import (
	"fmt"
	"net"

	"github.com/akekukh/gosearch/task-11/pkg/crawler"
	"github.com/akekukh/gosearch/task-11/pkg/crawler/spider"
	"github.com/akekukh/gosearch/task-11/pkg/netsrv"
)

var urls = []string{"https://go.dev", "https://golang.org"}

const network, addr, depth = "tcp4", "0.0.0.0:8000", 3

func main() {
	spider := spider.New()
	var data []crawler.Document
	for _, url := range urls {
		fmt.Println("Scan: ", url)
		docs, err := spider.Scan(url, depth)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
		data = append(data, docs...)
	}

	fmt.Println("Finished scan!")
	listener, err := net.Listen(network, addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go netsrv.Handler(conn, data)
	}
}
