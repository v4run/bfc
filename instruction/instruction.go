package instruction

// Symbol is used to denote the different symbols in the language
type Symbol byte

var (
	// PLUS means increment the value in the cell pointed by data pointer by 1
	PLUS Symbol = '+'

	// MINUS means decrement the value in current the cell pointed by data pointer by 1
	MINUS Symbol = '-'

	// MOVERIGHT means increment the data pointer by 1
	MOVERIGHT Symbol = '>'

	// MOVELEFT means decrement the data pointer by 1
	MOVELEFT Symbol = '<'

	// WRITE means write the value in current cell to output
	WRITE Symbol = '.'

	// READ means read value into current cell from input
	READ Symbol = ','

	// JUMPFORWARD means if the value in current cell is 0, then jump to
	// instruction after the matching ']'
	JUMPFORWARD Symbol = '['

	// JUMPBACKWARD means if the value in current cell is not 0, then jump to
	// instruction after the matching '['
	JUMPBACKWARD Symbol = ']'
)

// String converst symbol to string
func (s Symbol) String() string {
	return string(s)
}

// I represents an instruction in the code
type I struct {
	S        Symbol
	Argument uint8
}

// New creates and returns a new instruction
func New(s Symbol, a uint8) *I {
	return &I{
		S:        s,
		Argument: a,
	}
}
