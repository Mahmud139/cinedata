package validator

import "regexp"

// emailverifier "github.com/AfterShip/email-verifier"

var	EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// var verifier = emailverifier.NewVerifier()

type Validator struct {
	Errors map[string]string
}

func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) AddError(key, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

func In(value string, list ...string) bool {
	for i := range list {
		if value == list[i] {
			return true
		}
	}

	return false
}

func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

func Unique(values []string) bool {
	uniqueValues := make(map[string]bool)

	for _, value := range values {
		uniqueValues[value] = true
	}

	return len(values) == len(uniqueValues)
}

// func (f *Form) VerifyEmail(field string) {
// 	email := f.Get(field)
// 	result, err := verifier.Verify(email)
// 	if err != nil {
// 		f.Errors.Add(field, "invalid email address.")
// 	}

// 	if !result.Syntax.Valid {
// 		f.Errors.Add(field, "email address syntext is invalid.")
// 	}

// 	if result.Disposable {
// 		f.Errors.Add(field, "we don't accept disposable email address.")
// 	}

// 	if !result.HasMxRecords {
// 		f.Errors.Add(field, "no MX record found for the domain, no such host.")
// 	}

// 	if result.RoleAccount {
// 		f.Errors.Add(field, "we don't accept role-based account.")
// 	}
// }