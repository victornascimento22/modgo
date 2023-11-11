package main

import "github.com/victornascimento22/modgo/configs"

func main() {

	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)

}