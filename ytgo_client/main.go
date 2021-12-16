package main

import (
	"fmt"
	"strings"
)

func main() {
		var s string
		CreateFolder("pic")
		checkErr(err)
		fmt.Printf("Enter Youtube Url or urls separated by ',':")
		fmt.Scanln(&s)
		var urls = strings.Split(s, ",")
		for i:=0; i < len(urls); i++ {
			//fmt.Println(urls[i])

		}
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
