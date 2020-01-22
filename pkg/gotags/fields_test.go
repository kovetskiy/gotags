package gotags

import (
	"testing"
)

func TestParseFieldsEmpty(t *testing.T) {
	_, err := ParseFields("")
	if err != nil {
		t.Fatalf("unexpected error from ParseFields: %s", err)
	}
}

func TestParseFieldsLanguage(t *testing.T) {
	set, err := ParseFields("+l")
	if err != nil {
		t.Fatalf("unexpected error from ParseFields: %s", err)
	}
	if !set.Includes(Language) {
		t.Fatal("expected set to include Language")
	}
}

func TestParseFieldsInvalid(t *testing.T) {
	_, err := ParseFields("junk")
	if err == nil {
		t.Fatal("expected ParseFields to return error")
	}
	if _, ok := err.(ErrInvalidFields); !ok {
		t.Fatalf("expected ParseFields to return error of type ErrInvalidFields, got %T", err)
	}
}
