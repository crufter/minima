Minima
======

Minima is a toy interpreter written in Go (ofc the language is called the same).
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
if (< n 8)
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