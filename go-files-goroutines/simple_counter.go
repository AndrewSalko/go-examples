package main

import (
	"fmt"
	"sync"
	"time"
)

func _SimpleCounter(count int, wg *sync.WaitGroup) {
	defer wg.Done() //defer означает что вызов выполнится перед выходом из ф-ции

	for i := 0; i < count; i++ {
		fmt.Println("_SimpleCounter:", i)
		time.Sleep(1 * time.Second)
	}

	fmt.Println("_SimpleCounter finished")
}
