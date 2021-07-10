package main

import (
	"fmt"
	"github.com/SupreethDhareshwar/go-sqlite/config"
	// log "github.com/SupreethDhareshwar/go-sqlite/log"
	routes "github.com/SupreethDhareshwar/go-sqlite/routes"

	"os"
)

func main() {
	err := config.Set()
	if err != nil {
		os.Exit(1)
	}
	// err = log.Init()
	// if err != nil {
	// 	fmt.Println("Unable to setup logger! ")
	// 	os.Exit(1)
	// }
	r := routes.Init()
	// if err != nil {
	// 	fmt.Println("Unable to setup logger! ")
	// 	os.Exit(1)
	// }
	runOn := config.Get("APPIP") + ":" + config.Get("APPPORT")
	r.Run(runOn)
	fmt.Println("App Started")
}
