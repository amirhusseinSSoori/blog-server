package validation

func ErrorMessage() map[string]string {
	return map[string]string{
		"required": "the filed is required",
		"email":    "the filed must have valid email address",
		"min":      "should be more than the limit",
		"max":      "should be less than the limit",
	}
}
