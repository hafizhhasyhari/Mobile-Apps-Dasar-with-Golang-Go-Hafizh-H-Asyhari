package main

import "fmt"

func main(){

	fullName := "Alfiansyah Komeng"

	// TODO: convert string to byte
	convertByte := byte(fullName[0])
	// TODO: convert byte to string
	convertString := string(convertByte)


	fmt.Println(fullName)
	fmt.Println(convertByte)
	fmt.Println(convertString)

}
