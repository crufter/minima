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
	bo
)

func kind(str string) (interface{}, int) {
	if len(str) > 2  && string(str[0]) == `"` && string(str[len(str)-1]) == `"` {
		return str[1:len(str)-1], st
	} else if _int, err := strconv.ParseInt(str, 10, 32); err == nil {
		return int(_int), in
	} else if flo, err := strconv.ParseFloat(str, 32); err == nil {
		return flo, fl
	} else if str == "true" {
		return true, bo
	} else if str == "false" {
		return false, bo
	}
	return str, id
}

type Func struct {
	Vars_ *Vars
	Args []string
	Com *Cmd
}

func (f *Func) Eval(vars *Vars, params []interface{}) interface{} {
	nvar := &Vars{Sym:make([]map[string]interface{}, 50), Lev:f.Vars_.Lev}	// Support for recursion.
	copy(nvar.Sym, f.Vars_.Sym)
	for i, v := range f.Args {
		nvar.Set(v, params[i])
	}
	return f.Com.Eval(nvar)
}

// I sense some ignorance of multithreading here, but hey, it's just a prototype.
type Vars struct {
	Sym []map[string]interface{}
	Lev int
}

func (v Vars) Get(varname string) interface{} {
	var ret interface{}
	for i:=v.Lev-1; i>=0 ;i-- {
		if v.Sym[i] != nil && len(v.Sym[i]) > 0 {
			ret, ok := v.Sym[i][varname]
			if ok {
				return ret
			}
		}
	}
	return ret
}

// Equals to = in Go.
func (v Vars) Mod(varname string, val interface{}) {
	for i:=v.Lev-1; i>=0 ;i-- {
		if v.Sym[i] != nil && len(v.Sym[i]) > 0 {
			_, ok := v.Sym[i][varname]
			if ok {
				v.Sym[i][varname] = val
			}
		}
	}
}

// Equals to := in Go.
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

// TODO: refactor code to get rid of a lot of evaling inside builtins.
// A var should eval to it's value, a constant to a const etc...
func (c *Cmd) Eval(vars *Vars) interface{} {
	var v interface{}
	if c.Params != nil {
		vars.Lev++
		switch c.Op {
		case "+":
			v = c.Add(vars)
		case "&":
			v = c.And(vars)
		case "/":
			v = c.Div(vars)
		case "eq":
			v = c.Eq(vars)
		case "for":
			v = c.For(vars)
		case "func":
			v = c.Func(vars)
		case "get":
			v = c.Get(vars)
		case "if":
			v = c.If(vars)
		case "<":
			v = c.Less(vars)
		case "list":
			v = c.List(vars)
		case "map":
			v = c.Map(vars)
		case "mod":
			v = c.Mod(vars)
		case "*":
			v = c.Mul(vars)
		case "|":
			v = c.Or(vars)
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
		default:			// Not builtin function call.
			fun := vars.Get(c.Op)
			if val, k := fun.(Func); k {
				params := []interface{}{}
				for _, va := range c.Params {
					params = append(params, va.Eval(vars))
				}
				v = val.Eval(vars, params)
			} else {
				panic("Call of non-function " + c.Op)
			}
		}
		vars.Sym[vars.Lev] = nil
		vars.Lev--
	} else {
		val, ki := kind(c.Op)
		switch ki {
		case id:
			v = vars.Get(val.(string))
		default:
			v = val
		}
	}
	return v
}

func (c *Cmd) Add(vars *Vars) interface{} {
	var res int
	for _, v := range c.Params{
		res += v.Eval(vars).(int)
	}
	return res
}

func (c *Cmd) And(vars *Vars) interface{} {
	for _, v := range c.Params {
		val := v.Eval(vars)
		if value, _ := val.(bool); value == false {
			return false
		}
	}
	if len(c.Params) == 0 {
		return false
	}
	return true
}

func (c *Cmd) Div(vars *Vars) interface{} {
	res := c.Params[0].Eval(vars).(int)
	for i:=1; i<len(c.Params); i++ {
		res /= c.Params[i].Eval(vars).(int)
	}
	return res
}

func (c *Cmd) Eq(vars *Vars) interface{} {
	return c.Params[0].Eval(vars).(int) == c.Params[1].Eval(vars).(int)
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

func (c *Cmd) Func(vars *Vars) interface{} {
	var name string
	co := 0
	if c.Params[0].Params == nil {
		name = c.Params[0].Op
		co++
	} else {
		name = "lambda"
	}
	nvar := &Vars{Sym:make([]map[string]interface{}, 50), Lev:vars.Lev}	// TODO: think about the Lev+1 later.
	copy(nvar.Sym, vars.Sym)
	f := Func{Vars_: nvar}
	if len(c.Params) == co + 2 {		// Has parameters.
		args := []string{c.Params[co].Op}
		for _, v := range c.Params[co].Params {
			args = append(args, v.Op)
		}
		f.Args = args
		co++
	}
	f.Com = c.Params[co]
	vars.Set(name, f)					// Not sure if it will be kept.
	f.Vars_.Set(name, f)
	f.Vars_.Lev++
	// TODO: think about the possible inconsistency what a nils cause when we imagine vars as a []map[string]interface{} in terms of references.
	// For example: x := make(map[string]interface{}, 10); copying it to a new slice and Vars.Setting variables assuming that both will updated will only work
	// if the maps are already existing and not nil.
	return f
}

func (c *Cmd) Get(vars *Vars) interface{} {
	val, _ := kind(c.Params[0].Op)
	return vars.Get(val.(string))
}

// TODO: current syntax rules makes a the second and third param of if (call anything) tp ((call anything)) which means (run (call anything))
// Same applies to for and all "cotrol structures".
func (c *Cmd) If(vars *Vars) interface{} {
	v := c.Params[0].Eval(vars)
	var ret interface{}
	if v.(bool) {
		ret = c.Params[1].Eval(vars)
	} else if len(c.Params) > 2 {
		ret = c.Params[2].Eval(vars)
	}
	return ret
}

func (c *Cmd) Less(vars *Vars) interface{} {
	return c.Params[0].Eval(vars).(int) < c.Params[1].Eval(vars).(int)
}

func (c *Cmd) List(vars *Vars) interface{} {
	return nil
}

func (c *Cmd) Map(vars *Vars) interface{} {
	return nil
}

func (c *Cmd) Mod(vars *Vars) interface{} {
	vname := c.Params[0].Op
	var v interface{}
	v = c.Params[1].Eval(vars)
	vars.Mod(vname, v)
	return v
}

func (c *Cmd) Mul(vars *Vars) interface{} {
	res := 1
	for _, v := range c.Params{
		res *= v.Eval(vars).(int)
	}
	return res
}

func (c *Cmd) Or(vars *Vars) interface{} {
	for _, v := range c.Params {
		val := v.Eval(vars)
		if value, _ := val.(bool); value == true {
			return true
		}
	}
	return false
}

func (c *Cmd) Print(vars *Vars) interface{} {
	for _, v := range c.Params {
		fmt.Print(v.Eval(vars))
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
	v = c.Params[1].Eval(vars)
	vars.Set(vname, v)
	return v
}

func (c *Cmd) Sub(vars *Vars) interface{} {
	var res int
	first := true
	for _, v := range c.Params{
		va := v.Eval(vars).(int)
		if first {
			res = va
		} else {
			res -= va
		}
		first = false
	}
	return res
}