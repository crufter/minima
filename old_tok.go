package minima

import(
	"github.com/opesun/lexer"
)

var token_exprs = []lexer.TokenExpr{
    {`[ \t]+`,									itemIgnore},	// Whitespace
    {`\-\-[^\n]*`,								itemIgnore},	// Comment
	{`\n`,										itemIgnore},	// Newline
    {`\(`,										itemRParen},
    {`\)`,										itemLParen},
    {`[0-9]+`,									itemInt},
	{`"(?:[^"\\]|\\.)*"`,						itemString},
    {`[\<\>\!\=\+\-\*\/A-Za-z][A-Za-z0-9_]*`,	itemID},
}

func TokenizeOld(source string) []string {
	tokens, _ := lexer.Lex("\n(" + source + ")", token_exprs)
	toks := []string{}
	for _, v := range tokens {
		toks = append(toks, v.Text)
	}
	return toks
}