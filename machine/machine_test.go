package machine

import (
	"bytes"
	"testing"

	"github.com/v4run/bfc/compiler"
)

func TestExecute(t *testing.T) {
	testCases := []struct {
		in            string
		expected      string
		expectedError error
	}{
		{
			in:            `-[------->+<]>-.-[->+++++<]>++.+++++++..+++.[--->+<]>-----.---[->+++<]>.-[--->+<]>---.+++.------.--------.-[--->+<]>.`,
			expected:      `Hello World!`,
			expectedError: nil,
		},
		{
			in:            `-[]]`,
			expected:      `Hello World!`,
			expectedError: compiler.ErrUnexpectedJumpBack,
		},
		{
			in:            `[],`,
			expected:      ``,
			expectedError: nil,
		},
	}
	for _, c := range testCases {
		out := bytes.NewBufferString("")
		in := bytes.NewBufferString("h")
		m, err := New(c.in, in, out)
		if err != c.expectedError {
			t.Fail()
			t.Errorf("Invalid result. Expected: %v, Got: %v.\n", c.expectedError, err)
			continue
		}
		if err != nil {
			continue
		}
		m.Execute()
		if out.String() != c.expected {
			t.Fail()
			t.Errorf("Invalid result. Expected: %s, Got: %s.\n", c.expected, out.String())
		}
	}
}
