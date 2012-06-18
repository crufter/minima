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

// I sense some ignorance of multithreading here, but hey, it's just a prototype.
type Vars struct {
	Sym []map[string]interface{}
	Lev int
}

func (v Vars) Get(varname string) interface{} {
	var ret interface{}
	for i:=v.Lev-1; i>=0 ;i-- {
		if v.Sym[i] != nil {
			ret, ok := v.Sym[i][varname]
			if ok {
				return ret
			}
		}
	}
	return ret
}

func (v Vars) Set(varname string, val interface{}) {
	if v.Sym[v.Lev-1] == nil {
		v.Sym[v.Lev-1] = map[string]interface{}{}
	}
	v.Sym[v.Lev-1][varname] = val
}

type Cmd struct {
	Op string
	Params	[]*Cmd
}

func (c *Cmd) Eval(vars *Vars) interface{} {
	vars.Lev++
	var v interface{}
	switch c.Op {
	case "+":
		v = c.Add(vars)
	case "/":
		v = c.Div(vars)
	case "<":
		v = c.Less(vars)
	case "if":
		v = c.If(vars)
	case "for":
		v = c.For(vars)
	case "*":
		v = c.Mul(vars)
	case "print":
		v = c.Print(vars)
	case "println":
		v = c.Println(vars)
	case "read":
		v = c.Read(vars)
	case "run":
		v = c.Run(vars)
	case "set":
		v = c.Set(vars)
	case "-":
		v = c.Sub(vars)
	}
	vars.Lev--
	return v
}

func (c *Cmd) Add(vars *Vars) interface{} {
	var res int
	for _, v := range c.Params{
		if v.Params == nil {
			r, _ := kind(v.Op)
			res += r.(int)
		} else {
			res += v.Eval(vars).(int)
		}
	}
	return res
}

func (c *Cmd) Div(vars *Vars) interface{} {
	var res int
	for _, v := range c.Params{
		if v.Params == nil {
			r, _ := kind(v.Op)
			res /= r.(int)
		} else {
			res /= v.Eval(vars).(int)
		}
	}
	return res
}

// TODO: current syntax rules makes a the second and third param of if (call anything) tp ((call anything)) which means (run (call anything))
// Same applies to for and all "cotrol structures".
func (c *Cmd) If(vars *Vars) interface{} {
	v := c.Params[0].Eval(vars)
	var ret interface{}
	if v.(bool) {
		ret = c.Params[1].Eval(vars)
	} else {
		ret = c.Params[2].Eval(vars)
	}
	return ret
}

func (c *Cmd) For(vars *Vars) interface{} {
	v, k := kind(c.Params[0].Op)
	var val int
	if k == in {
		val = v.(int)
	} else if k == id {
		val = vars.Get(v.(string)).(int)
	}
	for i:=0;i<val;i++{
		if i == val-1 {
			return c.Params[1].Eval(vars)
		}
		c.Params[1].Eval(vars)
	}
	return nil
}

func (c *Cmd) Less(vars *Vars) interface{} {
	if c.Params[0].Params == nil && c.Params[1].Params == nil {
		v1, k1 := kind(c.Params[0].Op)
		v2, k2 := kind(c.Params[1].Op)
		var val1, val2 int
		if k1 == id {
			val1 = vars.Get(v1.(string)).(int)
		} else {
			val1 = v1.(int)
		}
		if k2 == id {
			val2 = vars.Get(v2.(string)).(int)
		} else {
			val2 = v2.(int)
		}
		return val1 < val2
	}
	return false
}

func (c *Cmd) Mul(vars *Vars) interface{} {
	var res int
	for _, v := range c.Params{
		if v.Params == nil {
			r, _ := kind(v.Op)
			res *= r.(int)
		} else {
			res *= v.Eval(vars).(int)
		}
	}
	return res
}

func (c *Cmd) Print(vars *Vars) interface{} {
	for _, v := range c.Params {
		val, k := kind(v.Op)
		if k == id {
			fmt.Print(vars.Get(val.(string)))
		} else {
			fmt.Print(val)
		}
	}
	return 1
}

func (c *Cmd) Println(vars *Vars) interface{} {
	c.Print(vars)
	fmt.Print("\n")
	return 1
}

func (c *Cmd) Read(vars *Vars) interface{} {
		return nil
}

func (c *Cmd) Run(vars *Vars) interface{} {
	if len(c.Params) == 1 {
		return c.Params[0].Eval(vars)
	}
	l := len(c.Params)
	for i, v := range c.Params {
		if i == l-1 {
			return v.Eval(vars)
		}
		v.Eval(vars)
	}
	return nil
}

func (c *Cmd) Set(vars *Vars) interface{} {
	vname := c.Params[0].Op
	var v interface{}
	if c.Params[1].Params != nil {
		v = c.Params[1].Eval(vars)
		vars.Set(vname, v)
	} else {
		v, k := kind(c.Params[1].Op)
		switch k {
		case id:
			vars.Set(vname, vars.Get(v.(string)))
		default:
			vars.Set(vname, v)
		}
	}
	return v
}

func (c *Cmd) Sub(vars *Vars) interface{} {
	var res int
	for _, v := range c.Params{
		if v.Params == nil {
			r, _ := kind(v.Op)
			res -= r.(int)
		} else {
			res -= v.Eval(vars).(int)
		}
	}
	return res
}