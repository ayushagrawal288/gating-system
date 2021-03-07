package custom_errors

type InvalidOperator struct {
	Type     string
	Operator string
	Key      string
}

func (m *InvalidOperator) Error() string {
	if m.Type == "" {
		return "Operation: " + m.Operator + "not allowed for key: " + m.Key + " whose value is of type: " + m.Type
	} else {
		return "Invalid Operator passed: " + m.Operator + " for key: " + m.Key
	}
}
