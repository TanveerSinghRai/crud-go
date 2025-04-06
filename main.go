package main

import "fmt"

type student struct {
	name string
	rn   int
}

var s1 student
var n string
var r int

func create(nc string, rc int) {
	fmt.Println("This is create function... I am here to create")
	s1 = student{name: nc}
	s1.rn = rc

}
func delete() {
	fmt.Println("This is delete func... Lets delete")
}
func edit() {
	fmt.Println("This is edit function.. Let me edit")
}
func read() {
	fmt.Println("This is Read fucntion.... Reading")
	fmt.Println(s1)
}
func main() {
	fmt.Println("This is main function \n")
	fmt.Println("Lets create... Enter your name , rollno and marks")
	fmt.Scanln(&n, &r)
	create(n, r)
	read()

}
