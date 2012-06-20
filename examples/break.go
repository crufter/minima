package main

import(
	"time"
	"fmt"
	"github.com/opesun/minima"
)

var src =
`
set x 20
set i 0
for x
	if (eq i 10)
		break
		run
			mod i (+ i 1)
			println "i is not 10 yet"
println "after break, i is " i
println "deep break"
for 6
	for 9
		run
			println "yo"
			break 2
println "yo should appear only once above"
for 3
	println "works"
`

func main() {
	t := time.Now()
	neu := minima.Tokenize(src)
	p := minima.Parse(neu)
	//minima.Visualize(p, "   ", 0)
	minima.Eval(p)
	fmt.Println(time.Since(t))
}