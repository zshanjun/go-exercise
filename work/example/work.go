package main

import (
	"log"
	"time"
	"zshanjun/go-exercise/work"
	"sync"
)

var names = []string{
	"aaa",
	"bbb",
	"ccc",
	"ddd",
	"eee",
}

type namePrinter struct {
	name string
}

func (n *namePrinter) Task() {
	log.Println(n.name)
	time.Sleep(time.Second)
}

func main() {
	p := work.New(2)

	var wg sync.WaitGroup
	wg.Add(100 * len(names))
	for i := 0; i < 100; i++ {
		for _, name := range names {
			np := namePrinter{name:name}
			go func() {
				p.Run(&np)
				wg.Done()
			}()
		}
	}
	wg.Wait()

	p.Shutdown()
}
