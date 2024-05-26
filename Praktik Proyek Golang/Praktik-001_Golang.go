package main

import "fmt"

// TODO: declare var
var name string = "hafizh asyhari"
var age int = 23

// TODO: declare multiple var
var (
	firstName = "hafizh h"
	lastName = "asyhari"
) 

// TODO: declare constant var
const (
	example1 = "hello"
	example2 = "bro"
)

func main() {
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println("Full Name = ", firstName + lastName)
	fmt.Println(example1 + example2)
}




