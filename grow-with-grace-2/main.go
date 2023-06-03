package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"html/template"
	"io"
	"log"
	"net/http"
)

var rootTemplate = template.Must(template.New("root").Parse(`
<!DOCTYPE html>
<html>

<body> <h1>hello</h1>
	<script>
    websocket = new WebSocket("ws://{{.}}/socket");
	</script>
</body>
</html>
`))

var partner = make(chan io.ReadWriteCloser)

type socket struct {
	conn *websocket.Conn
	done chan bool
}

func (s socket) Read(b []byte) (n int, err error) {
	return s.conn.Read(b)
}

func (s socket) Write(b []byte) (n int, err error) {
	return s.conn.Write(b)
}

func (s socket) Close() error {
	s.done <- true
	return nil
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.Handle("/socket", websocket.Handler(socketHandler))
	if err := http.ListenAndServe("localhost:3000", nil); err != nil {
		log.Fatal(err)
	}

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := rootTemplate.Execute(w, "localhost:3000")
	if err != nil {
		fmt.Println(err)
	}
}

func socketHandler(c *websocket.Conn) {
	sock := socket{conn: c, done: make(chan bool)}
	go match(sock)
	<-sock.done

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

	a.Close()
	b.Close()
}

func cp(r io.Writer, w io.Reader, ec chan<- error) {
	_, err := io.Copy(r, w)
	ec <- err

}
