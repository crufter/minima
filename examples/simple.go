package main

import(
	"fmt"
	"time"
	"github.com/opesun/minima"
)

var oldstyle =
`
-- This is a comment.
(set n (+ 2 1))
(set x 8)
(if (< n x)
	(run
		(println "Multiline if.")
		(println "n is smaller than " x))
	(println "n is greater or equal than " x))
(for n (println "n is " n))
`
var newstyle =
`
-- This is a comment
set n (+ 2 1)
set x 8
if (< n x)
	run
		println "Multiline if."
		println "n is smaller than " x
	println "n is greater or equal than " x
for n
	println "n is " n
`

func main() {
	t := time.Now()
	neu := minima.Tokenize(newstyle)
	p := minima.Parse(neu)
	minima.Eval(p)
	fmt.Println(time.Since(t))
}