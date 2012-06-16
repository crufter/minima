package main

import(
	"fmt"
	"time"
	"github.com/opesun/minima"
)

var oldstyle =
`
-- This is a comment.
(set n (+ 2 4))
(set x 8)
(if (< n x)
	(run
		(println "Multiline if.")
		(println "n is smaller than " x))
	(println "n is greater or equal than " x))
(for x (println "n is " n))
`
var newstyle =
`
-- This is a comment
set n (+ 2 4)
set x 8
if (< n 8)
	run
		println "Multiline if."
		println "n is smaller than " x
	println "n is greater or equal than " x
for x
	println "n is " n
`

func main() {
	//old := minima.TokenizeOld(oldstyle)
	neu := minima.Tokenize(newstyle)
	//fmt.Println(old)
	//fmt.Println(neu)
	p := minima.Parse(neu)
	//Visualize(p, "   ", 0)
	t := time.Now()
	minima.Eval(p)
	//fmt.Println(ev)
	fmt.Println(time.Since(t))
	//Visualize(p, "   ", 0)
}