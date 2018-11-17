package ppaster

import (
	"testing"
)

func TestConvertString(t *testing.T) {
	// Makes sure there are no explicit line blocks in the string
	testStr := `I
  Am
  A
  Test`

	convertedStr := ConvertString(testStr)

	for i, r := range convertedStr {
		if r == 10 {
			t.Errorf("We did not eliminate the line breaks correctly!!  Line break at index: %d", i)
		}
	}
}
