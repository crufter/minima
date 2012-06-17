Minima
======

Minima is a toy interpreter written in Go (the language is called the same).
We needed a way to massage our JSON data with a scripting language.

The syntax (or the lack of it) is inspired by Lisp, to be easy to parse for machines.
However, I tried to get rid of the zillions of parentheses to be easy to parse for humans too.
With significant whitespace and indentation, the outermost parentheses are there, but they are kinda transparent.

Approx nothing works yet, apart from that little example below.
No functions, no variable scopes, nothing.

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
n is 6
n is 6
n is 6
n is 6
n is 6
```

One can use the ";" as a shorthand for a newline with same indentation:°
```
set a 10; set b 20
```

Goals
======
- Create a language in pure Go
- Create a scripting language which is statically typed.

Roadmap
======
- []interface{} and map[string]interface{} types to be able to handle JSON data.
- More syntactic sugar (expect some neat things here).
- More builtin goodies.
- Functions, lexical closures, proper variable scoping.
- Static typing.
- Packages.
- A mongodb driver ;)
- Optimizations.