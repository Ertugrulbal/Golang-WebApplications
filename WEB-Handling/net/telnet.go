package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

// Telnet : acronis.com/en-us/articles/telnet/
// to Unicode : https://tldp.org/HOWTO/Unicode-HOWTO-4.html (kermit)

// streaming
func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		io.WriteString(conn, "\nHello from TCP Server!\n")
		fmt.Fprintf(conn, "Go nasıl ama?")
		fmt.Fprintf(conn, "%v", "Güzel gidiyyaa")

		conn.Close()
	}
}
