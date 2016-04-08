package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

const (
	AUTH       = "_A:"
	HB         = "_H:"
	MSG_PREFIX = "_M:"
	NEW_LINE   = "\n"
)

func main() {
	for i := 0; i < 3000; i++ {
		time.Sleep(10 * time.Millisecond)
		go run(strconv.Itoa(i))
	}
	time.Sleep(24 * time.Hour)
}

func run(uid string) {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("dial error:", err.Error())
	}

	conn.Write([]byte(AUTH + uid))
	reader := bufio.NewReader(conn)

	bb, _, err := reader.ReadLine()

	fmt.Println("AUTH response:" + string(bb))

	go hb(conn)

	for {
		bb, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("read error:", err.Error())
			break
		} else {

			respStr := string(bb)
			if !strings.HasPrefix(respStr, HB) {
				fmt.Println("Server response:[" + respStr + "][" + uid + "]")
				if strings.HasPrefix(respStr, MSG_PREFIX) {
					_, err := conn.Write([]byte(MSG_PREFIX + "OK" + NEW_LINE))
					if err != nil {
						fmt.Println("write error:", err.Error())
						break
					}
				}
			}
		}
	}
}

func hb(conn net.Conn) {
	for {
		_, err := conn.Write([]byte(HB + NEW_LINE))
		if err != nil {
			fmt.Println("write error:", err.Error())
			break
		}

		time.Sleep(6 * time.Second)
	}
}
