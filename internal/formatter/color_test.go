package formatter

import "testing"

func TestApplyANSISingle(t *testing.T) {
	const text = "Test ANSI"
	const expected = "\x1b[34m\x1b[KTest ANSI\x1b[m\x1b[K"
	found := ApplyANSI(text, Blue)

	if expected != found {
		t.Errorf("Expected format:\n%v\nFound format:\n%v", expected, found)
	}

}

func TestApplyANSIMultiple(t *testing.T) {
	const text = "Test multiple ANSI"
	const expected = "\x1b[01;34m\x1b[KTest multiple ANSI\x1b[m\x1b[K"
	found := ApplyANSI(text, Bold, Blue)

	if expected != found {
		t.Errorf("Expected format:\n%v\nFound format:\n%v", expected, found)
	}
}
