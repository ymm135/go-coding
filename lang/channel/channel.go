package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for n := range c { //使用range也行, 也可以使用n,ok := <-c
		fmt.Printf("Worker %d received %c\n",
			id, n)
	}
}

func createWorker(id int) chan<- int {//chan<- int表明用于发送数据
	c := make(chan int)               //<-chan int表明用于接收数据
	go worker(id, c)
	return c
}

func chanDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	//缓冲区时3个, 缓存超过3个以后才会block
	c := make(chan int, 3)  //channel 也是一等公民, 可以作为参数
	go worker(0, c)
	c <- 'a' // channel作为协程之间的通道, 如果存进没有取出, 就会block
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	fmt.Println("Channel as first-class citizen")
	chanDemo()
	fmt.Println("Buffered channel")
	bufferedChannel()
	fmt.Println("Channel close and range")
	channelClose()
}
