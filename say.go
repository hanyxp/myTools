package main

import "fmt"

func main() {
	Say("hahah")
}

func Say(name string) string {
	str := fmt.Sprintf("vim-go %s \n", name)
	fmt.Printf("%s", str)
	return str
}
