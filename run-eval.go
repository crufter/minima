package minima

import(
	"fmt"
)

func evalErr() {
	r := recover();
	if r != nil {
		fmt.Println("An error during eval occured:", r)
	}
}

func Eval(cmd Cmd) interface{} {
	//defer evalErr()
	vars := Vars{make([]map[string]interface{}, 50), 0}
	ev := cmd.Eval(&vars)
	return ev
}

func EvalWith(cmd Cmd, inp map[string]interface{}) interface{} {
	//defer evalErr()
	vars := Vars{make([]map[string]interface{}, 50), 0}
	vars.Sym[0] = map[string]interface{}{"en":inp}
	ev := cmd.Eval(&vars)
	return ev
}

func Run(src string) interface{} {
	toks := Tokenize(src)
	p := Parse(toks)
	return Eval(p)
}