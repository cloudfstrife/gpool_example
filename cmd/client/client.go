package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/cloudfstrife/gpool_example/dial"
)

func main() {
	wg := sync.WaitGroup{}
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				send(i, j)
			}
		}(i)
	}
	wg.Wait()
	dial.ClosePool()
	end := time.Now()
	fmt.Println(end.Sub(start))
}

func send(i, j int) {
	conn, err := dial.GetConnection()
	if err != nil {
		log.Fatalf("第%d个线程获取连接失败%v", i, err)
	}
	defer dial.CloseConnection(conn)
	_, err = conn.Write([]byte(fmt.Sprintf("%d %d\n", i, j)))
	if err != nil {
		log.Fatal(err)
	}
}
