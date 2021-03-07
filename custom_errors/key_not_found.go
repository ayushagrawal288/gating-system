package custom_errors

type KeyNotFoundError struct {
	Key string
}

func (m *KeyNotFoundError) Error() string {
	return "Value for key: " + m.Key + " not found in user's attributes"
}
