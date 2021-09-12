package main

import (
	"fmt"
	"learnPkg/numbers"
	"learnPkg/strings"
	"learnPkg/strings/greeting"
	str "strings" // Alias
)

func main() {
	fmt.Println(numbers.IsPrime(19))
	fmt.Println(greeting.WelcomeText)
	//fmt.Println(strings.Reverse('This is a string reverse'))
	fmt.Println(strings.Reverse("VIDYUT V PRABHU"))

	fmt.Println(str.Count("My name is Vidyut V Prabhu", "Vidyut"))

}
