package main

import "fmt"

var name = "bessie"
var age = 10
var ok = true

var (
	bes  = "bess"
	time = 29
	best = true
)

func main() {
	var name int
	name = 1
	var a = name + 1

	var hi = 1
	name2 := 10

	fmt.Print(name, a, name2, hi)

	var user1, user2, user3 string
	var u1, u2, u3 = "bes", 21, true //cannot write as: var u1 = "bes", u2 = 21

	fmt.Println(user1, user2, user3, u1, u2, u3)

}
