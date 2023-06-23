package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

var partner = make(chan io.ReadWriteCloser)

func main() {

	l, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal("error occured:", err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("error occured:", err)
		}
		fmt.Println("new conection from:", conn.LocalAddr())
		go match(conn)
	}

}

func match(c io.ReadWriteCloser) {
	fmt.Fprintln(c, "Looking for a partner ...hold on")
	select {
	case partner <- c:
		//now handled by other goroutine
	case p := <-partner:
		chat(c, p)
	}

}

func chat(a, b io.ReadWriteCloser) {
	fmt.Fprintln(a, "found one")
	fmt.Fprintln(b, "found one")
	ec := make(chan error)
	go cp(a, b, ec)
	go cp(b, a, ec)

	if err := <-ec; err != nil {
		log.Println(err)
	}
	fmt.Fprintln(a, "connection closed")
	fmt.Fprintln(b, "connection closed")
	a.Close()
	b.Close()
}

func cp(r io.Writer, w io.Reader, ec chan<- error) {
	_, err := io.Copy(r, w)
	ec <- err

}
