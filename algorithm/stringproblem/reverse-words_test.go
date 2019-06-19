package stringproblem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	input string
}

type ans struct {
	output string
}

type question struct {
	p para
	a ans
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				input: " the sky is blue",
			},
			a: ans{
				output: "blue is sky the",
			},
		},
		question{
			p: para{
				input: "a good   example",
			},
			a: ans{
				output: "example good a",
			},
		},
		question{
			p: para{
				input: "  hello world!  ",
			},
			a: ans{
				output: "world! hello",
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.output, ReverseWords(p.input), "输入:%v", p)
	}
}
