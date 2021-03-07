package main

import (
	"fmt"
	"gating-system/api"
	"strconv"
)

func main() {
	featureName := "Test"
	userAttributes := map[string]interface{}{
		"age":               24,
		"gender":            "female",
		"past_order_amount": 11000,
	}
	expression := "(age > 25 && gender == male) || past_order_amount > 10000"
	gateway := api.Gating{}
	allowed, err := gateway.IsAllowed(featureName, expression, userAttributes)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Feature is allowed: " + strconv.FormatBool(allowed))
	}
}
