package main

import "fmt"

func main() {
	x := "`"
	y := `package main

import "fmt"

func main() {
	x := "%s"
	y := %s%s%s
	fmt.Printf(y, x, x, y, x)
}
`
	fmt.Printf(y, x, x, y, x)
}
