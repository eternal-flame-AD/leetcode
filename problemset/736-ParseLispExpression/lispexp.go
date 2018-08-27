package lispexp

import (
	"bytes"
	"strconv"
	"strings"
)

const (
	OPCODE_LET = iota
	OPCODE_ADD
	OPCODE_MULT
	OPCODE_NUMERIC
	OPCODE_VARIABLE
)

type LispContext struct {
	variables map[string]int64
}

type LispExp struct {
	context  LispContext
	exptype  int8
	OPs      []*LispExp
	ancestor *LispExp
	varname  string
	result   *int64
}

func (c *LispExp) searchContextForVariable(key string) int64 {
	this := c
	for {
		if val, ok := this.context.variables[key]; ok {
			return val
		}
		this = this.ancestor
	}
}

func (c *LispExp) eval() int64 {
	if c.result != nil {
		return *c.result
	}
	switch c.exptype {
	case OPCODE_ADD:
		return c.OPs[0].eval() + c.OPs[1].eval()
	case OPCODE_MULT:
		return c.OPs[0].eval() * c.OPs[1].eval()
	case OPCODE_LET:
		for i := 0; i < len(c.OPs)-2; i += 2 {
			varname := c.OPs[i].varname
			c.context.variables[varname] = c.OPs[i+1].eval()
		}
		return c.OPs[len(c.OPs)-1].eval()
	case OPCODE_VARIABLE:
		return c.searchContextForVariable(c.varname)
	default:
		panic("HELP")
	}
}

func splitExprstrParts(s string) []string {
	pointer := 0
	res := make([]string, 0)
	thispart := bytes.NewBufferString("")
	depth := 0
	for pointer < len(s) {
		switch s[pointer] {
		case ')':
			depth--
			thispart.WriteByte(s[pointer])
		case '(':
			depth++
			thispart.WriteByte(s[pointer])
		case ' ':
			if depth == 0 {
				res = append(res, thispart.String())
				thispart = bytes.NewBufferString("")
				break
			}
			fallthrough
		default:
			thispart.WriteByte(s[pointer])
		}
		pointer++
	}
	res = append(res, thispart.String())
	return res
}

func ConsumeLispExpString(exprstr string, ancestor *LispExp) *LispExp {
	parts := splitExprstrParts(strings.TrimSuffix(strings.TrimPrefix(exprstr, "("), ")"))
	thisexp := new(LispExp)
	thisexp.context.variables = make(map[string]int64, 0)
	thisexp.ancestor = ancestor
	switch {
	case len(parts) == 1:
		if parts[0][0] < 'a' { //number
			number, _ := strconv.ParseInt(parts[0], 10, 32)
			thisexp.exptype = OPCODE_NUMERIC
			thisexp.result = &number
		} else { //variable
			thisexp.exptype = OPCODE_VARIABLE
			thisexp.varname = parts[0]
		}
	case parts[0] == "add":
		thisexp.exptype = OPCODE_ADD
		thisexp.OPs = []*LispExp{
			ConsumeLispExpString(parts[1], thisexp),
			ConsumeLispExpString(parts[2], thisexp),
		}
	case parts[0] == "mult":
		thisexp.exptype = OPCODE_MULT
		thisexp.OPs = []*LispExp{
			ConsumeLispExpString(parts[1], thisexp),
			ConsumeLispExpString(parts[2], thisexp),
		}
	case parts[0] == "let":
		thisexp.exptype = OPCODE_LET
		thisexp.OPs = make([]*LispExp, len(parts)-1)
		for i := 1; i < len(parts); i++ {
			thisexp.OPs[i-1] = ConsumeLispExpString(parts[i], thisexp)
		}
	}
	return thisexp
}

func evaluate(expression string) int {
	return int(ConsumeLispExpString(expression, nil).eval())
}
