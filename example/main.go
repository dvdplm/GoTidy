package main

import (
	"fmt"
	"log"
	"tidy"
)

func main() {
	output, err := tidy.Tidy("<title>Foo</title><p>Foo!")
	if err != nil {
		fmt.Println(output)
		log.Fatal(err)
	}
	fmt.Println(output)
}