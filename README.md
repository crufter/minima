Minima
======

Minima is an experimental interpreter written in Go (the language is called the same).
We needed a way to massage our JSON data with a scripting language.

The syntax (or the lack of it) is inspired by Lisp, to be easy to parse for machines.
However, I tried to get rid of the zillions of parentheses to be easy to parse for humans too.
With significant whitespace and indentation, the outermost parentheses are there, but they are kinda transparent.

Everything is subject to change.

```
-- This is a comment
set n (+ 2 4)
set x 8
if (< n x)
	run
		println "Multiline if."
		println "n is smaller than " x
	println "n is greater or equal than " x
for x
	println "n is " n
```

This above snippet will produce:

```
Multiline if.
n is smaller than 8
n is 6
n is 6
n is 6
```

One can use the ";" as a shorthand for a newline with same indentation:
```
set a 10; set b 20
```

Function definition and call:
```
set k 10
func l (u) (run
	println k
	+ u u u u)
println (l 20) 
println u
```

Produces:
```
10
80
<nil>
```

Closures are working too:
```
func l (u) (run
	set m 9
	func (v) (+ v u m))
set p (l 10)
println (p 30)
println (p 40)
```

Produces:
```
49
59
```

Goals
======
- Create a language in pure Go.
- Create a scripting language which is statically typed.

Latest additions
======
- Better recursion support.
- Eq, |, & operators.
- Variable scoping, functions, closures.

Roadmap
======
- []interface{} and map[string]interface{} types to be able to handle JSON data.
- More syntactic sugar (expect some neat things here).
- More builtin goodies.
- Static typing.
- Packages.
- A mongodb driver ;)
- Optimizations.