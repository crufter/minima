package minima

import(
	"fmt"
	"strconv"
)

const(
	id = iota
	st
	fl
	in
)

func kind(str string) (interface{}, int) {
	if len(str) > 2  && string(str[0]) == `"` && string(str[len(str)-1]) == `"` {
		return str[1:len(str)-1], st
	} else if _int, err := strconv.ParseInt(str, 10, 32); err == nil {
		return int(_int), in
	} else if flo, err := strconv.ParseFloat(str, 32); err == nil {
		return flo, fl
	}
	return str, id
}

type Env struct {
	Symbols map[string]interface{}
}

type Cmd struct {
	Op string
	Params	[]*Cmd
}

func (c *Cmd) Eval(env *Env) interface{} {
	switch c.Op {
	case "run":
		return c.Run(env)
	case "+":
		return c.Add(env)
	case "-":
		return c.Sub(env)
	case "*":
		return c.Mul(env)
	case "/":
		return c.Div(env)
	case "<":
		return c.Lesser(env)
	case "set":
		return c.Set(env)
	case "if":
		return c.If(env)
	case "for":
		return c.For(env)
	case "print":
		return c.Print(env)
	case "println":
		return c.Println(env)
	case "read":
		return c.Read(env)
	}
	return nil
}

func (c *Cmd) Add(env *Env) interface{} {
	var res int
	for _, v := range c.Params{
		if v.Params == nil {
			r, _ := kind(v.Op)
			res += r.(int)
		} else {
			res += v.Eval(env).(int)
		}
	}
	return res
}

func (c *Cmd) Sub(env *Env) interface{} {
	var res int
	for _, v := range c.Params{
		if v.Params == nil {
			r, _ := kind(v.Op)
			res -= r.(int)
		} else {
			res -= v.Eval(env).(int)
		}
	}
	return res
}

func (c *Cmd) Mul(env *Env) interface{} {
	var res int
	for _, v := range c.Params{
		if v.Params == nil {
			r, _ := kind(v.Op)
			res *= r.(int)
		} else {
			res *= v.Eval(env).(int)
		}
	}
	return res
}

func (c *Cmd) Div(env *Env) interface{} {
	var res int
	for _, v := range c.Params{
		if v.Params == nil {
			r, _ := kind(v.Op)
			res /= r.(int)
		} else {
			res /= v.Eval(env).(int)
		}
	}
	return res
}

func (c *Cmd) For(env *Env) interface{} {
	v, k := kind(c.Params[0].Op)
	var val int
	if k == in {
		val = v.(int)
	} else if k == id {
		val = env.Symbols[v.(string)].(int)
	}
	for i:=0;i<val;i++{
		if i == val-1 {
			return c.Params[1].Eval(env)
		}
		c.Params[1].Eval(env)
	}
	return nil
}

func (c *Cmd) Lesser(env *Env) interface{} {
	if c.Params[0].Params == nil && c.Params[1].Params == nil {
		v1, k1 := kind(c.Params[0].Op)
		v2, k2 := kind(c.Params[1].Op)
		var val1, val2 int
		if k1 == id {
			val1 = env.Symbols[v1.(string)].(int)
		} else {
			val1 = v1.(int)
		}
		if k2 == id {
			val2 = env.Symbols[v2.(string)].(int)
		} else {
			val2 = v2.(int)
		}
		return val1 < val2
	}
	return false
}

func (c *Cmd) Run(env *Env) interface{} {
	if len(c.Params) == 1 {
		return c.Params[0].Eval(env)
	}
	l := len(c.Params)
	for i, v := range c.Params {
		if i == l-1 {
			return v.Eval(env)
		}
		v.Eval(env)
	}
	return nil
}

func (c *Cmd) Set(env *Env) interface{} {
	if c.Params[1].Params != nil {
		v := c.Params[1].Eval(env)
		switch v.(type) {
		case int:
			env.Symbols[c.Params[0].Op] = v.(int)
		case string:
			env.Symbols[c.Params[0].Op] = v.(string)
		}
	} else {
		v, k := kind(c.Params[1].Op)
		switch k {
		case id:
			env.Symbols[c.Params[0].Op] = env.Symbols[v.(string)]
		case in:
			env.Symbols[c.Params[0].Op] = v.(int)
		case st:
			env.Symbols[c.Params[0].Op] = v.(string)
		}
	}
	return env.Symbols[c.Params[0].Op]
}

// TODO: current syntax rules makes a the second and third param of if (call anything) tp ((call anything)) which means (run (call anything))
// Same applies to for and all "cotrol structures".
func (c *Cmd) If(env *Env) interface{} {
	v := c.Params[0].Eval(env)
	var ret interface{}
	if v.(bool) {
		ret = c.Params[1].Eval(env)
	} else {
		ret = c.Params[2].Eval(env)
	}
	return ret
}

func (c *Cmd) Print(env *Env) interface{} {
	for _, v := range c.Params {
		val, k := kind(v.Op)
		if k == id {
			fmt.Print(env.Symbols[val.(string)])
		} else {
			fmt.Print(val.(string))
		}
	}
	return 1
}

func (c *Cmd) Println(env *Env) interface{} {
	c.Print(env)
	fmt.Print("\n")
	return 1
}

func (c *Cmd) Read(env *Env) interface{} {
		return nil
}