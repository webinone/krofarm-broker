package main

import (
	"fmt"
	_ "time"
)

func init() {

}

func say(s string) {

}

// 송신 채널
func sendChan ( ch chan <- string) {
	ch <- "Data"
}

// 수신 채널
func receiveChan (ch <- chan string) {
	data := <- ch
	fmt.Println(data)
}

func main() {

	done := make(chan bool, 2) // 버퍼가 2개인 비동기 채널 생성
	count := 4

	go func() {
		for i := 0 ; i < count; i++ {
			done <- true
			fmt.Println("고루틴 : " , i)
		}
	}()

	for i := 0; i < count; i++ {
		<- done
		fmt.Println("메인 함수 :  ", i)
	}
}