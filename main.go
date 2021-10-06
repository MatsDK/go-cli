package main

import(
	"fmt"
	"os"
)

func main() {
	if len(os.Args) <  2 {
		fmt.Println("off | on | set")
		return
	}

	switch os.Args[1] {
		case "set":
			fmt.Println("set")
		case "on":
			fmt.Println("on")
		case "off":
			fmt.Println("off")
		default: 
			fmt.Println("off | on | set")
	}

}
