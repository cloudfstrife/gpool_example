package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:11211")
	if err != nil {
		log.Fatal(err)
	}
	i := 0
	for {
		conn, err := listener.Accept()
		i++
		if err != nil {
			log.Fatal(err)
		}
		go Proc(conn, i)
	}

}

//Proc Proc
func Proc(conn net.Conn, i int) {
	defer conn.Close()
	buf := bufio.NewReader(conn)
	for {
		bs, _, err := buf.ReadLine()
		if err != nil {
			return
		}
		v := string(bs)
		slist := strings.Split(v, " ")

		if len(slist) == 2 {
			if slist[1] == "999" {
				fmt.Printf("第%d个连接收到消息 : 线程 %s 第 %s 次发送\n", i, slist[0], slist[1])
			}
		} else {
			fmt.Println(v)
		}
	}
}
