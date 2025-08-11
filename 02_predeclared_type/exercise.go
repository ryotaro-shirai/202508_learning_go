package main

import "fmt"

func no_1() {
	var i int = 20
	f := float64(i)
	fmt.Println(i)
	fmt.Println(f)
}

func no_2() {
	const value = 30
	i := value
	var f float64 = value
	fmt.Println(i)
	fmt.Println(f)
}

func no_3() {
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615
	b += 1
	smallI += 1
	bigI += 1
	fmt.Println(b)
	fmt.Println(smallI)
	fmt.Println(bigI)
}

func main() {
	fmt.Println("-----1-----")
	no_1()

	fmt.Println("-----2-----")
	no_2()

	fmt.Println("-----3-----")
	no_3()
}
