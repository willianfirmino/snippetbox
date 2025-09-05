package validator

import (
	"regexp"
	"testing"
)

func TestValidator(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		v := &Validator{}
		if !v.Valid() {
			t.Error("want valid; got invalid")
		}

		v.AddFieldError("foo", "bar")
		if v.Valid() {
			t.Error("want invalid; got valid")
		}
	})

	t.Run("AddFieldError", func(t *testing.T) {
		v := &Validator{}
		v.AddFieldError("foo", "bar")

		if _, ok := v.FieldErrors["foo"]; !ok {
			t.Error("want field error to be added")
		}
	})

	t.Run("CheckField", func(t *testing.T) {
		v := &Validator{}
		v.CheckField(true, "foo", "bar")

		if _, ok := v.FieldErrors["foo"]; ok {
			t.Error("want field error to not be added")
		}

		v.CheckField(false, "foo", "bar")
		if _, ok := v.FieldErrors["foo"]; !ok {
			t.Error("want field error to be added")
		}
	})

	t.Run("NotBlank", func(t *testing.T) {
		if NotBlank("") {
			t.Error("want not blank to be false")
		}
		if !NotBlank("foo") {
			t.Error("want not blank to be true")
		}
	})

	t.Run("MaxChars", func(t *testing.T) {
		if !MaxChars("foo", 3) {
			t.Error("want max chars to be true")
		}
		if MaxChars("foobar", 3) {
			t.Error("want max chars to be false")
		}
	})

	t.Run("MinChars", func(t *testing.T) {
		if !MinChars("foo", 3) {
			t.Error("want min chars to be true")
		}
		if MinChars("foo", 4) {
			t.Error("want min chars to be false")
		}
	})

	t.Run("Matches", func(t *testing.T) {
		rx := regexp.MustCompile("a.b")
		if !Matches("aob", rx) {
			t.Error("want matches to be true")
		}
		if Matches("axxb", rx) {
			t.Error("want matches to be false")
		}
	})

	t.Run("PermittedValue", func(t *testing.T) {
		if !PermittedValue("foo", "foo", "bar") {
			t.Error("want permitted value to be true")
		}
		if PermittedValue("baz", "foo", "bar") {
			t.Error("want permitted value to be false")
		}
	})
}
