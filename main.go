package main

import (
	"fmt"
	"os"

	"github.com/SupreethDhareshwar/go-sqlite/config"
)

func main() {
	err := config.Set()
	if err != nil {
		os.Exit(1)
	}
	fmt.Println("App Started")
}
