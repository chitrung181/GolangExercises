//Write a program with two goroutines that send messages back and forth over
//two unbuffered channels in ping-pong fashion. How many communications per second can
//the program sustain?

package main

import (
	"fmt"

	"time"
)

func main() {
	q := make(chan int)
	var i int64
	start := time.Now()
	go func() {
		q <- 1
		for {
			i++
			q <- <-q
		}a
	}()
	go func() {
		for {
			q <- <-q
		}
	}()
		
	timer:= time.NewTimer(time.Second)
	<- timer.C
	fmt.Println(float64(i)/float64(time.Since(start).Seconds()), "times per second")
}