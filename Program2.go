package main

import ( 

	"fmt"
	"strings"
)

func main() {

	input:= "zzz"
	output:= strings.Replace(input, "zzz", "Ashu",-1)

	fmt.Println("Hello " + output + " how are you?")

}