package api

import (
	"gating-system/custom_errors"
	"gating-system/operators"
	"strconv"
	"strings"
)

type Gating struct {
	userAttributes map[string]interface{}
}

func (e *Gating) IsAllowed(featureName string, conditionalExpression string, userAttributes map[string]interface{}) (allowed bool, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	e.userAttributes = userAttributes
	resultStr := e.evaluate(conditionalExpression)
	return strconv.ParseBool(resultStr)

}

func (e *Gating) evaluateExpression(exp string) (bool, error) {
	ok, err := strconv.ParseBool(exp)
	if err == nil {
		return ok, nil
	}

	arr := strings.Split(exp, " ")
	if len(arr) < 2 {
		return false, &custom_errors.InvalidExpressionError{Key: exp}
	}

	operator, ok := operators.AllowedOperators[arr[1]]
	if !ok {
		return false, &custom_errors.InvalidOperator{Operator: arr[1], Key: arr[0]}
	}

	return operator.IsValid(arr[0], arr[2], e.userAttributes)
}

func precedence(op string) int {
	if op == "&&" {
		return 2
	}
	if op == "||" {
		return 1
	}
	return 0
}

func (e *Gating) applyOp(a string, b string, op string) string {
	aOK, err := e.evaluateExpression(a)
	if err != nil {
		panic(err)
	}
	bOK, err := e.evaluateExpression(b)
	if err != nil {
		panic(err)
	}
	if op == "&&" {
		return strconv.FormatBool(aOK && bOK)
	} else {
		return strconv.FormatBool(aOK || bOK)
	}
}

func (e *Gating) evaluate(tokens string) string {
	values := []string{}
	ops := []string{}
	i := 0
	var op string
	brancesAndOps := map[string]bool{"&": true, "|": true, "(": true, ")": true}
	for i < len(tokens) {
		if tokens[i] == ' ' {
			i += 1
			continue
		} else if tokens[i] == '(' {
			ops = append(ops, string(tokens[i]))
		} else if ok, _ := brancesAndOps[string(tokens[i])]; !ok {
			s := ""
			for !ok && i < len(tokens) {
				s += string(tokens[i])
				i += 1
				if i < len(tokens) {
					ok, _ = brancesAndOps[string(tokens[i])]
				}
			}
			values = append(values, s)
			i -= 1
		} else if string(tokens[i]) == ")" {
			for len(ops) != 0 && ops[len(ops)-1] != "(" {
				val2, val1 := values[len(values)-1], values[len(values)-2]
				values = values[:len(values)-2]
				op, ops = ops[len(ops)-1], ops[:len(ops)-1]
				res := e.applyOp(val1, val2, op)
				values = append(values, res)
			}
			if len(ops) > 0 {
				ops = ops[:len(ops)-1]
			}
		} else {
			tempOp := string(tokens[i]) + string(tokens[i+1])
			i += 1
			for len(ops) != 0 &&
				precedence(ops[len(ops)-1]) >=
					precedence(tempOp) {

				val2, val1 := values[len(values)-1], values[len(values)-2]
				values = values[:len(values)-2]

				op, ops = ops[len(ops)-1], ops[:len(ops)-1]
				values = append(values, e.applyOp(val1, val2, op))
			}
			ops = append(ops, tempOp)
		}
		i += 1
	}

	for len(ops) != 0 {
		val2, val1 := values[len(values)-1], values[len(values)-2]
		values = values[:len(values)-2]
		op, ops = ops[len(ops)-1], ops[:len(ops)-1]
		res := e.applyOp(val1, val2, op)
		values = append(values, res)
	}
	return values[len(values)-1]
}
