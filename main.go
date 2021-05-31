package main

import (
	"fmt"

	"github.com/jrazmi/gograph-play/internal/hello"
)

func main() {
	message := hello.Hello("Josh")
	fmt.Println(message)
}
