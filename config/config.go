package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

var myEnv map[string]string

//Set Read the configuration file
func Set() error {
	myEnv, err := godotenv.Read()
	if err != nil {
		return err
	}
	fmt.Println(myEnv)
	return nil
}

//Update function updates the configuration
func Update() {

}
