package operators

import (
	"testing"
)

func TestEquals(t *testing.T) {
	m := map[string]interface{}{
		"value": 1.23,
	}
	op := Equals{}
	ok, err := op.IsValid("value", "1.23", m)
	if err != nil {
		t.Error("Failed while checking float: " + err.Error())
	}
	if !ok {
		t.Log("Should be equal")
		t.Fail()
	}

	m = map[string]interface{}{
		"value": "asd",
	}
	// op := Equals{}
	ok, err = op.IsValid("value", "asd", m)
	if err != nil {
		t.Error("Failed while checking string: " + err.Error())
	}
	if !ok {
		t.Log("Should be equal")
		t.Fail()
	}

	m = map[string]interface{}{
		"value": 5,
	}
	// op := Equals{}
	ok, err = op.IsValid("value", "5", m)
	if err != nil {
		t.Error("Failed while checking int: " + err.Error())
	}
	if !ok {
		t.Log("Should be equal")
		t.Fail()
	}

	m = map[string]interface{}{
		"value": true,
	}
	// op := Equals{}
	ok, err = op.IsValid("value", "true", m)
	if err != nil {
		t.Error("Failed while checking bool: " + err.Error())
	}
	if !ok {
		t.Log("Should be equal")
		t.Fail()
	}
}
