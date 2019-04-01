package main

import (
	"fmt"
	"time"
)

// time.After(duration)を使って、決まった時間を計るタイマーを作る
func main() {
	fmt.Println(time.Now())
	timeChan := make(chan time.Duration)
	go func() {
		timeChan <- 10 * time.Second
	}()
	duration := <- timeChan
	<-time.After(duration)
	fmt.Println(time.Now())
}