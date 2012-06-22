package minima

import(
	"fmt"
	"github.com/opesun/lexer"
	"strconv"
)

const(
	itemIgnore = iota
	itemRParen
	itemLParen
	itemInt
	itemID
	itemString
	itemNewLine
	itemSemi
	itemTab
)

var token_exrps_clear = []lexer.TokenExpr{
    {`[ ]+`,										itemIgnore},	// Whitespace
    {`\-\-[^\n]*`,									itemIgnore},	// Comment
	{`\n`,											-itemNewLine},	// Newline
	{`\t`,											-itemTab},
	{`;`,											itemSemi},
    {`\(`,											itemRParen},
    {`\)`,											itemLParen},
    {`[0-9]+`,										itemInt},
	{`"(?:[^"\\]|\\.)*"`,							itemString},
    {`[\<\>\!\=\+\-\|\&\*\/A-Za-z][A-Za-z0-9_]*`,	itemID},
}

// From beginning or end of source.
func clearNewlines(src string) string {
	if string(src[0]) == "\n" {
		src = src[1:]
	}
	if string(src[len(src)-1]) == "\n" {
		src = src[:len(src)-1]
	}
	return src
}

// This is where we handle all the new style rules, we transform to old style simply.
func Tokenize(source string) []string {
	source = clearNewlines(source)
	tokens, _ := lexer.Lex("(\n" + source + "\n)" , token_exrps_clear)
	toks := []string{}
	last_ind := 0
	for i := 0; i<len(tokens);i++ {
		v := tokens[i]
		if v.Text == "\n" {
			var next_ind int
			if tokens[i+1].Text != "\t" {
				next_ind = 0
			} else {
				next_ind = tokens[i+1].Occ
			}
			diff := next_ind - last_ind
			last_ind = next_ind 
			if len(tokens) != i + 1 && i > 0 && tokens[i-1].Text != "(" {
				if diff <= 0 {
					toks = append(toks, ")")			// 1 implicit záró
					for i:= 0; i<diff*(-1);i++ {		// plusz amennyit csökken
						toks = append(toks, ")")
					}
				}
			}
			if len(toks) > 0 && len(tokens) > i + 2 && tokens[i+2].Text != ")" {
				toks = append(toks, "(")
			}
		} else if v.Text == ";" {
			toks = append(toks, ")")
			toks = append(toks, "(")
		} else if v.Text != "\t" {
			toks = append(toks, v.Text)
		}
	}
	return toks
}

func parsErr() {
	r := recover();
	if r != nil {
		fmt.Println("An error during parsing occured:", r)
	}
}

func kind(str string) (interface{}, int) {
	if len(str) > 2  && string(str[0]) == `"` && string(str[len(str)-1]) == `"` {
		return str[1:len(str)-1], st
	} else if _int, err := strconv.ParseInt(str, 10, 32); err == nil {
		return int(_int), in
	} else if flo, err := strconv.ParseFloat(str, 32); err == nil {
		return flo, fl
	} else if str == true_str {
		return true, bo
	} else if str == false_str {
		return false, bo
	}
	return str, id
}

func Parse(tokens []string) Cmd {
	defer parsErr()
	s := []*Cmd{}
	for i := 0; i < len(tokens); {
		tok := tokens[i]
		if tok == "(" {
			var op string
			jump := 0
			if tokens[i+1] == "(" {		// Allows you to leave the "run" commmand and simply type (for 12 ((println "1") (println "2"))) instead of (for 12 (run (println "1") (println "2")))
				op = "run"
			} else {
				op = tokens[i+1]
				jump = 1
			}
			val, t := kind(op)
			cmd := &Cmd{op,builtinNum(op),[]*Cmd{},nil, nil,t,val}
			if len(s) > 0 {
				cmd.ParentCmd = s[len(s)-1] // To be able to traversal upward (for example in panic).
				s[len(s)-1].Params = append(s[len(s)-1].Params, cmd)
			}
			s = append(s, cmd)
			i += 1 + jump
		} else if tok == ")" {
			if len(s) == 1 {
				break
			}
			s = s[:len(s)-1]
			i++
		} else {
			val, t := kind(tokens[i])
			cmd := Cmd{Op:tokens[i],Value:val,Kind:t}
			c := s[len(s)-1]
			c.Params = append(c.Params, &cmd)
			i++
		}
	}
	if len(s) > 1 {
		panic("Parens are not matching.")
	}
	return *s[0]
}
	