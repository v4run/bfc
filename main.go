package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/v4run/bfc/machine"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: bf <filename>")
		os.Exit(1)
	}
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	m, err := machine.New(string(data), os.Stdin, os.Stdout)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	m.Execute()
}
