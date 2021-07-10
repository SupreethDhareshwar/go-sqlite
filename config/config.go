package config

import (
	"github.com/joho/godotenv"
)

var myEnv map[string]string

//Set Read the configuration file
func Set() error {
	var err error
	myEnv, err = godotenv.Read()
	if err != nil {
		return err
	}
	return nil
}

//Update function updates the configuration
func Update() {

}

func Get(key string) string {
	return myEnv[key]
}
