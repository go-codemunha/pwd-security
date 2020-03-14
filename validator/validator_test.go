package validator

import (
	"testing"
)

func TestPasswordOriginal(t *testing.T) {
	t.Run("FooBar123!", func(t *testing.T) {
		given := "FooBar123!"
		want := true
		get := Password(given)
		if want != get {
			t.Errorf("given %q want %t but get %t", given, want, get)
		}

	})
	t.Run("foobar", func(t *testing.T) {
		given := "foobar"
		want := false
		get := Password(given)
		if want != get {
			t.Errorf("given %q want %t but get %t", given, want, get)
		}

	})
	t.Run("foobar123", func(t *testing.T) {
		given := "foobar123"
		want := false
		get := Password(given)
		if want != get {
			t.Errorf("given %q want %t but get %t", given, want, get)
		}

	})
	t.Run("@", func(t *testing.T) {
		given := "@"
		want := false
		get := Password(given)
		if want != get {
			t.Errorf("given %q want %t but get %t", given, want, get)
		}

	})
	t.Run("@@@@@@@2wS@", func(t *testing.T) {
		given := "@@@@@@@2wS@ "
		want := false
		get := Password(given)
		if want != get {
			t.Errorf("given %q want %t but get %t", given, want, get)
		}

	})
}
