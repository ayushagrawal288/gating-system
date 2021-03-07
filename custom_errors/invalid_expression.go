package custom_errors

type InvalidExpressionError struct {
	Key string
}

func (m *InvalidExpressionError) Error() string {
	return "Invalid expression passed: " + m.Key
}
