package main

import (
	"time"
	"log"
	"os"
	"zshanjun/go-exercise/runner"
)

const timeout = 20 * time.Second

func main() {
	log.Println("Start working")

	r := runner.New(timeout)

	r.Add(createTask(), createTask(), createTask())
	r.Add(createTask(), createTask(), createTask())
	r.Add(createTask(), createTask(), createTask())
	r.Add(createTask(), createTask(), createTask())

	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("Terminating due to timeout")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("Terminatig due to interrupt")
			os.Exit(2)
		}
	}

	log.Println("Process finished")
}

func createTask() func(int) {
	return func(i int) {
		time.Sleep(time.Second)
		log.Printf("Processor - Task #%d.\n", i)
	}
}
