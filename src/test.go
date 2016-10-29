package main

import (
	"fmt"
	"time"
)

func sum(values []int, resultChan chan int) {
	sum := 0
	for _, value := range values {
		sum += value
	}
	resultChan <- sum
}

func main() {
	// values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// resultChan := make(chan int, 3)
	// go sum(values[:len(values)/2], resultChan)
	// go sum(values[len(values)/2:], resultChan)
	// go sum(values[len(values)/3:], resultChan)
	// sum1, sum2, sum3 := <-resultChan, <-resultChan, <-resultChan
	// fmt.Println("Result:", sum1, sum2, sum3)

	c = make(chan int)
	go ready("Tea", 2)
	go ready("Coffee", 1)
	<-c
	<-c
}

var c chan int

func ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, "is ready!")
	c <- 1
}
