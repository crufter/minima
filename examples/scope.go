package main

import (
	"fmt"
	"github.com/opesun/minima"
	"time"
)

var src = `
-- This is a comment
set a 1
if (< 0 1)
	run
		set b 2
		if (< 0 1)
			println b
	set c 3
println a b c
`

func main() {
	t := time.Now()
	minima.Run(src)
	fmt.Println(time.Since(t))
}
