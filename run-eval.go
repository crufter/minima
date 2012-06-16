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
	env := Env{map[string]interface{}{}}
	ev := cmd.Eval(&env)
	return ev
}

func EvalWith(cmd Cmd, inp map[string]interface{}) interface{} {
	defer evalErr()
	env := Env{inp}
	ev := cmd.Eval(&env)
	return ev
}

func Run(src string) interface{} {
	toks := Tokenize(src)
	p := Parse(toks)
	return Eval(p)
}