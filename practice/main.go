package main

import "fmt"

const constValue int = 9

func main() {

	var str string = "Hello ! This is Saheeshna"
	fmt.Println(str)
	fmt.Printf("This type is %T\n", str)

	var intValue int = 3
	fmt.Println(intValue)
	fmt.Printf("This type is %T\n", intValue)

	var defaultInt int //var operator works for outside functions
	fmt.Printf("Default int is %T\n", defaultInt)

	newString := "Hi again! from Saheeshna" //works inside functions
	fmt.Println(newString)
	fmt.Printf("This type is %T\n", newString)

	fmt.Println(constValue)

}
