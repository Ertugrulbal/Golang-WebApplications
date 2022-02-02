package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

// multiplexer (mux)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err.Error())
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	request(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			mux(conn, ln)
		} else {
			if ln == "" {
				break
			}
		}
		i++
	}
}

func mux(conn net.Conn, ln string) {
	method := strings.Fields(ln)[0]
	uri := strings.Fields(ln)[1]
	fmt.Println("--> Method : ", method)
	fmt.Println("--> URI : ", uri)

	if method == "GET" && uri == "/" {
		index(conn)
	}

	if method == "GET" && uri == "/about" {
		about(conn)
	}

	if method == "GET" && uri == "/register" {
		register(conn)
	}

	if method == "POST" && uri == "/register" {
		registerPost(conn)
	}

	if method == "GET" && uri == "/contact" {
		contact(conn)
	}
}

func index(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF=8"><title>Burası başlık!</title></head>
  <body>
  <strong>INDEX</strong><br>
  <a href="/">index</a><br>
  <a href="/about">about</a><br>
  <a href="/contact">contact</a><br>
  <a href="/register">register</a><br>
  </body>
  </html>`
	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length : %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, body)
}

func about(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF=8"><title>Burası başlık!</title></head>
  <body>
  <strong>ABOUT</strong><br>
  <a href="/">index</a><br>
  <a href="/about">about</a><br>
  <a href="/contact">contact</a><br>
  <a href="/register">register</a><br>
  </body>
  </html>`
	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length : %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, body)
}

func contact(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF=8"><title>Burası başlık!</title></head>
  <body>
  <strong>CONTACT</strong><br>
  <a href="/">index</a><br>
  <a href="/about">about</a><br>
  <a href="/contact">contact</a><br>
  <a href="/register">register</a><br>
  </body>
  </html>`
	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length : %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, body)
}

func register(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF=8"><title>E-Commerce System!</title></head>
  <body>
  <strong>REGISTER</strong><br>
  <a href="/">index</a><br>
  <a href="/about">about</a><br>
  <a href="/contact">contact</a><br>
  <a href="/register">register</a><br>
  </body>
  </html>`
	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length : %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, body)
}

func registerPost(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF=8"><title>E-Commerce System!</title></head>
  <body>
  <strong>REGISTER (POST)</strong><br>
  <a href="/">index</a><br>
  <a href="/about">about</a><br>
  <a href="/contact">contact</a><br>
  <a href="/register">register</a><br>
  </body>
  </html>`
	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length : %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, body)
}
