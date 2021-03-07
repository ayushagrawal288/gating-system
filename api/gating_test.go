package api

import "testing"

func TestGatewayEvaluate(t *testing.T) {
	mp := map[string]interface{}{
		"age":               24,
		"gender":            "female",
		"past_order_amount": 9000,
	}

	gt := Gating{}
	exp := "age <= 24"
	gt.userAttributes = mp
	ok, err := gt.evaluateExpression(exp)
	if err != nil {
		t.Error("Failed while checking " + exp + " : " + err.Error())
	}
	if !ok {
		t.Log("Should have passed")
		t.Fail()
	}

	exp = "age2 <= 24"
	ok, err = gt.evaluateExpression(exp)
	if err == nil {
		t.Error("Succeded while checking " + exp + " : " + err.Error())
	}
	if ok {
		t.Log("Should have failed")
		t.Fail()
	}

}

func TestGatewayIsAllowed(t *testing.T) {
	mp := map[string]interface{}{
		"age":               24,
		"gender":            "female",
		"past_order_amount": 9000,
	}

	gt := Gating{}
	exp := "(age > 25 && gender == male) || past_order_amount > 10000)"
	ok, err := gt.IsAllowed("Name", exp, mp)
	if err != nil {
		t.Error("Failed while checking " + exp + " : " + err.Error())
	}
	if ok {
		t.Log("Should have failed")
		t.Fail()
	}

	mp = map[string]interface{}{
		"age":               25,
		"gender":            "male",
		"past_order_amount": 9000,
	}
	exp = "(age > 25 && gender == male) || past_order_amount > 10000"
	ok, err = gt.IsAllowed("Name", exp, mp)
	if err != nil {
		t.Error("Failed while checking " + exp + " : " + err.Error())
	}
	if ok {
		t.Log("Should have failed")
		t.Fail()
	}

	mp = map[string]interface{}{
		"age":               24,
		"gender":            "female",
		"past_order_amount": 11000,
	}
	exp = "(age > 25 && gender == male) || past_order_amount > 10000"
	ok, err = gt.IsAllowed("Name", exp, mp)
	if err != nil {
		t.Error("Failed while checking " + exp + " : " + err.Error())
	}
	if !ok {
		t.Log("Should have failed")
		t.Fail()
	}
}
