package main

import (
	"bufio"
	"io"
	"log"
	"net"

	"github.com/golang-book/distributed-systems/demo/jsnet"
)

func main() {
	log.SetFlags(0)

	li := jsnet.Listen()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln("failed to accept new connection:", err)
		}
		go handle(li, conn)
	}
}

func handle(li net.Listener, conn net.Conn) {
	defer conn.Close()

	s := bufio.NewScanner(conn)
	for s.Scan() {
		log.Println("["+li.Addr().String()+":"+conn.LocalAddr().String()+"] received", s.Text())
		io.WriteString(conn, "Hello World\n")
		break
	}

}
