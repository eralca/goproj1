// From https://en.wikipedia.org/wiki/Go_%28programming_language%29
// Install with 'go install github.com/user/hello'
//

package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	omitNewline := flag.Bool("n", false, "do not print final newline")
	flag.Parse()

	str := strings.Join(flag.Args(), " ")
	if *omitNewline {
		fmt.Print(str)
	} else {
		fmt.Println(str)
	}
}

