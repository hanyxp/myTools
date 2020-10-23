package god

import "fmt"

func Say(name string) string {
	return fmt.Sprintf("vim-go %s", name)
}
