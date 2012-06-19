package main

import(
	"time"
	"fmt"
	"github.com/opesun/minima"
)

var src =
`
set x 10
println (- x 7)
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