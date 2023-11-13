package main

import (
	"fmt"

	"github.com/victornascimento22/modgo/configs"
)

func main() {
	config, _ := configs.LoadConfig(".")
	
	fmt.Println("DBHost:", config.DBHost)
}
