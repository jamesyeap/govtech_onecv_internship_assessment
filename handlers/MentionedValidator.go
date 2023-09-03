package handlers

import "regexp"

func IsValidMentionedFormat(mentioned string) bool {
	r, _ := regexp.Compile("^@[\\w.-]+@[\\w.-]+\\.[a-zA-Z]{2,}$")
	
	return r.MatchString(mentioned)
}