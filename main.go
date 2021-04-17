package main

import (
	"github.com/ptcar2009/avro-generator/cmd"
)

func main() {
	err := cmd.MainCommand.Execute()
	if err != nil {
		panic(err)
	}
}
