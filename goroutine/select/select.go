package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("worker %d received %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	var c1, c2 = generator(), generator()
	var worker = createWorker(0)

	var values []int
	// 总执行时间
	tm := time.After(30 * time.Second)
	// 一定时间间隔检查积压数据长度
	tick := time.Tick(time.Second)
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeValue = values[0]
			activeWorker = worker
		}

		select {
		case n := <-c1: //生产内容
			values = append(values, n)
		case n := <-c2: //生成内容
			values = append(values, n)
		case activeWorker <- activeValue: //消费内容
			values = values[1:]

		case <-time.After(800 * time.Millisecond): //两次请求时间间隔超时
			fmt.Println("timeout")
		case <-tick:
			fmt.Println("queue len = ", len(values))
		case <-tm:
			fmt.Println("bye")
			return
		}
	}
}
