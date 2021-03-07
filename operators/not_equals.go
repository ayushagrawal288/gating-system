package operators

import (
	"gating-system/custom_errors"
	"reflect"
	"strconv"
)

type NotEquals struct {
}

func (o *NotEquals) IsValid(key string, value string, attributes map[string]interface{}) (bool, error) {
	actualValue, ok := attributes[key]
	if !ok {
		return false, &custom_errors.KeyNotFoundError{Key: key}
	}

	expectedType := reflect.TypeOf(actualValue).String()
	switch t := actualValue.(type) {
	case string:
		return t != value, nil
	case int:
		v, err := strconv.Atoi(value)
		if err != nil {
			return false, &custom_errors.InvalidType{ExpectedType: expectedType, Err: err, Key: key}
		}
		return t != v, nil
	case float64:
		v, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return false, &custom_errors.InvalidType{ExpectedType: expectedType, Err: err, Key: key}
		}
		return t != v, nil
	case bool:
		v, err := strconv.ParseBool(value)
		if err != nil {
			return false, &custom_errors.InvalidType{ExpectedType: expectedType, Err: err, Key: key}
		}
		return t != v, nil
	default:
		return false, &custom_errors.InvalidOperator{Type: expectedType, Operator: NOT_EQUALS, Key: key}
	}
}
