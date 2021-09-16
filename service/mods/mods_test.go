package mods

import "testing"

// ### MissingChar
func TestMissingChar(t *testing.T) {
	str, err := MissingChar("Hello", "*")
	if err != nil {
		t.Fail()
	}
	t.Log(str)
}

func TestMissingCharEmptyWord(t *testing.T) {
	_, err := MissingChar("", "*")
	if err == nil {
		t.Fail()
	}
	t.Log(err)
}

func TestMissingCharEmptyToken(t *testing.T) {
	_, err := MissingChar("Hello", "")
	if err == nil {
		t.Fail()
	}
	t.Log(err)
}

func TestMissingCharLogToken(t *testing.T) {
	_, err := MissingChar("Hello", "oj")
	if err == nil {
		t.Fail()
	}
	t.Log(err)
}

func TestSwitchChars(t *testing.T) {
	str, err := SwitchChars("Hello")
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	t.Log(str)
}

func TestSwitchCharsEmptyWord(t *testing.T) {
	_, err := SwitchChars("")
	if err == nil {
		t.Fail()
	}
	t.Log(err)
}

func TestCharsToNumbers(t *testing.T) {
	str, err := CharsToNumbers("izeasgtbo IZEASGTBO")
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	t.Log(str)
}

func TestCharsToNumbers2(t *testing.T) {
	str, err := CharsToNumbers("HelloWorld")
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	t.Log(str)
}

func TestCharsToNumbersEmptyWord(t *testing.T) {
	_, err := CharsToNumbers("")
	if err == nil {
		t.Fail()
	}
	t.Log(err)
}

func TestReverse(t *testing.T) {
	str, err := Reverse("Hello")
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	t.Log(str)
}

func TestReverseEmptyWord(t *testing.T) {
	_, err := Reverse("")
	if err == nil {
		t.Fail()
	}
	t.Log(err)
}