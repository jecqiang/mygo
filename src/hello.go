package main

import (
	"fmt"
	//"time"
)

var a int = 10

//const s string = "constant"

type person struct {
	human
	Name string
	Age  int
}

type human struct {
	Sex int
}

func main() {
	// fmt.Printf("%s\n", "hello world")
	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("%d\n", i)
	// }
	// p()
	// q()
	// p()
	//fmt.Printf("%s\n", s)
	// fmt.Println(time.Saturday)
	// fmt.Println(time.Sunday)
	// fmt.Println(time.Now().Weekday())
	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("It is weekend")
	// default:
	// 	fmt.Println("Tt is weekday")
	// }

	// t := time.Now().Day()
	// fmt.Println(t)
	// switch {
	// case t.Hour() < 12:
	// 	fmt.Println("It is before noon")
	// default:
	// 	fmt.Println("It is after noon")
	// }

	// var arr [5]int
	// arr2 := [6]int{1, 2, 3, 4, 5}
	// arr[0] = 1
	// arr[1] = 2
	// append(arr2, 6)
	// fmt.Println(len(arr))
	// fmt.Println(arr2)

	// a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	// s := a[0:]
	// s1 := append(s, 11, 12, 13)
	// sa := a[2:7]
	// sb := sa[3:5]
	// fmt.Println(a, len(a), cap(a))
	// fmt.Println(s, len(s), cap(s))
	// fmt.Println(s1, len(s1), cap(s1))
	// fmt.Println(sa, len(sa), cap(sa))
	// fmt.Println(sb, len(sb), cap(sb))

	// a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	// sa := a[2:5]
	// sa = append(a, 10, 11, 12)
	// fmt.Println(a)
	// fmt.Println(sa)
	// s := make([]int, 10, 10)
	// fmt.Println(s)

	// m := make(map[int]string)
	// fmt.Println(m)
	// m := map[int]string{1: "a", 2: "b", 3: "c"}
	// for _, v := range m {
	// 	fmt.Println(v)
	// }
	//delete(m, 1)
	//fmt.Println(m)

	// a := [4]int{1, 2, 3, 4}
	// fmt.Println(a)

	// a, b := plus(10, 11)
	// fmt.Println(a, b)
	// c := closeue(10)
	// c(12)

	// sp := person{Name: "Jec", Age: 25, human: human{Sex: 1}}

	// sp.structMethod()
	//fmt.Println(sp.human.Sex)
}

func closeue(x int) func(int) {
	f := func(y int) {
		fmt.Println(x + y)
	}
	return f
}

func (s person) structMethod() {
	fmt.Println("struct method")
}

func plus(a, b int) (int, int) {
	a = 20
	b = 21
	return a, b
}

func p() {
	fmt.Printf("%d\n", a)
}

func q() {
	a = 5
	fmt.Printf("%d\n", a)
}
