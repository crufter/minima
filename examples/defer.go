package main

import(
	"time"
	"fmt"
	"github.com/opesun/minima"
)

var src =
`
func f (run
	defer (println 0)
	defer (println 1)
	println "This shall run before the deferred functions.")
f
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