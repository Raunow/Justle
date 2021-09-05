package main

import (
	// "flag"
	"fmt"

	"github.com/raunow/justle/src/config"
)

func main() {
	config.ReadConfig()
	fmt.Println("Hello World!")

}
