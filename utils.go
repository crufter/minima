package minima

import(
	"fmt"
	"strings"
)

func Visualize(cmd Cmd, indent string, lev int) {
	fmt.Println(strings.Repeat(indent, lev), cmd.Op)
	for _, v := range cmd.Params {
		Visualize(*v, indent, lev+1)
	}
}