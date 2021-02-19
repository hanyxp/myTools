package god

import "fmt"

func Say(name string) {
	str := fmt.Sprintf("vim-go %s", name)

	fmt.Printf("out print: %s \n", str)
}
