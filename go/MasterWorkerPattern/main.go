package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, ch chan string, workerName string) {
	defer wg.Done()
	for v := range ch {
		fmt.Println(workerName, "is working on", v)
		//do stuff
		time.Sleep(time.Second)
	}
}

func master(wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
	defer close(ch)
	commands := []string{
		"walking",
		"waiting",
		"going forward",
		"going backward",
		"going left",
		"going right",
	}

	for _, v := range commands {
		ch <- v
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan string)
	wg.Add(3)
	go worker(&wg, ch, "John")
	go worker(&wg, ch, "Ricky")
	go master(&wg, ch)
	wg.Wait()
}
