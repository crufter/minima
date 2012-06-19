package main

import(
	"time"
	"fmt"
	"github.com/opesun/minima"
)

var src =
`
-- Lambda refers to the last defined lambda function, in this case the function itself.
set n 10000
set z 0
func (run
	mod z (+ z 1)
	if (< z n)
		lambda)
println "Calling a recursive labda " n " times..."
lambda
println "Done."
-- You can't leave out the fib function when you talk about recursion...
set y 15
println "Calculating the " y "th Fibonacci number..."
func fib (x)
	if (| (eq x 0) (eq x 1))
		get x
		+ (fib (- x 1)) (fib (- x 2))
println "Done. Result: " (fib 15)
`

func main() {
	t := time.Now()
	neu := minima.Tokenize(src)
	p := minima.Parse(neu)
	//minima.Visualize(p, "   ", 0)
	minima.Eval(p)
	fmt.Println(time.Since(t))
}