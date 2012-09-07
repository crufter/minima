package main

import (
	"fmt"
	"github.com/opesun/minima"
	"time"
)

var src = `
func l (u) (run
	set m 9
	func (v) (+ v u m))
set p (l 10)
println (p 30)
println (p 40)
`

func main() {
	t := time.Now()
	tok := minima.Tokenize(src)
	p := minima.Parse(tok)
	//minima.Visualize(p, "   ", 0)
	//minima.Run(src)
	minima.Eval(p)
	fmt.Println(time.Since(t))
}
