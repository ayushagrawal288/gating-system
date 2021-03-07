package custom_errors

type InvalidType struct {
	ExpectedType string
	Err          error
	Key          string
}

func (m *InvalidType) Error() string {
	return "For key: " + m.Key + " the expected value type is: " + m.ExpectedType + " got following error on parsing: " + m.Err.Error()
}
