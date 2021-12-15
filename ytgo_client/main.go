package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Printf("\nUsage: enter exit to exit, enter link to download")
	for s != "exit" {
		fmt.Printf("\nEnter Youtube Url or exit: ")
		fmt.Scanf("%s ", &s)

	}

}

/*func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}*/
