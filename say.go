package main

import "fmt"

func Say(name string) {
	str := fmt.Sprintf("vim-go %s", name)

	fmt.Printf("god say:out print: %s \n", str)
}
