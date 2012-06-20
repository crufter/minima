package main

import(
	"time"
	"fmt"
	"github.com/opesun/minima"
)

var src =
`
func l (run
	recover (+ 1 1)
	println "Just a casual println..."
	panic "Get out of here."
	println "This shall not run.")
println (l)
println "Recovered"
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