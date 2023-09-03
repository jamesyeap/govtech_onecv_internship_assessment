package handlers

import "testing"

type emailTest struct {
	emailArg string
	expected bool
}

var validEmailTests = []emailTest{
	emailTest{"teacherken@gmail.com", true},
	emailTest{"teacherken@hotmail.com", true},
	emailTest{"teacherken@mail.com.sg", true},
}

var invalidEmailTests = []emailTest{
	emailTest{"teacherken", false},
	emailTest{"teacherken@gmail", false},
	emailTest{"teacherken@gmail", false},
}

func TestValidEmails(t *testing.T) {
	for _, test := range validEmailTests {
		if output := IsValidEmailFormat(test.emailArg); output != test.expected {
			t.Errorf("got %t, wanted %t", output, test.expected)
		}
	}
}

func TestInValidEmails(t *testing.T) {
	for _, test := range invalidEmailTests {
		if output := IsValidEmailFormat(test.emailArg); output != test.expected {
			t.Errorf("got %t, wanted %t", output, test.expected)
		}
	}
}
