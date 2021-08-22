package forms

// errors is a map that stores a field and the value of its error message
type errors map[string][]string

// Add sets an error message as the value of the field parameter
func (err errors) Add(field, message string) {
	err[field] = append(err[field], message)
}

// Get returns the errors belonging to a specific field
func (err errors) Get(field string) string {
	errs := err[field]
	if len(errs) == 0 {
		return ""
	}

	return errs[0]
}
