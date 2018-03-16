// sync.WaitGroup只有3个方法，Add()，Done()，Wait()。其中Done()是Add(-1)的别名。简单的来说，使用Add()添加计数，Done()减掉一个计数，计数不为0, 阻塞Wait()的运行。
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wp := new(sync.WaitGroup)
	wp.Add(10)

	for i := 0; i < 10; i++ {
		go func(index int) {
			fmt.Println("done ", index)
			time.Sleep(10 * time.Second)
			wp.Done()
		}(i)
	}

	wp.Wait()
	fmt.Println("wait end")
}
