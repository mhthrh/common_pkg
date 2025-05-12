package validation_test

import (
	"github.com/mhthrh/common_pkg/util/validation"
	"testing"
)

type test struct {
	input    string
	expected bool
}

func TestValidateCanadianMobile(t *testing.T) {
	cases := []test{
		{"14165551234", true},
		{"(416) 555-1234", true},
		{"416-555-1234", true},
		{"0114165551234", false},
		{"1234567", false},
		{"555-5555", false},
	}

	for _, c := range cases {
		if got := validation.MobilePhone(c.input); got != c.expected {
			t.Errorf("ValidateMobile(%q) = %v; want %v", c.input, got, c.expected)
		}
	}
}

func TestValidateFirstName(t *testing.T) {
	cases := []test{
		{"John", true},
		{"Anne-Marie", true},
		{"O'Connor", true},
		{"", false},
		{" A", false},
		{"J", false},
		{"Jo3n", false},
		{"   John", true},
	}

	for _, c := range cases {
		if got := validation.Name(c.input); got != c.expected {
			t.Errorf("ValidateName(%q) = %v; want %v", c.input, got, c.expected)
		}
	}
}

func TestValidateEmail(t *testing.T) {
	cases := []test{
		{"john.doe@example.com", true},
		{"john+label@gmail.com", true},
		{"john.doe@sub.example.co.uk", true},
		{"john@localhost", false},
		{"john@.com", false},
		{"@example.com", false},
	}

	for _, c := range cases {
		if got := validation.Email(c.input); got != c.expected {
			t.Errorf("ValidateEmail(%q) = %v; want %v", c.input, got, c.expected)
		}
	}
}

func TestValidatePassword(t *testing.T) {
	cases := []test{
		{"Password123!", true},
		{"password", false},
		{"Pass123", false},
		{"PASSWORD123$", false},
		{"StrongPass1@", true},
		{"Short1!", false},
	}

	for _, c := range cases {
		if got := validation.Password(c.input); got != c.expected {
			t.Errorf("ValidatePassword(%q) = %v; want %v", c.input, got, c.expected)
		}
	}
}
