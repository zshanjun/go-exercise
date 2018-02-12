package done

import (
	"fmt"
	"sync"
)

type worker struct {
	in   chan int
	done func()
}

func doWorkOne(id int, w worker) {
	for ch := range w.in {
		fmt.Printf("worker %d recevice %c\n", id, ch)
		w.done()
	}
}

func createWorkerOne(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWorkOne(id, w)
	return w
}

func main() {
	var wg sync.WaitGroup

	workers := [10]worker{}
	for i := 0; i < 10; i++ {
		workers[i] = createWorkerOne(i, &wg)
	}

	wg.Add(20)
	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		workers[i].in <- 'A' + i
	}
	wg.Wait()
}
