package handlers

import "regexp"

func IsValidEmailFormat(email string) bool {
	r, _ := regexp.Compile("^[\\w.-]+@[\\w.-]+\\.[a-zA-Z]{2,}$")
	
	return r.MatchString(email)
}