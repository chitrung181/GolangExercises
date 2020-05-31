//Using a select statement, add a timeout to the echo server from Section 8.3 so
//that it disconnects any client that shouts nothing within 10 seconds
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle multiple connection at a time
	}
}

func echo(c net.Conn, shout string, delay time.Duration, wg *sync.WaitGroup) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
	wg.Done()
}

func handleConn(c net.Conn) {
	wg := sync.WaitGroup{}
	defer func() {
		wg.Wait()
		c.Close()
	}()
	text := make(chan string)
	go textScan(c, text)
	delay := 10 * time.Second
	timer := time.NewTimer(10 * time.Second)
	for {
		select {
		case line := <-text:
			timer.Reset(delay)
			wg.Add(1)
			go echo(c, line, 1*time.Second, &wg)
		case <-timer.C:
			return
		}
	}
}

func textScan(r io.Reader, text chan<- string) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		text <- scanner.Text()
	}

	if scanner.Err() != nil {
		fmt.Println("scan: ", scanner.Err())
	}
}
