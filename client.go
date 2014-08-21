package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

const (
	AUTH       = "_A:"
	HB         = "_H:"
	MSG_PREFIX = "_M:"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	checkError(err)

	conn.Write([]byte(AUTH + "upton"))
	reader := bufio.NewReader(conn)

	bb, _, err := reader.ReadLine()

	fmt.Println("AUTH response:" + string(bb))

	//conn.Close()

	go hb(conn)

	for {
		bb, _, err := reader.ReadLine()
		checkError(err)

		fmt.Println("Server response:[" + string(bb) + "]")
	}
}

func hb(conn net.Conn) {

	for {
		conn.Write([]byte(HB))
		time.Sleep(2 * time.Second)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err.Error())
		os.Exit(1)
	}
}
