package main

import (
	"fmt"
	"github.com/opesun/minima"
	"time"
)

var src = `
set k 10
func l (u) (run
	println k
	+ u u u u)
println (l 20) 
println u
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
