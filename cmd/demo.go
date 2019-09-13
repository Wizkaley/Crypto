package main

import (
	"fmt"
	"secret"
)

func main() {
	v := secret.Memory("fake-keys")

	v.Set("demo-key", "some-value ")

	plain, err := v.Get("demo-key")
	if err != nil {
		panic(err)
	}
	fmt.Println("PlainText : ", plain)
}
