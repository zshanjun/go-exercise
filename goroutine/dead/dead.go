package dead

import (
	"time"
	"fmt"
	"runtime"
)

func main() {
	//dead()
	//alive()
	//alivePrint()
}

// goroutine是轻量级“线程”
// 非抢占式多任务处理，由协程主动交出控制权
// 编译器/解析器/虚拟机层面的多任务

func dead() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				//由于没有交出控制器的机会（io操作等耗时操作），会一直运行下去
				//main也是一个goroutine，所以下面的sleep也不起作用了
				a[i]++
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}

func alive() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				a[i]++
				//主动交出控制权，让其他线程也有运行的机会
				runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}

func alivePrint() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				a[i]++
				fmt.Println(a[i])
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
}
