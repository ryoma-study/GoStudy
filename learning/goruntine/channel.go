package main

//var a string
//
//func f() {
//	print(a)
//}
//
//func hello() {
//	a = "hello, world"
//	go f()
//}
//
//func main() {
//	hello()
//}

//var a string
//
//func hello() {
//	go func() { a = "hello" }()
//	print(a)
//}
//
//func main() {
//	hello()
//}

//var c = make(chan int, 10)
//var a string
//
//func f() {
//	a = "hello, world"
//	//c <- 0
//	close(c)
//}
//
//func main() {
//	go f()
//	<-c
//	print(a)
//}

var c = make(chan int)

//var c = make(chan int, 1)
var a string

func f() {
	a = "hello, world"
	<-c
}

func main() {
	go f()
	c <- 0
	print(a)
}
