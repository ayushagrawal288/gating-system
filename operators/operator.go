package operators

type Operator interface {
	IsValid(key string, value string, attributes map[string]interface{}) (bool, error)
}

const (
	GREATER_THAN          = ">"
	GREATER_THAN_EQUAL_TO = ">="
	LESS_THAN_EQUAL_TO    = "<="
	LESS_THAN             = "<"
	EQUALS                = "=="
	BETWEEN               = "BETWEEN"
	ALL_OF                = "ALLOF"
	NONE_OF               = "NONEOF"
	NOT_EQUALS            = "!="
)

var AllowedOperators = map[string]Operator{
	EQUALS:                &Equals{},
	GREATER_THAN:          &GreaterThan{},
	GREATER_THAN_EQUAL_TO: &GreaterThanEqualTo{},
	LESS_THAN_EQUAL_TO:    &LessThanEqualTo{},
	LESS_THAN:             &LessThan{},
	NOT_EQUALS:            &NotEquals{},
}
