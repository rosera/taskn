package main


import (
	"fmt"
)


const (
	emailToken = "lab-architects@google.com"
	nameToken  = "lab-architects@google.com"
)

func apiOwnerUpdate(emailToken, nameToken string) {
	fmt.Printf("# Lab Owner\n")
	if nameToken == "" {
		fmt.Println(emailToken)
	} else {
		fmt.Printf("%s # %s\n", emailToken, nameToken)
	}
}

func main() {
	apiOwnerUpdate(emailToken, nameToken)
}
