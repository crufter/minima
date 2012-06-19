package main

import(
	"github.com/opesun/minima"
)

var src =
`
println (& false false)
println (& false true)
println (& true false)
println (& (& true false) true)
println (& true (& true true))
println (| false true)	
println (| true false)
println (| false (| true false))
println (| (| false true) false)
`

func main() {
	minima.Run(src)
}