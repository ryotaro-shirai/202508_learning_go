package main

import "fmt"

func no_1() {
	greetings := []string{"Hello", "Hola", "السلام", "こんにちは", "Привет"}
	subSlice1 := greetings[:2]
	subSlice2 := greetings[1:5]
	subSlice3 := greetings[3:]
	fmt.Println(greetings)
	fmt.Println(subSlice1)
	fmt.Println(subSlice2)
	fmt.Println(subSlice3)
}

func no_2() {
	message := "Hi 👶 and 👧"
	runes := []rune(message)
	fmt.Println(string(runes[3]))
}

func no_3() {
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}
	emp1 := Employee{
		"リオネル",
		"メッシ",
		1,
	}
	emp2 := Employee{
		firstName: "クリスティアーノ",
		lastName:  "ロナウド",
		id:        2,
	}
	var emp3 Employee
	emp3.firstName = "ルカ"
	emp3.lastName = "モドリッチ"
	emp3.id = 3

	fmt.Println(emp1)
	fmt.Println(emp2)
	fmt.Println(emp3)
}

func main() {
	fmt.Println("-----1-----")
	no_1()

	fmt.Println("-----2-----")
	no_2()

	fmt.Println("-----3-----")
	no_3()
}
