package main

import (
	"fmt"

	depa "github.com/phoenix2x/bot-test-dep-a"
	depb "github.com/phoenix2x/bot-test-dep-b"
)

func main() {
	fmt.Println(depb.Dummy)

	depa.F()
}
