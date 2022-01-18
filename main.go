package main

import (
	"time"
)

func main() {
	Say(time.Now().Format(time.Kitchen))
}
