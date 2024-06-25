package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("changes done in the source")
	fmt.Printf(
		"Hello world from %s/%s\n",
		runtime.GOOS,
		runtime.GOARCH,
	)
}
