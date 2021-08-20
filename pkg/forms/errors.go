package forms

type errors map[string][]string

func (err errors) Add(field, message string) {
	err[field] = append(err[field], message)
}

func (err errors) Get(field string) string {
	errs := err[field]
	if len(errs) == 0 {
		return ""
	}

	return errs[0]
}
