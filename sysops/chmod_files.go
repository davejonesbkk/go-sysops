package main 

import (
		"os"
		"log"
		"fmt"
		"flag"

)



func main()	{

		filepath := flag.String("filepath", "file.csv", "The file we are changing the permissions for")

		flag.Parse()

		file, err = os.Chmod(filepath, 0444)
		if err != nil {
			log.Fatal(err)
		}

		file := &filepath

		info, err := os.Stat(file)

		if err != nil {
			fmt.Println("Error:", err)
				os.Exit(1)
		}

		mode := info.Mode()
		fmt.Println(file, ": ", mode, "n")

}