package main 

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file, err := ioutil.ReadFile("test.txt") 
	if err != nil {
		fmt.Println(err)
	}

	str := string(file) //converts file contents from bytes to string
	fmt.Println(str)
}


