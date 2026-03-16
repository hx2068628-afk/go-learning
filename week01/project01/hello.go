package main

import "fmt"

func main() {
	test(5,test02)
}

func test02()int{
	fmt.Println("b")
	return 5
}

func test(a int,funcvar func()int){
	fmt.Println(a)
	funcvar()
	

}