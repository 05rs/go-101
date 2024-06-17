package main

import (
	"fmt"
	"time"
)

func work(tasks []string) {
	for _, task := range tasks {
		time.Sleep(1 * time.Second)
		fmt.Println(task, "completed")
	}

}

func main() {
	tasks := []string{"1", "2", "3", "4"}
	go work(tasks[:2])
	work(tasks[2:])

}
