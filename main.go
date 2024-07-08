package main

import "fmt"

func main() {
	recurse_fuga()
}

func recurse_fuga() {
	fmt.Println("fuga")
	recurse_fuga()
}
