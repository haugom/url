package url

import "testing"

func TestParse(t *testing.T) {
	if err := Parse("broken url"); err != nil {
		t.Log(err)
		t.Fail()
	}
}

