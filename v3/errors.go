package v3

// InvalidKeyError is an error when the key length not long enough or contains non-alphanumeric characters
type InvalidKeyError struct{}

func (e *InvalidKeyError) Error() string {
	return "key must be at least 1 character in length and only contain alphanumeric characters"
}

// ParamLengthError is an error when the param length is less than 1 or greater than 512 characters
type ParamLengthError struct{}

func (e *ParamLengthError) Error() string {
	return "params must between 1 and 512 characters in length"
}
