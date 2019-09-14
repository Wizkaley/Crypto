package main

import (
	"fmt"
	"secret"
)

func main() {
	v := secret.File("Eshaaaaan", ".secrets")

	v.Set("Eshan-demo", "some-value ")

	plain, err := v.Get("Eshan-demo")
	if err != nil {
		panic(err)
	}
	fmt.Println("PlainText : ", plain)
}
