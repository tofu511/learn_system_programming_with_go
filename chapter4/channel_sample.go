package main

import "fmt"

func main() {
	fmt.Println("start sub()")
	done := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("sub() is finished")
			done <- true
		}
	}()
	<-done
	fmt.Println("all tasks are finished")
}
