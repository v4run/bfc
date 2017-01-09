package machine

import (
	"fmt"
	"io"
	"os"

	"github.com/v4run/bfc/compiler"
	"github.com/v4run/bfc/instruction"
)

// M defines an interface for a machine
type M interface {
	Execute()
}

type machine struct {
	memory            [30000]uint8
	dp                uint8
	ip                uint8
	instructions      []*instruction.I
	instructionLength uint8
	output            io.Writer
	input             io.Reader
	buf               []byte
}

// Execute executes the instructions in the machine
func (m *machine) Execute() {
	for m.ip < m.instructionLength {
		i := m.instructions[m.ip]
		switch i.S {
		case instruction.PLUS:
			m.memory[m.dp] += i.Argument
		case instruction.MINUS:
			m.memory[m.dp] -= i.Argument
		case instruction.MOVELEFT:
			m.dp -= i.Argument
		case instruction.MOVERIGHT:
			m.dp += i.Argument
		case instruction.READ:
			var j uint8
			for ; j < i.Argument; j++ {
				m.readChar()
			}
		case instruction.WRITE:
			var j uint8
			for ; j < i.Argument; j++ {
				m.writeChar()
			}
		case instruction.JUMPFORWARD:
			if m.memory[m.dp] == 0 {
				m.ip = i.Argument
				continue
			}
		case instruction.JUMPBACKWARD:
			if m.memory[m.dp] != 0 {
				m.ip = i.Argument
				continue
			}
		}
		m.ip++
	}
}

func (m *machine) readChar() {
	bc, err := m.input.Read(m.buf)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if bc != 1 {
		fmt.Println("Invalid number of characters read")
		os.Exit(1)
	}
	m.memory[m.dp] = uint8(m.buf[0])
}

func (m *machine) writeChar() {
	m.buf[0] = byte(m.memory[m.dp])
	bc, err := m.output.Write(m.buf)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if bc != 1 {
		fmt.Println("Invalid number of characters wrote")
		os.Exit(1)
	}
}

// New returns a new instance of the machine
func New(c string, i io.Reader, o io.Writer) (M, error) {
	ins, err := compiler.New(c).Compile()
	if err != nil {
		return nil, err
	}
	m := &machine{
		input:             i,
		output:            o,
		instructions:      ins,
		buf:               make([]byte, 1),
		instructionLength: uint8(len(ins)),
	}
	return m, nil
}
