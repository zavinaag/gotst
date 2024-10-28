package main

import "fmt"

type outs string

type Person struct {
	Name string
	Age  int8
}

func Myscruct() {
	var name outs = "Test outs type"
	print(name)

	var Jo Person
	var Tom = Person{}

	fmt.Println(Jo)
	fmt.Println(Tom)

	Jo.Name = "Jo"
	Jo.Age = 30
	fmt.Println(Jo)
}
