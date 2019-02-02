package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func listdir() {
	file, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range file {
		fmt.Println(f.Name())
	}
}

func main() {

	listdir()
}



