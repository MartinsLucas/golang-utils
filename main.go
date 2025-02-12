package main

import (
	"fmt"
	"log"

	"github.com/MartinsLucas/golang-utils/src/util/src/util"
)

func main() {
	// Read a single character and print it
	// char, err := readChar()
	// if err != nil {
	//   log.Fatal(err)
	// }

	// fmt.Printf("%c\n", char)

	// ------------------------------------

	// Read character until a 0 is inputed
	// for {
	//   char, err := readChar()
	//   if err != nil {
	//     log.Fatal(err)
	//     break
	//   }

	//   if char == '\n' {
	//     continue
	//   }

	//   if char == '0' {
	//     break
	//   }

	//   fmt.Printf("Character: %c\n", char)
	// }

	// Read a single line and print it
	str, err := util.ReadLine()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", str)
}
