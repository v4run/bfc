package compiler

import (
	"testing"

	"github.com/v4run/bfc/instruction"
)

func TestCompile(t *testing.T) {
	testCases := []struct {
		in            string
		expected      []*instruction.I
		expectedError error
	}{
		{
			in: "+++++--->><.,[++][-[-]++]",
			expected: []*instruction.I{
				instruction.New(instruction.PLUS, 5),
				instruction.New(instruction.MINUS, 3),
				instruction.New(instruction.MOVERIGHT, 2),
				instruction.New(instruction.MOVELEFT, 1),
				instruction.New(instruction.WRITE, 1),
				instruction.New(instruction.READ, 1),
				instruction.New(instruction.JUMPFORWARD, 8),
				instruction.New(instruction.PLUS, 2),
				instruction.New(instruction.JUMPBACKWARD, 6),
				instruction.New(instruction.JUMPFORWARD, 15),
				instruction.New(instruction.MINUS, 1),
				instruction.New(instruction.JUMPFORWARD, 13),
				instruction.New(instruction.MINUS, 1),
				instruction.New(instruction.JUMPBACKWARD, 11),
				instruction.New(instruction.PLUS, 2),
				instruction.New(instruction.JUMPBACKWARD, 9),
			},
			expectedError: nil,
		},
		{
			in: "-[------->+<]>-.-[->+++++<]>++.+++++++..+++.[--->+<]>-----.---[->+++<]>.-[--->+<]>---.+++.------.--------.-[--->+<]>.",
			expected: []*instruction.I{
				instruction.New(instruction.MINUS, 1),
				instruction.New(instruction.JUMPFORWARD, 6),
				instruction.New(instruction.MINUS, 7),
				instruction.New(instruction.MOVERIGHT, 1),
				instruction.New(instruction.PLUS, 1),
				instruction.New(instruction.MOVELEFT, 1),
				instruction.New(instruction.JUMPBACKWARD, 1),
				instruction.New(instruction.MOVERIGHT, 1),
				instruction.New(instruction.MINUS, 1),
				instruction.New(instruction.WRITE, 1),
				instruction.New(instruction.MINUS, 1),
				instruction.New(instruction.JUMPFORWARD, 16),
				instruction.New(instruction.MINUS, 1),
				instruction.New(instruction.MOVERIGHT, 1),
				instruction.New(instruction.PLUS, 5),
				instruction.New(instruction.MOVELEFT, 1),
				instruction.New(instruction.JUMPBACKWARD, 11),
				instruction.New(instruction.MOVERIGHT, 1),
				instruction.New(instruction.PLUS, 2),
				instruction.New(instruction.WRITE, 1),
				instruction.New(instruction.PLUS, 7),
				instruction.New(instruction.WRITE, 2),
				instruction.New(instruction.PLUS, 3),
				instruction.New(instruction.WRITE, 1),
				instruction.New(instruction.JUMPFORWARD, 29),
				instruction.New(instruction.MINUS, 3),
				instruction.New(instruction.MOVERIGHT, 1),
				instruction.New(instruction.PLUS, 1),
				instruction.New(instruction.MOVELEFT, 1),
				instruction.New(instruction.JUMPBACKWARD, 24),
				instruction.New(instruction.MOVERIGHT, 1),
				instruction.New(instruction.MINUS, 5),
				instruction.New(instruction.WRITE, 1),
				instruction.New(instruction.MINUS, 3),
				instruction.New(instruction.JUMPFORWARD, 39),
				instruction.New(instruction.MINUS, 1),
				instruction.New(instruction.MOVERIGHT, 1),
				instruction.New(instruction.PLUS, 3),
				instruction.New(instruction.MOVELEFT, 1),
				instruction.New(instruction.JUMPBACKWARD, 34),
				instruction.New(instruction.MOVERIGHT, 1),
				instruction.New(instruction.WRITE, 1),
				instruction.New(instruction.MINUS, 1),
				instruction.New(instruction.JUMPFORWARD, 48),
				instruction.New(instruction.MINUS, 3),
				instruction.New(instruction.MOVERIGHT, 1),
				instruction.New(instruction.PLUS, 1),
				instruction.New(instruction.MOVELEFT, 1),
				instruction.New(instruction.JUMPBACKWARD, 43),
				instruction.New(instruction.MOVERIGHT, 1),
				instruction.New(instruction.MINUS, 3),
				instruction.New(instruction.WRITE, 1),
				instruction.New(instruction.PLUS, 3),
				instruction.New(instruction.WRITE, 1),
				instruction.New(instruction.MINUS, 6),
				instruction.New(instruction.WRITE, 1),
				instruction.New(instruction.MINUS, 8),
				instruction.New(instruction.WRITE, 1),
				instruction.New(instruction.MINUS, 1),
				instruction.New(instruction.JUMPFORWARD, 64),
				instruction.New(instruction.MINUS, 3),
				instruction.New(instruction.MOVERIGHT, 1),
				instruction.New(instruction.PLUS, 1),
				instruction.New(instruction.MOVELEFT, 1),
				instruction.New(instruction.JUMPBACKWARD, 59),
				instruction.New(instruction.MOVERIGHT, 1),
				instruction.New(instruction.WRITE, 1),
			},
			expectedError: nil,
		},
		{
			in:            `-[]]`,
			expected:      nil,
			expectedError: ErrUnexpectedJumpBack,
		},
	}

	for _, c := range testCases {
		ins, err := New(c.in).Compile()
		if err != c.expectedError {
			t.Errorf("Error mismatch. Expected: %v. Got: %v\n", c.expectedError, err)
			t.Fail()
			continue
		}
		if len(ins) != len(c.expected) {
			t.Errorf("Instruction count mismatch. Expected: %d. Got: %d\n", len(c.expected), len(ins))
			t.Fail()
			continue
		}
		for i, e := range c.expected {
			if e.S != ins[i].S || e.Argument != ins[i].Argument {
				t.Errorf("Wrong instruction at position %d. Expected: %v. Got: %v\n", i, e, ins[i])
				t.Fail()
			}
		}
	}
}
