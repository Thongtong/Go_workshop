package main

import (
	"GO_workshop/calculator"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//  var x int
	//  var y func(string) string {
	// 	return "Hello"= name
	//  }

	// func sum(a int , b int) int {

	// }

	// scores := 50
	// _ = scores ///blank for exception
	// number := [5]int{1, 2, 3, 4, 5}

	// for i := 0; i < len(number); i++ {
	// 	println(number[i])
	// }

	// for i, v := range number {
	// 	fmt.Printf("index=%v value=%v\n", i, v)
	// }

	// for _, v := range number { ///ถ้าไม่ใช้ i ให้ blank _
	// 	fmt.Printf( "value=%v\n", v)
	// }

	// func helloworld(name string) string{
	// 	return name
	// }
	// sum := cal(10,10,c func(a, b int) int )int{
	// 	return a*b
	// })
  arg := os.Args[1]
  number , err :=strconv.Atoi(arg)

	if err != nil {
panic(err)
	}
	result := calculator.FizzBuzz(number)
	fmt.Printf(result)

}

// func cal(a, b int, c func(int, int) int) int {
// 	return c(a, b)
// }

// func Hello()(string, int){     ////////tuple value
// 	return "Hello" , 10
// }