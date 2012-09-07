package minima

// Half ugly optimization (doesn't help much btw, maybe 10-25% and more consistent performance.)
var builtins = [...]func(*Cmd, *Vars) interface{}{
	(*Cmd).Add,
	(*Cmd).And,
	(*Cmd).Break,
	(*Cmd).Defer,
	(*Cmd).Div,
	(*Cmd).Eq,
	(*Cmd).For,
	(*Cmd).Func,
	(*Cmd).Get,
	(*Cmd).If,
	(*Cmd).Less,
	(*Cmd).List,
	(*Cmd).Map,
	(*Cmd).Mod,
	(*Cmd).Mul,
	(*Cmd).Or,
	(*Cmd).Panic,
	(*Cmd).Print,
	(*Cmd).Println,
	(*Cmd).Read,
	(*Cmd).Recover,
	(*Cmd).Run,
	(*Cmd).Set,
	(*Cmd).Sub,
}

func builtinNum(str string) int {
	switch str {
	case "+":
		return 1
	case "&":
		return 2
	case "break":
		return 3
	case "defer":
		return 4
	case "/":
		return 5
	case "eq":
		return 6
	case "for":
		return 7
	case "func":
		return 8
	case "get":
		return 9
	case "if":
		return 10
	case "<":
		return 11
	case "list":
		return 12
	case "map":
		return 13
	case "mod":
		return 14
	case "*":
		return 15
	case "|":
		return 16
	case "panic":
		return 17
	case "print":
		return 18
	case "println":
		return 19
	case "read":
		return 20
	case "recover":
		return 21
	case "run":
		return 22
	case "set":
		return 23
	case "-":
		return 24
	default:
		return -1
	}
	return -1
}
