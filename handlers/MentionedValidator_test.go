package handlers

import "testing"

type mentionedTest struct {
	mentionedArg string
	expected     bool
}

var validMentionedTests = []mentionedTest{
	mentionedTest{"@student1@gmail.com", true},
}

var invalidMentionedTests = []mentionedTest{
	mentionedTest{"student1@gmail.com", false},
	mentionedTest{"@student1", false},
}

func TestValidMentions(t *testing.T) {
	for _, test := range validMentionedTests {
		if output := IsValidMentionedFormat(test.mentionedArg); output != test.expected {
			t.Errorf("got %t, wanted %t", output, test.expected)
		}
	}
}

func TestInValidMentions(t *testing.T) {
	for _, test := range invalidMentionedTests {
		if output := IsValidMentionedFormat(test.mentionedArg); output != test.expected {
			t.Errorf("got %t, wanted %t", output, test.expected)
		}
	}
}
