package main

import(
	"time"
	"fmt"
	"github.com/opesun/minima"
)

var src =
`
func k (panic "OMG")
func f (run
	recover (run(println "Recovering from " prob) (+ 1 1))
	println "Panicking in next function call."
	k
	println "This shall not run.")
func l (run
	println "Just a casual println..."
	set ret (f)
	println "This shall run."
	get ret)
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