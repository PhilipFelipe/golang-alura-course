package main

import (
	"fmt"
	"reflect"
)

func main() {
	var name string = "Felipe"
	var age int
	var job = "Backend Developer" // Inferência de tipo
	version := 1.2                // Declaração curta de variável
	fmt.Println("Hello,", name, "you're", age, "years old. And you're a", job)
	fmt.Println("Version", version)

	fmt.Println("The 'job' variable type is", reflect.TypeOf(job))
}
