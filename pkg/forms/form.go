package forms

import (
	"net/url"
	"regexp"
	"strings"
)

// URLRegExp Source "https://regexr.com/39nr7"
var URLRegExp = regexp.MustCompile("[(http(s)?):\\/\\/(www\\.)?a-zA-Z0-9@:%._\\+~#=]{2,256}\\.[a-z]{2,6}\\b([-a-zA-Z0-9@:%_\\+.~#?&//=]*)")

// Form contains the data passed to a HTML form
type Form struct {
	url.Values
	Errors errors
}

// New creates a new object of the Form type
func New(data url.Values) *Form {
	return &Form{data, errors(map[string][]string{})}
}

// Required adds a field required message to the specified parameter fields
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// MatchesPattern validates a form field input data using a regular expression
func (f *Form) MatchesPattern(field string, pattern *regexp.Regexp) {
	value := f.Get(field)
	if value == "" {
		return
	}

	if !pattern.MatchString(value) {
		f.Errors.Add(field, "This field is invalid")
	}
}

// Valid checks if a user input matches all the validation rules
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
