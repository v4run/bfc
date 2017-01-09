package compiler

import (
	"errors"

	"github.com/v4run/bfc/ds"
	"github.com/v4run/bfc/instruction"
)

var (
	// ErrUnexpectedJumpBack occurs when an unexpected ']' occurs
	ErrUnexpectedJumpBack = errors.New("Unexpected ']'")
)

// C defines an interface for the compiler
type C interface {
	Compile() ([]*instruction.I, error)
}

type compiler struct {
	position     int
	code         string
	codeLength   int
	instructions []*instruction.I
}

func (c *compiler) compileFoldableInstruction(s instruction.Symbol) {
	var count uint8 = 1
	for ; c.position < c.codeLength-1 && instruction.Symbol(c.code[c.position+1]) == s; c.position++ {
		count++
	}
	c.addInstruction(s, count)
}

func (c *compiler) addInstruction(s instruction.Symbol, arg uint8) uint8 {
	i := instruction.New(s, arg)
	c.instructions = append(c.instructions, i)
	return uint8(len(c.instructions) - 1)
}

// Compile takes in a code and compiles it into instruction set
func (c *compiler) Compile() ([]*instruction.I, error) {
	st := ds.NewStack()
	for ; c.position < c.codeLength; c.position++ {
		s := instruction.Symbol(c.code[c.position])
		switch s {
		case instruction.PLUS:
			fallthrough
		case instruction.MINUS:
			fallthrough
		case instruction.MOVELEFT:
			fallthrough
		case instruction.MOVERIGHT:
			fallthrough
		case instruction.READ:
			fallthrough
		case instruction.WRITE:
			c.compileFoldableInstruction(s)
		case instruction.JUMPFORWARD:
			st.Push(c.addInstruction(s, 0))
		case instruction.JUMPBACKWARD:
			startPos := st.Pop()
			if startPos == nil {
				return nil, ErrUnexpectedJumpBack
			}
			endPos := c.addInstruction(s, *startPos)
			c.instructions[*startPos].Argument = endPos
		}
	}
	return c.instructions, nil
}

// New returns a new compiler instance
func New(code string) C {
	cl := len(code)
	c := &compiler{
		code:         code,
		codeLength:   cl,
		instructions: make([]*instruction.I, 0, cl),
	}
	return c
}
